package handlers

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/xackery/tinywebeq/tlog"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
}

type Handlers struct {
	logger    Logger
	templates fs.FS
}

func New(logger Logger, templates fs.FS) *Handlers {
	return &Handlers{
		logger:    logger,
		templates: templates,
	}
}

func viewJSON(w http.ResponseWriter, headers http.Header, data any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		tlog.Errorf("json.Encode: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
