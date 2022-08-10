package helper

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	expectedLayout   = time.RFC850  // "Monday, 02-Jan-06 15:04:05 MST"
	unexpectedLayout = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

func TestCheckTime(t *testing.T) {
	testCtx := context.Background()

	currentTime := time.Now().UTC()

	Convey("CheckTime should return false when current time is before configured time", t, func() {
		timeInFuture := currentTime.Add(1 * time.Minute).Format(expectedLayout)
		So(CheckTime(testCtx, timeInFuture), ShouldEqual, false)
	})

	Convey("CheckTime should return true when current time is after configured time", t, func() {
		timeInPast := currentTime.Add(-1 * time.Minute).Format(expectedLayout)
		So(CheckTime(testCtx, timeInPast), ShouldEqual, true)
	})

	Convey("CheckTime should return false when the configured time is in the incorrect time format", t, func() {
		timeInPast := currentTime.Add(-1 * time.Minute).Format(unexpectedLayout)
		So(CheckTime(testCtx, timeInPast), ShouldEqual, false)
	})
}
