package main

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/xackery/tinywebeq/npc"
	"github.com/xackery/tinywebeq/player"
	"github.com/xackery/tinywebeq/quest"
	"github.com/xackery/tinywebeq/spell"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/zone"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * time.Second))

	// Index Route
	r.Get("/", app.handlers.Root())

	// Item Routes
	r.Route("/items", func(r chi.Router) {
		r.Get("/", app.handlers.ItemIndex())
		r.Route("/{itemID}", func(r chi.Router) {
			r.Use(app.itemContext)
			r.Get("/", app.handlers.ItemView())
			r.Get("/peek", app.handlers.ItemPeek())
			r.Get("/image", app.handlers.ItemImage())
		})
	})

	// Player Routes
	r.Route("/players", func(r chi.Router) {
		r.Get("/", player.View(template.FS))
	})

	// Spell Routes
	r.Route("/spells", func(r chi.Router) {
		r.Get("/view", spell.View(template.FS))
		r.Get("/search", spell.Search)
		r.Get("/preview.png", spell.PreviewImage)
	})

	// NPC Routes
	r.Route("/npcs", func(r chi.Router) {
		r.Get("/", app.handlers.NPCIndex())
		r.Route("/{npcID}", func(r chi.Router) {
			r.Use(app.npcContext)
			r.Get("/", app.handlers.NPCView())
			r.Get("/peek", npc.Peek(template.FS))
			r.Get("/search", npc.Search)
			r.Get("/preview.png", npc.PreviewImage)
		})
	})

	// Quest Routes
	r.Route("/quests", func(r chi.Router) {
		r.Get("/view", quest.View(template.FS))
		r.Get("/search", quest.Search)
		r.Get("/preview.png", quest.PreviewImage)
	})

	// Zone Routes
	r.Route("/zones", func(r chi.Router) {
		r.Get("/view", zone.View(template.FS))
		r.Get("/search", zone.Search)
		r.Get("/preview.png", zone.PreviewImage)
	})

	r.Get("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
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

	return r
}
