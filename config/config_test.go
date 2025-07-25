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
				So(cfg.CacheCensusTopicUpdateInterval, ShouldEqual, 60*time.Second)
				So(cfg.CacheNavigationUpdateInterval, ShouldEqual, 60*time.Second)
				So(cfg.CacheUpdateInterval, ShouldEqual, 10*time.Second)
				So(cfg.DatasetFinderEnabled, ShouldEqual, false)
				So(cfg.EnableCensusTopicSubsection, ShouldEqual, false)
				So(cfg.EnableGetDataCard, ShouldEqual, false)
				So(cfg.EnableNewNavBar, ShouldEqual, false)
				So(cfg.EnablePreviewSiteTile, ShouldEqual, false)
				So(cfg.PreviewSiteURL, ShouldEqual, "https://nwp-prototype.ons.gov.uk/")
				So(cfg.FeedbackAPIURL, ShouldEqual, "http://localhost:23200/v1/feedback")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.IsPublishingMode, ShouldEqual, false)
				So(cfg.PatternLibraryAssetsPath, ShouldEqual, "//cdn.ons.gov.uk/dp-design-system/f3e1909")
				So(cfg.SiteDomain, ShouldEqual, "localhost")
				So(cfg.SupportedLanguages, ShouldResemble, []string{"en", "cy"})
			})
		})
	})
}
