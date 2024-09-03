package handlers

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

type ContextKey string

const (
	ContextKeyItem ContextKey = "item"
)

type Logger interface {
	Info(args ...any)
	Error(args ...any)
	Debug(args ...any)
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

func (h *Handlers) render(w http.ResponseWriter, name, fileName, layoutName string, data any) {
	path := fmt.Sprintf("%s/%s", name, fileName)

	view, err := template.Compile(name, path, h.templates)
	if err != nil {
		h.logger.Error("error compiling templates", zap.Error(err))
		h.serverErrorResponse(w, fmt.Errorf("template.Compile: %w", err))
		return
	}

	if err = view.ExecuteTemplate(w, layoutName, data); err != nil {
		h.logger.Error("viewTemplate.Execute", zap.Error(err))
		return
	}
}
