package helper

import (
	"github.com/ONSdigital/dp-renderer/helper"
	"testing"

	"dp-frontend-homepage-controller/helper/helper.go"
)
func TestCheckTime(t *testing.T) {
	Convey("CheckTime should return false when current time is before go live time", t, func() {
		So(helper.CheckTime(99, 1), ShouldEqual, 100)
	})
}

