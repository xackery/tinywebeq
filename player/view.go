package player

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/site"
)

var (
	viewTemplate *template.Template
)

func init() {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"template/player/view.go.tpl",    // data
		"template/head.go.tpl",           // head
		"template/header.go.tpl",         // header
		"template/footer.go.tpl",         // footer
		"template/layout/content.go.tpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		tlog.Fatalf("template.ParseFS: %v", err)
		return
	}
}

// View handles player view requests
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
		tlog.Errorf("viewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("viewRender: id: %d done", id)
}

func viewRender(ctx context.Context, id int, w http.ResponseWriter) error {

	path := fmt.Sprintf("player/view_%d.html", id)

	ok, cacheData := cache.Read(path)
	if ok {
		_, err := w.Write(cacheData)
		if err != nil {
			return fmt.Errorf("from cache write: %w", err)
		}
		return nil
	}

	query := "SELECT * FROM character_data WHERE id=:id LIMIT 1"
	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return fmt.Errorf("query players: %w", err)
	}
	defer rows.Close()

	type TemplateData struct {
		Site   site.BaseData
		Player *Table
	}

	data := TemplateData{Site: site.BaseDataInit("Player View")}
	player := &Table{}

	if !rows.Next() {
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	err = rows.StructScan(player)
	if err != nil {
		return fmt.Errorf("rows.StructScan: %w", err)
	}
	data.Player = player

	buf := &bytes.Buffer{}

	err = viewTemplate.ExecuteTemplate(buf, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	err = cache.Write(path, buf.Bytes())
	if err != nil {
		return fmt.Errorf("cache.Write: %w", err)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
