package player

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	viewTemplate  *template.Template
	isInitialized bool
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"player/view.go.tpl",    // data
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

// View handles player view requests
func View(w http.ResponseWriter, r *http.Request) {
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

	err = viewRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("viewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("viewRender: id: %d done", id)
}

func viewRender(ctx context.Context, id int64, w http.ResponseWriter) error {
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

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
