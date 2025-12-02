package runner

import (
	"path/filepath"
	"strings"

	"github.com/b3nten/ssss/parser"
)

type SchemaFile struct {
	Schema    *parser.Schema
	Lang      string
	Input     string
	Output    string
	Generated map[string]string
}

func getFileNameFromPath(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}
