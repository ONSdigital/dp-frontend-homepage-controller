package helper

import (
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckTime(t *testing.T) {
	Convey("CheckTime should return false when current time is before Census Results time", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)

		So(CheckTime(cfg), ShouldEqual, false)
	})

	Convey("CheckTime should return true when current time is after Census Results time", t, func() {
		cfg := initialiseMockConfig()

		So(CheckTime(cfg), ShouldEqual, true)
	})
}

func initialiseMockConfig() *config.Config {
	return &config.Config{
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
		CensusFirstResults:         "Monday, 20-Jun-22 14:29:00 UTC",
	}
}
