package request_id

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestId(t *testing.T) {
	cfg := CreateConfig()
	cfg.Header = "some-header"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "test-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeaderNotEmpty(t, req, "some-header")
}

func assertHeaderNotEmpty(t *testing.T, req *http.Request, key string) {
	t.Helper()

	if req.Header.Get(key) == "" {
		t.Errorf("empty header value")
	}
}
