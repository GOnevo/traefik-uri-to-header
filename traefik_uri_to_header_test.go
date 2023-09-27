package traefik_uri_to_header

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnabled(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "X-Custom-Header-Requested-URI"
	cfg.Enabled = true

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "uri2header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, cfg.HeaderName, "")
}

func TestEnabledLowerCase(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "x-custom-header"
	cfg.Enabled = true

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "uri2header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Custom-Header", "path")
}

func TestEnabledCustomPath(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "X-Custom-Header-Requested-URI"
	cfg.Enabled = true

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "uri2header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/some/beautiful/path?with=query", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, cfg.HeaderName, "some/beautiful/path?with=query")
}

func TestDisabled(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = "X-Custom-Header-Requested-URI"
	cfg.Enabled = false

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "uri2header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, cfg.HeaderName, "")
}

func TestEnabledEmptyHeader(t *testing.T) {
	cfg := CreateConfig()
	cfg.HeaderName = ""
	cfg.Enabled = false

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "uri2header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, cfg.HeaderName, "")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
