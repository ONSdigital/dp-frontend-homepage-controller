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
				So(cfg.BindAddr, ShouldEqual, ":24400")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 10*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 60*time.Second)
				So(cfg.RendererURL, ShouldEqual, "http://localhost:20010")
			})
		})
	})
}
