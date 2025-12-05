package writers

import (
	_ "embed"
)

//go:embed js.lua
var JavascriptTemplate string

//go:embed csharp.lua
var CSharpTemplate string

//go:embed go.lua
var GoTemplate string
