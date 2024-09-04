package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

func TestItemPeek(t *testing.T) {
	if os.Getenv("SINGLE_TEST") != "1" {
		t.Skip("skipping test; SINGLE_TEST not set")
	}

	ctx := context.Background()
	cfg, err := config.NewTestConfig(context.Background())
	if err != nil {
		t.Fatalf("newConfig: %s", err)
	}
	cfg.Item.IsEnabled = true

	err = db.Init(ctx)
	if err != nil {
		t.Fatalf("db.Init: %s", err)
	}
	err = store.Init(ctx)
	if err != nil {
		t.Fatalf("store.Init: %s", err)
	}

	req, err := http.NewRequest("GET", "/peek?id=1001", nil)
	if err != nil {
		t.Fatalf("http.NewRequest: %s", err)
	}

	h := New(tlog.Sugar, template.FS)

	rr := httptest.NewRecorder()
	h.PeekItem().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
