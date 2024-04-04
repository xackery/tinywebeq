package site

import (
	"io/fs"
	"os"
)

var (
	templateFS fs.FS
)

func init() {
	templateFS = os.DirFS(".")
}

// TemplateFS returns the content filesystem
func TemplateFS() fs.FS {
	return templateFS
}
