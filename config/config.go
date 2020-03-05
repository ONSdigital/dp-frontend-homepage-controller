package config

import (
	"time"
	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-frontend-homepage-controller
type Config struct {
	BindAddr             	    string          `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout     time.Duration   `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval         time.Duration   `envconfig:"HEALTHCHECK_INTERVAL"`
    HealthCheckCriticalTimeout  time.Duration   `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	Emphasise                   bool            `envconfig:"EMPHASISE"`

}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg := &Config{
		BindAddr:                   ":",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        10 * time.Second,
        HealthCheckCriticalTimeout: time.Minute,
		Emphasise:                  true,


	}

	return cfg, envconfig.Process("", cfg)
}