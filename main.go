package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/item"
	"github.com/xackery/tinywebeq/player"
	"github.com/xackery/tinywebeq/tlog"
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

	_, err := config.NewConfig(ctx)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %w", err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Get().Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	err = db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
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
	mux.HandleFunc("/item/preview/", item.Preview)
	mux.HandleFunc("/player/view/", player.View)
	tlog.Info("Listening on http://127.0.0.1:8080")
	return http.ListenAndServe(":8080", mux)
}
