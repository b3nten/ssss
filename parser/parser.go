package parser

import (
	"fmt"
	"hash/fnv"

	_ "embed"

	lua "github.com/yuin/gopher-lua"
)

//go:embed schema_globals.lua
var schemaGlobals string

func GenerateSchema(file, name string, debug bool) (*Schema, error) {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(schemaGlobals); err != nil {
		return nil, fmt.Errorf("failed to load schema globals: %w", err)
	}

	if err := L.DoString(string(file)); err != nil {
		return nil, fmt.Errorf("failed to execute schema file: %w", err)
	}

	structs := map[string]StructType{}
	seenStructs := map[uint16]bool{}
	lstructs := []*lua.LTable{}

	if globalTable, ok := L.GetGlobal("__Structs").(*lua.LTable); ok {
		// iterate all globals
		globalTable.ForEach(func(key lua.LValue, value lua.LValue) {
			if structTbl, ok := value.(*lua.LTable); ok {
				if structTbl.RawGet(lua.LString("type")) == lua.LString("struct") {
					lstructs = append(lstructs, structTbl)
				}
			}
		})
	} else {
		return nil, fmt.Errorf("failed to get global table")
	}

	for _, structTbl := range lstructs {

		id, name, err := extractStructMetadata(structTbl)

		if debug {
			fmt.Printf("Processing struct: %s (ID: %d)\n", name, id)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to extract metadata for struct: %w", err)
		}

		if seenStructs[id] {
			return nil, fmt.Errorf("hash collision for struct name: %s. Please change the name or set an ID manually.", name)
		} else {
			seenStructs[id] = true
		}

		sv := StructType{
			Name: name,
			ID:   id,
		}

		structs[sv.Name] = sv
	}

	for _, structTbl := range lstructs {
		seenFields := map[uint16]bool{}
		fields := map[string]*lua.LTable{}
		_, name, _ := extractStructMetadata(structTbl)
		sv := structs[name]

		if fieldTable, ok := structTbl.RawGet(lua.LString("fields")).(*lua.LTable); ok {
			fieldTable.ForEach(func(key lua.LValue, value lua.LValue) {
				if fieldTbl, ok := value.(*lua.LTable); ok {
					fields[key.String()] = fieldTbl
				}
			})
		} else {
			return nil, fmt.Errorf("struct %s is missing fields table", sv.Name)
		}

		for name, fieldTbl := range fields {
			typ, err := mapType(fieldTbl, structs, debug)

			if err != nil {
				return nil, fmt.Errorf("failed to map type for field %s in struct %s: %w", name, sv.Name, err)
			}

			metadata := fieldTbl.RawGet(lua.LString("metadata"))

			if metadata == lua.LNil {
				return nil, fmt.Errorf("field %s in struct %s is missing metadata", name, sv.Name)
			}

			id := metadata.(*lua.LTable).RawGet(lua.LString("id"))
			if id == lua.LNil {
				return nil, fmt.Errorf("field %s in struct %s is missing ID in metadata", name, sv.Name)
			}

			i, ok := id.(lua.LNumber)
			if !ok {
				return nil, fmt.Errorf("field %s in struct %s has invalid ID in metadata", name, sv.Name)
			}

			if seenFields[uint16(i)] {
				return nil, fmt.Errorf("duplicate field ID %d in struct %s", uint16(i), sv.Name)
			}

			seenFields[uint16(i)] = true

			sv.Fields = append(sv.Fields, Field{
				ID:   uint16(i),
				Name: name,
				Type: typ,
			})
		}

		structs[sv.Name] = sv
	}

	version := 1
	if v := L.GetGlobal("version"); v.Type() == lua.LTNumber {
		version = int(v.(lua.LNumber))
	}

	structList := []StructType{}
	for _, strct := range structs {
		structList = append(structList, strct)
	}

	return &Schema{name, version, structList}, nil
}

func mapType(tbl *lua.LTable, structs map[string]StructType, debug bool) (Type, error) {
	switch tbl.RawGet(lua.LString("type")).String() {

	case "primitive":
		if debug {
			fmt.Printf("    Primitive type: %s\n", tbl.RawGet(lua.LString("name")).String())
		}
		return PrimitiveType{Name: tbl.RawGet(lua.LString("name")).String()}, nil

	case "struct":
		if debug {
			fmt.Printf("    Struct type: %s\n", tbl.RawGet(lua.LString("name")).String())
		}
		_, name, err := extractStructMetadata(tbl)
		if err != nil {
			return nil, err
		}
		strct, ok := structs[name]
		if !ok {
			return nil, fmt.Errorf("unknown struct type: %s", name)
		}
		return strct, nil

	case "list":
		if debug {
			fmt.Printf("    List type\n")
		}

		rawOf := tbl.RawGet(lua.LString("of"))
		if rawOf.Type() == lua.LNil.Type() {
			panic(rawOf.String())
		}

		of, err := mapType(tbl.RawGet(lua.LString("of")).(*lua.LTable), structs, debug)
		if err != nil {
			return nil, err
		}
		return ListType{
			ElementType: of,
		}, nil

	case "map":
		if debug {
			fmt.Printf("    Map type\n")
		}
		key, err := mapType(tbl.RawGet(lua.LString("from")).(*lua.LTable), structs, debug)
		if err != nil {
			return nil, err
		}
		value, err := mapType(tbl.RawGet(lua.LString("to")).(*lua.LTable), structs, debug)
		if err != nil {
			return nil, err
		}
		return MapType{
			KeyType:   key,
			ValueType: value,
		}, nil

	default:
		return nil, fmt.Errorf("unknown type: %s", tbl.RawGet(lua.LString("type")).String())
	}
}

func hash(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

func extractStructMetadata(tbl *lua.LTable) (uint16, string, error) {
	metadataLValue := tbl.RawGet(lua.LString("metadata"))
	if metadataLValue == lua.LNil {
		return 0, "", fmt.Errorf("struct has invalid metadata")
	}

	metadata, ok := metadataLValue.(*lua.LTable)
	if !ok {
		return 0, "", fmt.Errorf("struct has invalid metadata")
	}

	nameLValue := metadata.RawGet(lua.LString("name"))
	if nameLValue == lua.LNil {
		return 0, "", fmt.Errorf("struct is missing name in metadata")
	}

	nameLStr, ok := nameLValue.(lua.LString)
	if !ok {
		return 0, "", fmt.Errorf("struct has invalid name in metadata")
	}

	name := nameLStr.String()

	if name == "" {
		return 0, "", fmt.Errorf("struct is missing name in metadata")
	}

	var id uint16
	idLValue := metadata.RawGet(lua.LString("id"))
	if idLValue == lua.LNil {
		id = uint16(hash(name) % 65536)
	} else {
		idNumber, ok := idLValue.(lua.LNumber)
		if !ok {
			return 0, "", fmt.Errorf("struct has invalid ID in metadata")
		}
		id = uint16(idNumber)
	}

	return id, name, nil
}
