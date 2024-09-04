package main

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) cssHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := app.templates.Open("style.css")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, "style.css", time.Now(), bytes.NewReader(data))
	}
}

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * time.Second))

	// Basic Routes
	r.Get("/", app.handlers.Root())
	r.Get("/css/style.css", app.cssHandler())

	// Item Routes
	r.Route("/items", func(r chi.Router) {
		r.Get("/", app.handlers.IndexItems())
		r.Route("/{itemID}", func(r chi.Router) {
			r.Use(app.itemCtx)
			r.Get("/", app.handlers.ViewItem())
			r.Get("/peek", app.handlers.PeekItem())
			r.Get("/image", app.handlers.GenerateItemImage())
		})
	})

	// NPC Routes
	r.Route("/npcs", func(r chi.Router) {
		r.Get("/", app.handlers.IndexNpcs())
		r.Route("/{npcID}", func(r chi.Router) {
			r.Use(npcCtx)
			r.Get("/", app.handlers.ViewNpc())
			r.Get("/peek", app.handlers.PeekNpc())
			r.Get("/image", app.handlers.GenerateNpcImage())
		})
	})

	// Player Routes
	r.Route("/players", func(r chi.Router) {
		r.Route("/{playerID}", func(r chi.Router) {
			r.Use(playerCtx)
			r.Get("/", app.handlers.ViewPlayer())
		})
	})

	// Quest Routes
	r.Route("/quests", func(r chi.Router) {
		r.Get("/", app.handlers.IndexQuests())
		r.Route("/{questID}", func(r chi.Router) {
			r.Use(questCtx)
			r.Get("/", app.handlers.ViewQuest())
			r.Get("/image", app.handlers.GenerateQuestImage())
		})
	})

	// Spell Routes
	r.Route("/spells", func(r chi.Router) {
		r.Get("/", app.handlers.IndexSpells())
		r.Route("/{spellID}", func(r chi.Router) {
			r.Use(spellCtx)
			r.Get("/", app.handlers.ViewSpell())
			r.Get("/image", app.handlers.GenerateSpellImage())
		})
	})

	// Zone Routes
	r.Route("/zones", func(r chi.Router) {
		r.Get("/", app.handlers.IndexZones())
		r.Route("/{zoneID}", func(r chi.Router) {
			r.Use(zoneCtx)
			r.Get("/view", app.handlers.ViewZone())
			r.Get("/image", app.handlers.GenerateZoneImage())
		})
	})

	return r
}
