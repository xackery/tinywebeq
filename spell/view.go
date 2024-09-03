package spell

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/site"
)

// View handles spell view requests
func View(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int

		tlog.Debugf("view: %s", r.URL.String())

		strID := r.URL.Query().Get("id")
		if len(strID) > 0 {
			id, err = strconv.Atoi(strID)
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
			if err.Error() == "spell not found" {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}
			tlog.Errorf("viewRender: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("viewRender: id: %d done", id)
	}
}

func viewRender(ctx context.Context, templates fs.FS, id int, w http.ResponseWriter) error {
	if id < 1 {
		return fmt.Errorf("id too low")
	}

	se := store.SpellByID(int32(id))
	if se == nil {
		return fmt.Errorf("spell not found")
	}

	_, info := store.SpellInfo(int32(id), 0)

	data := struct {
		Site                 site.BaseData
		Spell                *models.Spell
		SpellInfo            []string
		IsSpellSearchEnabled bool
	}{
		Site:                 site.BaseDataInit(se.Name),
		Spell:                se,
		SpellInfo:            info,
		IsSpellSearchEnabled: config.Get().Spell.Search.IsEnabled,
	}

	if config.Get().Spell.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/spell/preview.png?id=%d", id)
	}

	view, err := template.Compile("spell", "spell/view.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("viewTemplate.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
