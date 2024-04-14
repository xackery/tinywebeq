package template

import "embed"

//go:embed *.go.tpl layout/*.go.tpl item/*.go.tpl player/*.go.tpl
var FS embed.FS
