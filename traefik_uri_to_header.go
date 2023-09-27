package traefik_uri_to_header

import (
	"context"
	"net/http"
	"strings"
)

// Config is the configuration for this plugin
type Config struct {
	Enabled    bool   `json:"enabled"`
	HeaderName string `json:"headerName"`
}

// CreateConfig creates a new configuration for this plugin
func CreateConfig() *Config {
	return &Config{}
}

// UriToHeader represents the basic properties of this plugin
type UriToHeader struct {
	next   http.Handler
	name   string
	config *Config
}

// New creates a new instance of this plugin
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &UriToHeader{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (r *UriToHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if r.config.Enabled == true && r.config.HeaderName != "" {
		rawPath := req.URL.RequestURI()
		req.Header.Add(r.config.HeaderName, strings.TrimLeft(rawPath, "/"))
	}
	r.next.ServeHTTP(rw, req)
}
