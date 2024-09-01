package item

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
)

func TestPeek(t *testing.T) {
	if os.Getenv("SINGLE_TEST") != "1" {
		t.Skip("skipping test; SINGLE_TEST not set")
	}
	ctx := context.Background()
	cfg, err := config.NewTestConfig(context.Background())
	if err != nil {
		t.Fatalf("newConfig: %s", err)
	}
	cfg.Item.IsEnabled = true

	err = site.Init()
	if err != nil {
		t.Fatalf("site.Init: %s", err)
	}
	err = db.Init(ctx)
	if err != nil {
		t.Fatalf("db.Init: %s", err)
	}
	err = store.Init(ctx)
	if err != nil {
		t.Fatalf("store.Init: %s", err)
	}

	err = Init()
	if err != nil {
		t.Fatalf("Init: %s", err)
	}

	req, err := http.NewRequest("GET", "/peek?id=1001", nil)
	if err != nil {
		t.Fatalf("http.NewRequest: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Peek)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
