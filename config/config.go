package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-homepage-controller
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	APIRouterURL               string        `envconfig:"API_ROUTER_URL"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	RendererURL                string        `envconfig:"RENDERER_URL"`
	CacheUpdateInterval        time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	IsPublishingMode           bool          `envconfig:"IS_PUBLISHING_MODE"`
	Languages                  string        `envconfig:"LANGUAGES"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		BindAddr:                   ":24400",
		APIRouterURL:               "http://localhost:23200/v1",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		RendererURL:                "http://localhost:20010",
		CacheUpdateInterval:        10 * time.Second,
		IsPublishingMode:           false,
		Languages:                  "en,cy",
	}

	return cfg, envconfig.Process("", cfg)
}
