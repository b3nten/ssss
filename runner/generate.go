package runner

import (
	_ "embed"
	"fmt"

	"github.com/b3nten/ssss/parser"
	writers "github.com/b3nten/ssss/writer"
	lua "github.com/yuin/gopher-lua"
)

func generateGo(file *SchemaFile, debug bool) error {
	return generateCustom(file, writers.GoTemplate, debug)
}

func generateJS(file *SchemaFile, debug bool) error {
	return generateCustom(file, writers.JavascriptTemplate, debug)
}

func generateCSharp(file *SchemaFile, debug bool) error {
	return generateCustom(file, writers.CSharpTemplate, debug)
}

func generateCustom(file *SchemaFile, templater string, debug bool) error {
	L := parser.CreateLuaState(file.Schema, debug)
	defer L.Close()
	if err := L.DoString(templater); err != nil {
		panic(err)
	}
	result := L.GetGlobal("Output")
	if tbl, ok := result.(*lua.LTable); ok {
		tbl.ForEach(func(key lua.LValue, value lua.LValue) {
			file.Generated[key.String()] = value.String()
		})
	} else {
		panic("Output is not a table")
	}
	if len(file.Generated) == 0 {
		return fmt.Errorf("No output generated from Lua template")
	}
	return nil
}
