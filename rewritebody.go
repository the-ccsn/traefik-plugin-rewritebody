// Package rewrite_body a plugin to rewrite response body.
package traefik_plugin_rewritebody

import (
	"context"
	"net/http"

	"github.com/the-ccsn/traefik-plugin-rewritebody/handler"
	"github.com/the-ccsn/traefik-plugin-rewritebody/httputil"
)

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *handler.Config {
	return &handler.Config{
		LastModified: false,
		Rewrites:     nil,
		LogLevel:     0,
		Monitoring:   *httputil.CreateMonitoringConfig(),
	}
}

// New creates and returns a new rewrite body plugin instance.
func New(context context.Context, next http.Handler, config *handler.Config, name string) (http.Handler, error) {
	return handler.New(context, next, config, name)
}
