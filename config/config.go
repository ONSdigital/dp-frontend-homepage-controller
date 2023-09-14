package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-homepage-controller
type Config struct {
	APIRouterURL                   string        `envconfig:"API_ROUTER_URL"`
	BindAddr                       string        `envconfig:"BIND_ADDR"`
	CacheCensusTopicUpdateInterval time.Duration `envconfig:"CACHE_CENSUS_TOPIC_UPDATE_INTERVAL"`
	CacheNavigationUpdateInterval  time.Duration `envconfig:"CACHE_NAVIGATION_UPDATE_INTERVAL"`
	CacheUpdateInterval            time.Duration `envconfig:"CACHE_UPDATE_INTERVAL"`
	CensusTopicID                  string        `envconfig:"CENSUS_TOPIC_ID"`
	DatasetFinderEnabled           bool          `envconfig:"DATASET_FINDER_ENABLED"`
	Debug                          bool          `envconfig:"DEBUG"`
	EnableCensusTopicSubsection    bool          `envconfig:"ENABLE_CENSUS_TOPIC_SUBSECTION"`
	EnableFeedbackAPI              bool          `envconfig:"ENABLE_FEEDBACK_API"`
	EnableGetDataCard              bool          `envconfig:"ENABLE_GET_DATA_CARD"`
	EnableCustomDataset            bool          `envconfig:"ENABLE_CUSTOM_DATASET"`
	EnableNewNavBar                bool          `envconfig:"ENABLE_NEW_NAVBAR"`
	FeedbackAPIURL                 string        `envconfig:"FEEDBACK_API_URL"`
	GracefulShutdownTimeout        time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckCriticalTimeout     time.Duration `envconfig:"HEALTH_CHECK_CRITICAL_TIMEOUT"`
	HealthCheckInterval            time.Duration `envconfig:"HEALTH_CHECK_INTERVAL"`
	IsPublishingMode               bool          `envconfig:"IS_PUBLISHING_MODE"`
	PatternLibraryAssetsPath       string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
	ServiceAuthToken               string        `envconfig:"SERVICE_AUTH_TOKEN"   json:"-"`
	SiteDomain                     string        `envconfig:"SITE_DOMAIN"`
	SupportedLanguages             []string      `envconfig:"SUPPORTED_LANGUAGES"`
	SixteensVersion                string        `envconfig:"SIXTEENS_VERSION"`
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
		cfgVar.PatternLibraryAssetsPath = "//cdn.ons.gov.uk/dp-design-system/c44b4f8"
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
		CacheCensusTopicUpdateInterval: 1 * time.Minute,
		CacheNavigationUpdateInterval:  1 * time.Minute,
		CacheUpdateInterval:            10 * time.Second,
		CensusTopicID:                  "4445",
		DatasetFinderEnabled:           false,
		Debug:                          false,
		EnableCensusTopicSubsection:    false,
		EnableCustomDataset:            false,
		EnableFeedbackAPI:              false,
		EnableGetDataCard:              false,
		EnableNewNavBar:                false,
		FeedbackAPIURL:                 "http://localhost:23200/v1/feedback",
		GracefulShutdownTimeout:        5 * time.Second,
		HealthCheckCriticalTimeout:     90 * time.Second,
		HealthCheckInterval:            30 * time.Second,
		IsPublishingMode:               false,
		ServiceAuthToken:               "",
		SiteDomain:                     "localhost",
		SupportedLanguages:             []string{"en", "cy"},
		SixteensVersion:                "2366f2e",
	}

	return cfg, envconfig.Process("", cfg)
}
