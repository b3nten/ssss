package parser

import (
	"fmt"
	"hash/fnv"

	_ "embed"

	lua "github.com/yuin/gopher-lua"
)

//go:embed schema_globals.lua
var schemaGlobals string

func GenerateSchema(file, name string) (*Schema, error) {
	var err error
	L := lua.NewState()
	defer L.Close()
	if err = L.DoString(schemaGlobals); err != nil {
		return nil, err
	}
	if err = L.DoString(string(file)); err != nil {
		return nil, err
	}

	lstructs := []*lua.LTable{}
	structs := map[string]*StructType{}
	ltos := map[*lua.LTable]*StructType{}

	if tbl, ok := L.GetGlobal("_G").(*lua.LTable); ok {
		tbl.ForEach(func(key lua.LValue, value lua.LValue) {
			if structTbl, ok := value.(*lua.LTable); ok {
				if structTbl.RawGet(lua.LString("type")) == lua.LString("struct") {
					sv := StructType{
						Name: key.String(),
						ID:   uint16(hash(key.String())),
						UUID: structTbl.RawGet(lua.LString("uuid")).String(),
					}
					structs[sv.UUID] = &sv
					ltos[structTbl] = &sv
					lstructs = append(lstructs, structTbl)
				}
			}
		})

		if len(structs) == 0 {
			return nil, fmt.Errorf("no structs defined in schema")
		}

		for _, lstruct := range lstructs {
			if fields, ok := lstruct.RawGet(lua.LString("fields")).(*lua.LTable); ok {
				fields.ForEach(func(key lua.LValue, value lua.LValue) {
					if fieldTbl, ok := value.(*lua.LTable); ok {
						sv := ltos[lstruct]
						typ := mapType(fieldTbl, structs)
						metadata := fieldTbl.RawGet(lua.LString("metadata"))
						if metadata == lua.LNil {
							err = fmt.Errorf("field %s is missing metadata", key.String())
							return
						}
						id := metadata.(*lua.LTable).RawGet(lua.LString("id"))
						if id == lua.LNil {
							err = fmt.Errorf("field %s is missing ID in metadata", key.String())
							return
						}
						i, ok := id.(lua.LNumber)
						if !ok {
							err = fmt.Errorf("field %s has invalid ID in metadata", key.String())
							return
						}
						sv.Fields = append(sv.Fields, Field{
							ID:   uint16(i),
							Name: key.String(),
							Type: typ,
						})
					}
				})
			}
		}
	}

	if err != nil {
		return nil, err
	}

	version := 1
	if v := L.GetGlobal("version"); v.Type() == lua.LTNumber {
		version = int(v.(lua.LNumber))
	}

	structList := []StructType{}
	for _, sl := range structs {
		structList = append(structList, *sl)
	}

	return &Schema{name, version, structList}, nil
}

func mapType(tbl *lua.LTable, structs map[string]*StructType) Type {
	switch tbl.RawGet(lua.LString("type")).String() {
	case "primitive":
		return PrimitiveType{Name: tbl.RawGet(lua.LString("name")).String()}
	case "struct":
		return structs[tbl.RawGet(lua.LString("uuid")).String()]
	case "list":
		return ListType{
			ElementType: mapType(tbl.RawGet(lua.LString("of")).(*lua.LTable), structs),
		}
	default:
		panic(fmt.Sprintf("unknown type: %s", tbl.RawGet(lua.LString("type")).String()))
	}
}

func hash(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}
