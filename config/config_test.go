package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given an environment with no environment variables set", t, func() {
		cfg, err := Get()

		Convey("When the config values are retrieved", func() {

			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("That the values should be set to the expected defaults", func() {
				So(cfg.APIRouterURL, ShouldEqual, "http://localhost:23200/v1")
				So(cfg.BindAddr, ShouldEqual, ":24400")
				So(cfg.CacheNavigationUpdateInterval, ShouldEqual, 10*time.Second)
				So(cfg.CacheUpdateInterval, ShouldEqual, 10*time.Second)
				So(cfg.CensusFirstResults, ShouldEqual, "Wednesday, 27-Jul-22 11:00:00 BST")
				So(cfg.EnableNewNavBar, ShouldEqual, false)
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.IsPublishingMode, ShouldEqual, false)
				So(cfg.Languages, ShouldEqual, "en,cy")
				So(cfg.PatternLibraryAssetsPath, ShouldEqual, "//cdn.ons.gov.uk/dp-design-system/109aa8d")
				So(cfg.SiteDomain, ShouldEqual, "localhost")
			})
		})
	})
}
