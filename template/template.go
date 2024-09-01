package template

import "embed"

//go:embed *.go.tmpl layout/*.go.tmpl item/*.go.tmpl player/*.go.tmpl spell/*.go.tmpl npc/*.go.tmpl quest/*.go.tmpl zone/*.go.tmpl style.css
var FS embed.FS
