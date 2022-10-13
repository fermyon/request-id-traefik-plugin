// Package example a example plugin.
package main

import (
	"context"
	"net/http"
)

var defaultHeader string = "X-Request-Id"

// Config the plugin configuration.
type Config struct {
	Header string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Header: defaultHeader,
	}
}

// Example a plugin.
type RequestId struct {
	next   http.Handler
	name   string
	header string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// TODO make sure header is valid
	return &RequestId{
		next:   next,
		name:   name,
		header: config.Header,
	}, nil
}

func (r *RequestId) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.Header.Set(r.header, "something")
	r.next.ServeHTTP(rw, req)
}
