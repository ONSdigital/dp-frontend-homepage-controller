package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-homepage-controller
type Config struct {
	APIRouterURL                   string        `envconfig:"API_ROUTER_URL"`
	BindAddr                       string        `envconfig:"BIND_ADDR"`
	CacheCensusTopicUpdateInterval time.Duration `envconfig:"CACHE_CENSUS_TOPICS_UPDATE_INTERVAL"`
	CacheNavigationUpdateInterval  time.Duration `envconfig:"CACHE_NAVIGATION_UPDATE_INTERVAL"`
	CacheUpdateInterval            time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	CensusTopicID                  string        `envconfig:"CENSUS_TOPIC_ID"`
	Debug                          bool          `envconfig:"DEBUG"`
	EnableNewNavBar                bool          `envconfig:"ENABLE_NEW_NAVBAR"`
	GracefulShutdownTimeout        time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckCriticalTimeout     time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	HealthCheckInterval            time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	IsPublishingMode               bool          `envconfig:"IS_PUBLISHING_MODE"`
	Languages                      string        `envconfig:"LANGUAGES"`
	PatternLibraryAssetsPath       string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	SiteDomain                     string        `envconfig:"SITE_DOMAIN"`
	ServiceAuthToken               string        `envconfig:"SERVICE_AUTH_TOKEN"   json:"-"`
	SupportedLanguages             [2]string     `envconfig:"SUPPORTED_LANGUAGES"`
	CensusTopicsSubsectionFeature  bool          `envconfig:"CENSUS_TOPICS_SUBSECTION_FEATURE"`
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
		cfgVar.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/dp-design-system/67b5a7e"
	}
	return cfgVar, nil
}

func get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		APIRouterURL:                   "http://localhost:23200/v1",
		BindAddr:                       ":24400",
		CacheCensusTopicUpdateInterval: 5 * time.Second, // TODO 30 * time.Minute
		CacheNavigationUpdateInterval:  10 * time.Second,
		CacheUpdateInterval:            10 * time.Second,
		//CensusTopicID:                  "4445",
		CensusTopicID:                 "2747",
		Debug:                         false,
		EnableNewNavBar:               false,
		CensusTopicsSubsectionFeature: true, // TODO Set to false
		GracefulShutdownTimeout:       5 * time.Second,
		HealthCheckCriticalTimeout:    90 * time.Second,
		HealthCheckInterval:           30 * time.Second,
		IsPublishingMode:              false,
		Languages:                     "en,cy",
		ServiceAuthToken:              "",
		SiteDomain:                    "localhost",
		SupportedLanguages:            [2]string{"en", "cy"},
	}

	return cfg, envconfig.Process("", cfg)
}
