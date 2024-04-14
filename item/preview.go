package item

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/site"
)

var (
	previewTemplate *template.Template
)

// https://allaclone.wayfarershaven.com/?a=item&id=1004
func previewInit() error {
	var err error
	previewTemplate = template.New("preview")
	previewTemplate, err = previewTemplate.ParseFS(site.TemplateFS(),
		"item/preview.go.tpl",   // preview
		"layout/preview.go.tpl", // layout (requires preview)
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}
	return nil
}

// Preview handles item preview requests
func Preview(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("preview: %s", r.URL.String())

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

	tlog.Debugf("previewRender: id: %d", id)

	err = previewRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("previewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("previewRender: id: %d done", id)
}

func previewRender(ctx context.Context, id int, w http.ResponseWriter) error {
	if id <= 1000 {
		return fmt.Errorf("invalid id")
	}

	path := fmt.Sprintf("item/preview_%d.html", id)

	ok, cacheData := cache.Read(path)
	if ok {
		_, err := w.Write(cacheData)
		if err != nil {
			return fmt.Errorf("from cache write: %w", err)
		}
		return nil
	}

	query := "SELECT * FROM items WHERE id=:id LIMIT 1"
	if config.Get().IsDiscoveredOnly {
		query += "SELECT * FROM items, discovered items WHERE items.id=:id AND discovered_items.item_id=:id LIMIT 1"
	}

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return fmt.Errorf("query items: %w", err)
	}
	defer rows.Close()

	type TemplateData struct {
		Item *Table
	}
	item := &Table{}

	data := TemplateData{}

	if !rows.Next() {
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	err = rows.StructScan(item)
	if err != nil {
		return fmt.Errorf("rows.StructScan: %w", err)
	}
	data.Item = item

	buf := &bytes.Buffer{}

	err = previewTemplate.ExecuteTemplate(buf, "preview.go.tpl", data)
	if err != nil {
		return fmt.Errorf("previewTemplate.Execute: %w", err)
	}

	err = cache.Write(path, buf.Bytes())
	if err != nil {
		return fmt.Errorf("cache write: %w", err)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
