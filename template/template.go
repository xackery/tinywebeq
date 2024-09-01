package template

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
)

//go:embed *.go.tmpl layout/*.go.tmpl item/*.go.tmpl player/*.go.tmpl spell/*.go.tmpl npc/*.go.tmpl quest/*.go.tmpl zone/*.go.tmpl style.css
var FS embed.FS

// Compile returns a parsed view with that includes the full site layout.
func Compile(view, file string, f fs.FS) (*template.Template, error) {
	var err error

	tmpl := template.New(view)
	tmpl, err = template.ParseFS(f,
		file,                     // data
		"head.go.tmpl",           // head
		"header.go.tmpl",         // header
		"sidebar.go.tmpl",        // sidebar
		"footer.go.tmpl",         // footer
		"layout/content.go.tmpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		return nil, fmt.Errorf("template.ParseFS: %w", err)
	}
	return tmpl, nil
}
