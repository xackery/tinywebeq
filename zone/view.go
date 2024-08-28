package zone

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	viewTemplate *template.Template
)

func viewInit() error {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"zone/view.go.tpl",      // data
		"head.go.tpl",           // head
		"header.go.tpl",         // header
		"sidebar.go.tpl",        // sidebar
		"footer.go.tpl",         // footer
		"layout/content.go.tpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}

	return nil
}

// View handles zone view rezones
func View(w http.ResponseWriter, r *http.Request) {
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

	err = viewRender(ctx, id, w)
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

func viewRender(ctx context.Context, id int, w http.ResponseWriter) error {
	zone, err := store.ZoneByZoneIDNumber(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("store.ZoneByZoneIDNumber: %w", err)
	}

	type TemplateData struct {
		Site                site.BaseData
		Zone                *model.Zone
		Library             *library.Library
		ZoneInfo            []string
		IsZoneSearchEnabled bool
	}

	data := TemplateData{
		Site:                site.BaseDataInit(zone.LongName),
		Zone:                zone,
		Library:             library.Instance(),
		IsZoneSearchEnabled: config.Get().Zone.Search.IsEnabled,
	}
	if config.Get().Zone.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/zone/preview.png?id=%d", id)
	}

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
