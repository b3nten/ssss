package parser

import (
	_ "embed"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

//go:embed gen_utils.lua
var customUtils string

func CreateLuaState(s *Schema, debug bool) *lua.LState {
	L := lua.NewState()

	schema := L.NewTable()
	schema.RawSet(lua.LString("name"), lua.LString(s.Name))
	schema.RawSet(lua.LString("version"), lua.LNumber(s.Version))

	if debug {
		fmt.Printf("Creating Lua state for schema: %s (version %d)\n", s.Name, s.Version)
	}

	tAllStructs := L.NewTable()

	for _, s := range s.Structs {
		tStruct := L.NewTable()

		tStruct.RawSet(lua.LString("id"), lua.LNumber(s.ID))
		tStruct.RawSet(lua.LString("name"), lua.LString(s.Name))

		if debug {
			fmt.Printf("Processing struct: %s (ID: %d)\n", s.Name, s.ID)
		}

		tFields := L.NewTable()
		for _, f := range s.Fields {
			fieldTable := L.NewTable()

			fieldTable.RawSet(lua.LString("name"), lua.LString(f.Name))
			fieldTable.RawSet(lua.LString("id"), lua.LNumber(f.ID))

			if debug {
				fmt.Printf("  Processing field: %s (ID: %d)\n", f.Name, f.ID)
			}

			typeTable := generateTypeTable(L, f.Type, debug)
			fieldTable.RawSet(lua.LString("type"), typeTable)

			tFields.RawSet(lua.LString(f.Name), fieldTable)
		}
		tStruct.RawSet(lua.LString("fields"), tFields)
		tAllStructs.RawSet(lua.LString(s.Name), tStruct)
	}

	schema.RawSet(lua.LString("structs"), tAllStructs)
	L.SetGlobal("Schema", schema)
	L.SetGlobal("Output", L.NewTable())

	err := L.DoString(customUtils)

	if err != nil {
		panic(err)
	}

	return L
}

func generateTypeTable(L *lua.LState, lt Type, debug bool) *lua.LTable {
	if debug {
		fmt.Printf("Generating type table for type kind: %s\n", lt.TypeKind())
	}
	tbl := L.NewTable()
	switch lt.TypeKind() {
	case "primitive":
		pt := lt.(PrimitiveType)
		tbl.RawSet(lua.LString("kind"), lua.LString("primitive"))
		tbl.RawSet(lua.LString("name"), lua.LString(pt.Name))
	case "struct":
		st := lt.(StructType)
		tbl.RawSet(lua.LString("kind"), lua.LString("struct"))
		tbl.RawSet(lua.LString("name"), lua.LString(st.Name))
		tbl.RawSet(lua.LString("id"), lua.LNumber(st.ID))
	case "list":
		lt := lt.(ListType)
		tbl.RawSet(lua.LString("kind"), lua.LString("list"))
		tbl.RawSet(lua.LString("of"), generateTypeTable(L, lt.ElementType, debug))
	case "map":
		mt := lt.(MapType)
		tbl.RawSet(lua.LString("kind"), lua.LString("map"))
		tbl.RawSet(lua.LString("from"), generateTypeTable(L, mt.KeyType, debug))
		tbl.RawSet(lua.LString("to"), generateTypeTable(L, mt.ValueType, debug))
	default:
		panic(fmt.Sprintf("unknown type kind: %T", lt))
	}
	return tbl
}
