package spell

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/site"
)

var (
	viewTemplate *template.Template
)

func viewInit() error {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"spell/view.go.tpl",     // data
		"head.go.tpl",           // head
		"header.go.tpl",         // header
		"footer.go.tpl",         // footer
		"layout/content.go.tpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}
	return nil
}

// View handles spell view requests
func View(w http.ResponseWriter, r *http.Request) {
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

	err = viewRender(ctx, id, w)
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

func viewRender(ctx context.Context, id int, w http.ResponseWriter) error {
	if id < 1 {
		return fmt.Errorf("id too low")
	}

	se := library.SpellByID(id)
	if se == nil {
		return fmt.Errorf("spell not found")
	}

	type TemplateData struct {
		Site                 site.BaseData
		Spell                *library.Spell
		SpellInfo            []string
		IsSpellSearchEnabled bool
	}

	data := TemplateData{
		Site:                 site.BaseDataInit("Spell View"),
		Spell:                se,
		SpellInfo:            library.SpellInfo(id),
		IsSpellSearchEnabled: config.Get().Spell.Search.IsEnabled,
	}

	err := viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
