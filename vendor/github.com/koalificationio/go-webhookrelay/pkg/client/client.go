// Package client provides convenience wrapper for creating WebhookRelay OpenAPI client
package client

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/koalificationio/go-webhookrelay/pkg/openapi/client"
)

// Config settings for client
type Config struct {
	APIKey    string
	APISecret string

	// Override default host
	Host string

	// Override default base path
	BasePath string

	// Override default schemes
	Schemes []string
}

// New creates new webhookrelay client using defaults and overrides
func New(cfg *Config) *apiclient.Openapi {
	if cfg.Host == "" {
		cfg.Host = apiclient.DefaultHost
	}

	if cfg.BasePath == "" {
		cfg.BasePath = apiclient.DefaultBasePath
	}

	if cfg.Host == "" {
		cfg.Schemes = apiclient.DefaultSchemes
	}

	runtime := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	runtime.DefaultAuthentication = httptransport.BasicAuth(cfg.APIKey, cfg.APISecret)
	client := apiclient.New(runtime, strfmt.Default)

	return client
}
