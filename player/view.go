package player

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

// View handles player view requests
func View(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int64

		tlog.Debugf("view: %s", r.URL.String())

		strID := r.URL.Query().Get("id")
		if len(strID) > 0 {
			id, err = strconv.ParseInt(strID, 10, 64)
			if err != nil {
				tlog.Errorf("strconv.Atoi: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tlog.Debugf("viewRender: id: %d", id)

		err = viewRender(ctx, templates, id, w)
		if err != nil {
			tlog.Errorf("viewRender: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("viewRender: id: %d done", id)
	}
}

func viewRender(ctx context.Context, templates fs.FS, id int64, w http.ResponseWriter) error {
	if id < 1 {
		return fmt.Errorf("id too low")
	}

	player, err := store.PlayerByCharacterID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.PlayerByCharacterID: %w", err)
	}

	type TemplateData struct {
		Site   site.BaseData
		Player *model.Player
	}

	data := TemplateData{
		Site:   site.BaseDataInit("Player View"),
		Player: player,
	}

	view, err := template.Compile("player", "player/view.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("template.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
