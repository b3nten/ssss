package runner

import (
	_ "embed"
	"fmt"

	"github.com/b3nten/ssss/parser"
	writers "github.com/b3nten/ssss/writer"
	gowriter "github.com/b3nten/ssss/writer/go"
	lua "github.com/yuin/gopher-lua"
)

func generateGo(file *SchemaFile) error {
	namespace := getFileNameFromPath(file.Input)
	code, err := gowriter.Print(file.Schema, namespace)
	if err != nil {
		return err
	}
	file.Generated[namespace+".go"] = code
	return nil
}

func generateJS(file *SchemaFile) error {
	return generateCustom(file, writers.JavascriptTemplate)
}

func generateCSharp(file *SchemaFile) error {
	return generateCustom(file, writers.CSharpTemplate)
}

func generateCustom(file *SchemaFile, templater string) error {
	L := parser.CreateLuaState(file.Schema)
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
