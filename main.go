package main

import (
	"context"
	"os"

	runner "github.com/b3nten/ssss/runner"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "stupidly-simple-serialization-system",
		Description: "ssss is a simple binary de/serialization library which generates single file, zero dependency artifacts.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Usage:    "path to input schema file",
				Required: true,
				Aliases:  []string{"i", "file", "in"},
			},
			&cli.StringFlag{
				Name:     "output",
				Usage:    "directory for generated code output",
				Required: true,
				Aliases:  []string{"o", "out"},
			},
			&cli.StringFlag{
				Name:     "lang",
				Usage:    "target programming language or path to custom templater file",
				Required: true,
				Aliases:  []string{"l", "generator"},
			},
		},
		Action: runner.Run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)
		boldRed.Println(err.Error())
		os.Exit(1)
	}
}
