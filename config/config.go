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
	Debug                      bool          `envconfig:"DEBUG"`
	PatternLibraryAssetsPath   string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	SupportedLanguages         [2]string     `envconfig:"SUPPORTED_LANGUAGES"`
	CacheUpdateInterval        time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	IsPublishingMode           bool          `envconfig:"IS_PUBLISHING_MODE"`
	Languages                  string        `envconfig:"LANGUAGES"`
	SiteDomain                 string        `envconfig:"SITE_DOMAIN"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	cfg, err := get()
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		cfg.PatternLibraryAssetsPath = "http://localhost:9000/dist"
	} else {
		cfg.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/f80be2c"
	}
	return cfg, nil
}

func get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		BindAddr:                   ":24400",
		APIRouterURL:               "http://localhost:23200/v1",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		Debug:                      false,
		SupportedLanguages:         [2]string{"en", "cy"},
		CacheUpdateInterval:        10 * time.Second,
		IsPublishingMode:           false,
		Languages:                  "en,cy",
		SiteDomain:                 "localhost",
	}

	return cfg, envconfig.Process("", cfg)
}
