package main

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/xackery/tinywebeq/item"
	"github.com/xackery/tinywebeq/npc"
	"github.com/xackery/tinywebeq/player"
	"github.com/xackery/tinywebeq/quest"
	"github.com/xackery/tinywebeq/spell"
	"github.com/xackery/tinywebeq/zone"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.handlers.Root())
	mux.HandleFunc("GET /items", app.handlers.ItemIndex())
	mux.HandleFunc("GET /items/{itemID}", app.handlers.ItemView())
	mux.HandleFunc("GET /items/{itemID}/peek", item.Peek(app.templates))
	mux.HandleFunc("GET /items/search", item.Search)
	mux.HandleFunc("GET /items/preview.png", item.PreviewImage)
	mux.HandleFunc("GET /player/view/", player.View(app.templates))
	mux.HandleFunc("GET /spell/view", spell.View(app.templates))
	mux.HandleFunc("GET /spell/search", spell.Search)
	mux.HandleFunc("GET /spell/preview.png", spell.PreviewImage)
	mux.HandleFunc("GET /npcs/view/", npc.View(app.templates))
	mux.HandleFunc("GET /npcs/peek/", npc.Peek(app.templates))
	mux.HandleFunc("GET /npcs/search", npc.Search)
	mux.HandleFunc("GET /npcs/preview.png", npc.PreviewImage)
	mux.HandleFunc("GET /quest/view/", quest.View(app.templates))
	mux.HandleFunc("GET /quest/search", quest.Search)
	mux.HandleFunc("GET /quest/preview.png", quest.PreviewImage)
	// mux.HandleFunc("/recipe/view/", recipe.View)
	// mux.HandleFunc("/recipe/search", recipe.Search)
	// mux.HandleFunc("/recipe/preview.png", recipe.PreviewImage)
	mux.HandleFunc("GET /zone/view/", zone.View(app.templates))
	mux.HandleFunc("GET /zone/search", zone.Search)
	mux.HandleFunc("GET /zone/preview.png", zone.PreviewImage)
	mux.HandleFunc("GET /css/style.css", func(w http.ResponseWriter, r *http.Request) {
		fi, err := app.templates.Open("style.css")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer fi.Close()
		data, err := io.ReadAll(fi)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, "style.css", time.Now(), bytes.NewReader(data))
	})

	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	return mux
}
