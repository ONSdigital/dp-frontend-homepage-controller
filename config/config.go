package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-homepage-controller
type Config struct {
	APIRouterURL                  string        `envconfig:"API_ROUTER_URL"`
	BindAddr                      string        `envconfig:"BIND_ADDR"`
	CacheNavigationUpdateInterval time.Duration `envconfig:"CACHE_NAVIGATION_UPDATE_INTERVAL"`
	CacheUpdateInterval           time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	Debug                         bool          `envconfig:"DEBUG"`
	EnableNewNavBar               bool          `envconfig:"ENABLE_NEW_NAVBAR"`
	GracefulShutdownTimeout       time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckCriticalTimeout    time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	HealthCheckInterval           time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	IsPublishingMode              bool          `envconfig:"IS_PUBLISHING_MODE"`
	Languages                     string        `envconfig:"LANGUAGES"`
	PatternLibraryAssetsPath      string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	SiteDomain                    string        `envconfig:"SITE_DOMAIN"`
	SupportedLanguages            [2]string     `envconfig:"SUPPORTED_LANGUAGES"`
	EnableCensusTopicSubsection   bool          `envconfig:"ENABLE_CENSUS_TOPIC_SUBSECTION"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	cfgVar, err := get()
	if err != nil {
		return nil, err
	}

	if cfgVar.Debug {
		cfgVar.PatternLibraryAssetsPath = "http://localhost:9002/dist/assets"
	} else {
		cfgVar.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/dp-design-system/297621d"
	}
	return cfgVar, nil
}

func get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		BindAddr:                      ":24400",
		APIRouterURL:                  "http://localhost:23200/v1",
		CacheNavigationUpdateInterval: 10 * time.Second,
		CacheUpdateInterval:           10 * time.Second,
		Debug:                         false,
		EnableNewNavBar:               false,
		GracefulShutdownTimeout:       5 * time.Second,
		HealthCheckCriticalTimeout:    90 * time.Second,
		HealthCheckInterval:           30 * time.Second,
		IsPublishingMode:              false,
		Languages:                     "en,cy",
		SiteDomain:                    "localhost",
		SupportedLanguages:            [2]string{"en", "cy"},
		EnableCensusTopicSubsection:   false,
	}

	return cfg, envconfig.Process("", cfg)
}
