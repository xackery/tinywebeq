package template

import "embed"

//go:embed *.go.tpl layout/*.go.tpl item/*.go.tpl player/*.go.tpl spell/*.go.tpl npc/*.go.tpl quest/*.go.tpl zone/*.go.tpl style.css
var FS embed.FS
