package helper

import (
	"time"
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/log.go/v2/log"
)
//checks if current time is equal/after Census First Results time
func CheckTime(req *http.Request, cfg *config.Config) bool {
	var currentTime, goLiveTime time.Time
	currentTime = time.Now()
	ctx := req.Context()
	goLiveTime, err := time.Parse(time.RFC850, cfg.CensusFirstResults)
	if err != nil {
		log.Warn(ctx, "failing to parse configuration of census first results needed for "+
			"going live", log.FormatErrors([]error{err}))
	}
	if currentTime.Equal(goLiveTime) || currentTime.After(goLiveTime) {
		return true
	}
	return false
}
