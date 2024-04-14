package site

import (
	"io/fs"
	"os"

	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	templateFS    fs.FS
	isInitialized bool
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true
	fi, err := os.Stat("template")
	if err == nil && fi.IsDir() {
		tlog.Infof("Using local template directory")
		templateFS = os.DirFS("template")
		return nil
	}

	templateFS = template.FS
	return nil
}

// TemplateFS returns the content filesystem
func TemplateFS() fs.FS {
	return templateFS
}
