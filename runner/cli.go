package runner

import (
	"context"
	"fmt"

	_ "embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/b3nten/ssss/parser"
	"github.com/fatih/color"

	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, cli *cli.Command) error {
	input := cli.String("input")
	output := cli.String("output")
	lang := cli.String("lang")
	debug := cli.Bool("debug")

	if input == "" || output == "" || lang == "" {
		return fmt.Errorf("input, output, and lang parameters are required")
	}

	inputFile, err := os.ReadFile(input)
	if err != nil {
		return fmt.Errorf("error reading schema file '%s': %v", input, err)
	}

	if debug {
		fmt.Println(color.YellowString("Debug mode enabled"))
	}

	schema, err := parser.GenerateSchema(string(inputFile), getFileNameFromPath(input), debug)

	if err != nil {
		return fmt.Errorf("error parsing schema file '%s': %v", input, err)
	}

	file := &SchemaFile{
		Schema:    schema,
		Lang:      lang,
		Input:     input,
		Output:    output,
		Generated: make(map[string]string),
	}

	switch strings.ToLower(file.Lang) {
	case "go":
		err = generateGo(file, debug)
	case "js":
		err = generateJS(file, debug)
	case "c#":
		err = generateCSharp(file, debug)
	default:
		codeGenScript, err := os.ReadFile(file.Lang)
		if err != nil {
			return fmt.Errorf("Error reading templater file '%s': %v\n", file.Lang, err)
		}
		err = generateCustom(file, string(codeGenScript), debug)
	}

	if err != nil {
		return fmt.Errorf("Error generating code: %v\n", err)
	}

	if err := os.MkdirAll(file.Output, 0755); err != nil {
		return fmt.Errorf("Error creating output directory '%s': %v\n", file.Output, err)
	}

	for filename, generatedCode := range file.Generated {
		outputPath := filepath.Join(file.Output, filename)
		err = os.WriteFile(outputPath, []byte(generatedCode), 0644)
		if err != nil {
			return fmt.Errorf("Error writing to output file '%s': %v\n", outputPath, err)
		}
	}

	if debug {
		fmt.Println(color.YellowString("Generated files:"))
		for filename := range file.Generated {
			fmt.Println(color.YellowString(" - %s", filename))
		}
	}

	green := color.GreenString("Successfully generated schema `%s` to dir `%s` for %s", getFileNameFromPath(file.Input), file.Output, file.Lang)
	fmt.Println(green)

	return nil
}
