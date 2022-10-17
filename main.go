package request_id_traefik_plugin

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
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
	if len(config.Header) == 0 {
		return nil, errors.New("invalid header name")
	}

	return &RequestId{
		next:   next,
		name:   name,
		header: config.Header,
	}, nil
}

func (r *RequestId) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	request_id := uuid.New().String()
	req.Header.Set(r.header, string(request_id))
	r.next.ServeHTTP(rw, req)
}
