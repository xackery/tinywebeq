package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/item"
	"github.com/xackery/tinywebeq/player"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/tlog"
	"github.com/xackery/tinywebeq/util"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Get().Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	var err error

	err = site.Init()
	if err != nil {
		return fmt.Errorf("site.Init: %w", err)
	}
	err = cache.Init()
	if err != nil {
		return fmt.Errorf("cache.Init: %w", err)
	}

	err = item.Init()
	if err != nil {
		return fmt.Errorf("item.Init: %w", err)
	}
	err = player.Init()
	if err != nil {
		return fmt.Errorf("player.Init: %w", err)
	}

	_, err = config.NewConfig(ctx)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %w", err)
	}

	err = db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}

	err = util.Init()
	if err != nil {
		return fmt.Errorf("util.Init: %w", err)
	}
	err = os.MkdirAll("cache", 0755)
	if err != nil {
		return fmt.Errorf("make cache: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/item/view/", item.View)
	mux.HandleFunc("/item/preview.png", item.PreviewImage)
	mux.HandleFunc("/player/view/", player.View)
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	tlog.Infof("Listening on http://127.0.0.1:8080")
	return http.ListenAndServe(":8080", mux)
}
