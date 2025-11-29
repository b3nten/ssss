package parser

import (
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type Type interface {
	TypeKind() string
}

type PrimitiveType struct {
	Name string
}

func (p PrimitiveType) TypeKind() string { return "primitive" }

type Field struct {
	Name string
	ID   uint16
	Type Type
}

type StructType struct {
	ID     uint16
	UUID   string
	Name   string
	Fields []Field
}

func (s StructType) TypeKind() string { return "struct" }

type ListType struct {
	ElementType Type
}

func (l ListType) TypeKind() string { return "list" }

type Schema struct {
	Name string
	Version int
	Structs []StructType
}

func GenerateSchema(file string) (*Schema, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return &Schema{}, err
	}

	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(prelude + "\n" + string(bytes)); err != nil {
		panic(err)
	}

	lstructs := []*lua.LTable{}
	structs := map[string]*StructType{}
	ltos := map[*lua.LTable]*StructType{}
	if tbl, ok := L.GetGlobal("_G").(*lua.LTable); ok {
		// Iterate over the global table to find structs defined in the global scope
		tbl.ForEach(func(key lua.LValue, value lua.LValue) {
			if structTbl, ok := value.(*lua.LTable); ok {
				if structTbl.RawGet(lua.LString("type")) == lua.LString("struct") {
					sv := StructType{
						Name: key.String(),
						ID:   uint16(fnv32a(key.String())),
						UUID: structTbl.RawGet(lua.LString("uuid")).String(),
					}
					structs[sv.UUID] = &sv
					lstructs = append(lstructs, structTbl)
					ltos[structTbl] = &sv
				}
			}
		})
		// gather all fields in structs
		for _, lstruct := range lstructs {
			if fields, ok := lstruct.RawGet(lua.LString("fields")).(*lua.LTable); ok {
				fields.ForEach(func(key lua.LValue, value lua.LValue) {
					if fieldTbl, ok := value.(*lua.LTable); ok {
						sv := ltos[lstruct]
						typ := mapType(fieldTbl, structs)
						metadata := fieldTbl.RawGet(lua.LString("metadata"))
						if metadata == lua.LNil {
							panic("field " + key.String() + " is missing metadata. You probably forget to set an ID")
						}
						id := metadata.(*lua.LTable).RawGet(lua.LString("id"))
						if id == lua.LNil {
							panic("field " + key.String() + " is missing an ID in metadata")
						}
						i, ok := id.(lua.LNumber)
						if !ok {
							panic("field " + key.String() + " has a non-numeric ID in metadata")
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

	// get version global
	version := 1
	if v := L.GetGlobal("version"); v.Type() == lua.LTNumber {
		version = int(v.(lua.LNumber))
	}

	structList := []StructType{}
	for _, sl := range structs {
		structList = append(structList, *sl)
	}

	return &Schema{getFileNameFromPath(file), version, structList}, nil
}

func generateTypeTable(L *lua.LState, lt Type) *lua.LTable {
	tbl := L.NewTable()
	switch lt.TypeKind() {
	case "primitive":
		pt := lt.(PrimitiveType)
		tbl.RawSet(lua.LString("kind"), lua.LString("primitive"))
		tbl.RawSet(lua.LString("name"), lua.LString(pt.Name))
	case "struct":
		st := lt.(*StructType)
		tbl.RawSet(lua.LString("kind"), lua.LString("struct"))
		tbl.RawSet(lua.LString("name"), lua.LString(st.Name))
		tbl.RawSet(lua.LString("uuid"), lua.LString(st.UUID))
		tbl.RawSet(lua.LString("id"), lua.LNumber(st.ID))
	case "list":
		lt := lt.(ListType)
		tbl.RawSet(lua.LString("kind"), lua.LString("list"))
		tbl.RawSet(lua.LString("of"), generateTypeTable(L, lt.ElementType))
	default:
		panic(fmt.Sprintf("unknown type kind: %T", lt))
	}
	return tbl
}

func CreateLuaState(s *Schema) *lua.LState {
	// create state from schema so lua files can generate code
	L := lua.NewState()
	// set schema properties
	schema := L.NewTable()
	schema.RawSet(lua.LString("name"), lua.LString(s.Name))
	schema.RawSet(lua.LString("version"), lua.LNumber(s.Version))

	structsTable := L.NewTable()

	// set structs
	for _, s := range s.Structs {
		structTable := L.NewTable()

		// set struct properties
		structTable.RawSet(lua.LString("id"), lua.LNumber(s.ID))
		structTable.RawSet(lua.LString("uuid"), lua.LString(s.UUID))
		structTable.RawSet(lua.LString("name"), lua.LString(s.Name))

		// set field properties
		fieldsTable := L.NewTable()
		for _, f := range s.Fields {
			fieldTable := L.NewTable()

			fieldTable.RawSet(lua.LString("name"), lua.LString(f.Name))
			fieldTable.RawSet(lua.LString("id"), lua.LNumber(f.ID))

			// set type properties
			typeTable := generateTypeTable(L, f.Type)
			// assign type to field
			fieldTable.RawSet(lua.LString("type"), typeTable)
			// assign field to fields table
			fieldsTable.RawSet(lua.LString(f.Name), fieldTable)
		}
		// assign fields to struct
		structTable.RawSet(lua.LString("fields"), fieldsTable)
		// assign struct to structs table
		structsTable.RawSet(lua.LString(s.Name), structTable)
	}

	schema.RawSet(lua.LString("structs"), structsTable)
	L.SetGlobal("Schema", schema)
	L.SetGlobal("Output", L.NewTable())
	err := L.DoString(utils)
	if err != nil {
		panic(err)
	}
	return L
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
		return nil
	}
}

func fnv32a(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

func getFileNameFromPath(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

var prelude = `
local function uuid()
    local template ='xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function (c)
        local v = (c == 'x') and math.random(0, 0xf) or math.random(8, 0xb)
        return string.format('%x', v)
    end)
end

-- Helper to make a type callable for metadata
local function make_type(base_type)
    return setmetatable(base_type, {
        __call = function(self, metadata)
            local result = {}
            for k, v in pairs(self) do
                result[k] = v
            end
            result.metadata = metadata
            return result
        end
    })
end

-- Primitives
bool = make_type({ type = "primitive", name = "bool" })
int8 = make_type({ type = "primitive", name = "int8" })
uint8 = make_type({ type = "primitive", name = "uint8" })
int16 = make_type({ type = "primitive", name = "int16" })
uint16 = make_type({ type = "primitive", name = "uint16" })
int32 = make_type({ type = "primitive", name = "int32" })
uint32 = make_type({ type = "primitive", name = "uint32" })
int64 = make_type({ type = "primitive", name = "int64" })
uint64 = make_type({ type = "primitive", name = "uint64" })
str = make_type({ type = "primitive", name = "string" })

function struct(fields)
    local s = {
        type = "struct",
        fields = fields,
        uuid = uuid(),
    }
    return setmetatable(s, {
        __call = function(self, metadata)
            local result = {}
            for k, v in pairs(self) do
                result[k] = v
            end
            result.metadata = metadata
            return result
        end
    })
end

function list(type)
    return make_type({type = "list", of = type})
end

function map(keyType, valueType)
    return make_type({type = "map", key = keyType, value = valueType})
end
`

var utils = `
function sprintf(s, tab)
	return (s:gsub('($%b{})', function(w) return tab[w:sub(3, -2)] or w end))
end
getmetatable("").__mod = sprintf

function unindent(str, args)
	str = sprintf(str, args)
	str = str:gsub("^%s*\n", ""):gsub("\n%s*$", "")
	local min_indent = nil
	local min_indent_len = math.huge
	for line in str:gmatch("[^\n]+") do
		local indent = line:match("^%s*")
		local content = line:gsub("^%s+", "")
		if #content > 0 then         -- ignore empty lines
			if #indent < min_indent_len then
				min_indent_len = #indent
				min_indent = indent
			end
		end
	end
	if min_indent and #min_indent > 0 then
		local pattern = "^" .. min_indent:gsub("([%^%$%(%)%%%.%[%]%*%+%-%?])", "%%%1")
		str = str:gsub("([^\n]+)", function(line)
			return line:gsub(pattern, "", 1)
		end)
	end
	if args.tabs and args.tabs > 0 then
		local indent_str = string.rep("\t", args.tabs)
		str = str:gsub("([^\n]+)", indent_str .. "%1")
	end
	return str
end

function str_block(args)
	if type(args) == "string"
	then
		return unindent(args, {})
	else
		return function(str)
			return unindent(str, args or {})
		end
	end
end

function pascal_case(str)
	local result = str:gsub("[%w]+", function(word)
		return word:sub(1, 1):upper() .. word:sub(2):lower()
	end)
	return (result:gsub("[^%w]", ""))
end

`
