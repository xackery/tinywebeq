package zone

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

// View handles zone view rezones
func View(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var err error
		var id int

		if !config.Get().Zone.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

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
			if err.Error() == "zone not found" {
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
	zone, err := store.ZoneByZoneIDNumber(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("store.ZoneByZoneIDNumber: %w", err)
	}

	data := struct {
		Site                site.BaseData
		Zone                *model.Zone
		Library             *library.Library
		ZoneInfo            []string
		IsZoneSearchEnabled bool
	}{
		Site:                site.BaseDataInit(zone.LongName),
		Zone:                zone,
		Library:             library.Instance(),
		IsZoneSearchEnabled: config.Get().Zone.Search.IsEnabled,
	}

	if config.Get().Zone.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/zone/preview.png?id=%d", id)
	}

	view, err := template.Compile("zone", "zone/view.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("template.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
