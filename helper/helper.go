package helper

import (
	"context"
	"time"

	"github.com/ONSdigital/log.go/v2/log"
)
// CheckTime checks if the current time is equal/before the value of time variable passed into function

func CheckTime(timeToCompare string) bool {
	var currentTime, goLiveTime time.Time
	currentTime = time.Now()
	goLiveTime, err := time.Parse(time.RFC850, timeToCompare)
	if err != nil {
		log.Warn(context.Background(), "failing to parse configuration of census first results needed for "+
			"going live", log.FormatErrors([]error{err}))
	}
	if currentTime.Equal(goLiveTime) || currentTime.After(goLiveTime) {
		return true
	}
	return false
}
