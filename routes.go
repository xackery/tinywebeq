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
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/spell"
	"github.com/xackery/tinywebeq/zone"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, world!")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("GET /items/view/{itemID}", item.View)
	mux.HandleFunc("GET /items/peek", item.Peek)
	mux.HandleFunc("GET /items/search", item.Search)
	mux.HandleFunc("GET /items/preview.png", item.PreviewImage)
	mux.HandleFunc("/player/view/", player.View)
	mux.HandleFunc("/spell/view", spell.View)
	mux.HandleFunc("/spell/search", spell.Search)
	mux.HandleFunc("/spell/preview.png", spell.PreviewImage)
	mux.HandleFunc("/npc/view/", npc.View)
	mux.HandleFunc("/npc/peek/", npc.Peek)
	mux.HandleFunc("/npc/search", npc.Search)
	mux.HandleFunc("/npc/preview.png", npc.PreviewImage)
	mux.HandleFunc("/quest/view/", quest.View)
	mux.HandleFunc("/quest/search", quest.Search)
	mux.HandleFunc("/quest/preview.png", quest.PreviewImage)
	// mux.HandleFunc("/recipe/view/", recipe.View)
	// mux.HandleFunc("/recipe/search", recipe.Search)
	// mux.HandleFunc("/recipe/preview.png", recipe.PreviewImage)
	mux.HandleFunc("/zone/view/", zone.View)
	mux.HandleFunc("/zone/search", zone.Search)
	mux.HandleFunc("/zone/preview.png", zone.PreviewImage)
	mux.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		fi, err := site.TemplateFS().Open("style.css")
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

	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	return mux
}
