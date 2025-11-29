package main

import (
	"github.com/b3nten/ssss/writers/go"
	"github.com/b3nten/ssss/parser"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type SchemaFile struct {
	schema *parser.Schema
	lang string
	input string
	output string
	generated map[string]string
}

func main() {
	langFlag := flag.String(
		"lang",
		"",
		"The output language for the generated file (e.g., 'go', 'js'), or a path to a lua template file. (Required)",
	)
	inputFlag := flag.String("i", "", "Path to the input schema file. (Required)")
	outputFlag := flag.String("o", "", "Path to the output directory. (Required)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *inputFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: Input file path is required. Use the -i flag.")
		flag.Usage()
		os.Exit(1)
	}

	if *outputFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: Output directory path is required. Use the -o flag.")
		flag.Usage()
		os.Exit(1)
	}

	if *langFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: Output language is required. Use the -lang flag.")
		flag.Usage()
		os.Exit(1)
	}

	schema, err := parser.GenerateSchema(*inputFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input file '%s': %v\n", *inputFlag, err)
		os.Exit(1)
	}

	file := &SchemaFile{
		schema: schema,
		lang:   *langFlag,
		input:  *inputFlag,
		output: *outputFlag,
		generated: make(map[string]string),
	}

	switch strings.ToLower(*langFlag) {
	case "go": err = generateGoCode(file)
	case "js": err = generateJSCode(file)
	case "c#": err = generateCSharpCode(file)
	default:
		templater, err := os.ReadFile(*langFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading templater file '%s': %v\n", *langFlag, err)
			os.Exit(1)
		}
		err = generateWithLuaTemplate(file, string(templater))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(*outputFlag, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory '%s': %v\n", *outputFlag, err)
		os.Exit(1)
	}

	for filename, generatedCode := range file.generated {
		outputPath := filepath.Join(*outputFlag, filename)
		err = os.WriteFile(outputPath, []byte(generatedCode), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to output file '%s': %v\n", outputPath, err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully generated schema", getFileNameFromPath(file.input), "to", *outputFlag)
}

//go:embed writers/js.lua
var jsTemplater string
//go:embed writers/csharp.lua
var csharpTemplater string

func generateGoCode(file *SchemaFile) error {
	namespace := getFileNameFromPath(file.input)
	code, err := gowriter.Print(file.schema, namespace)
	if err != nil {
		return err
	}
	file.generated[namespace + ".go"] = code
	return nil
}

func generateJSCode(file *SchemaFile) error {
	return generateWithLuaTemplate(file, jsTemplater)
}

func generateCSharpCode(file *SchemaFile) error {
	return generateWithLuaTemplate(file, csharpTemplater)
}

func generateWithLuaTemplate(file *SchemaFile, templater string) error {
	L := parser.CreateLuaState(file.schema)
	defer L.Close()
	if err := L.DoString(templater); err != nil {
		panic(err)
	}
	result := L.GetGlobal("Output")
	if tbl, ok := result.(*lua.LTable); ok {
		tbl.ForEach(func(key lua.LValue, value lua.LValue) {
			file.generated[key.String()] = value.String()
		})
	} else {
	  panic("Output is not a table")
	}
	if len(file.generated) == 0 {
		return fmt.Errorf("No output generated from Lua template")
	}
	return nil
}

func getFileNameFromPath(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}
