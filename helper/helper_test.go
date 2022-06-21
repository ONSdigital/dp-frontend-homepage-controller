package helper

import (
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckTime(t *testing.T) {
	Convey("CheckTime should return false when current time is before Census Results time", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)

		So(CheckTime(cfg.CensusFirstResults), ShouldEqual, false)
	})

	Convey("CheckTime should return true when current time is after Census Results time", t, func() {
		cfg := "Monday, 20-Jun-22 14:29:00 UTC"

		So(CheckTime(cfg), ShouldEqual, true)
	})
}
