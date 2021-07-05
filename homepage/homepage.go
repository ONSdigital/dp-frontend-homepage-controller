package homepage

import (
	"context"
	"net/http"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
	"github.com/ONSdigital/log.go/log"
)

const (
	// HomepagePath is the string value which contains the URI to get the homepage's data.json
	HomepagePath = "/"

	// ImageVariant is the image variant to use for the homepage
	ImageVariant = "original"
)

type MainFigure struct {
	uris               []string
	trendURI           string
	datePeriod         string
	data               zebedee.TimeseriesMainFigure
	differenceInterval string
}

var mainFigureMap map[string]MainFigure

// Handler handles requests to homepage endpoint
func Handler(homepageClient HomepageClienter) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, accessToken, collectionID, lang, homepageClient)
	})
}

func handle(w http.ResponseWriter, req *http.Request, userAccessToken, collectionID, lang string, homepageClient HomepageClienter) {
	ctx := req.Context()
	homepageHTML, err := homepageClient.GetHomePage(ctx, userAccessToken, collectionID, lang)
	if err != nil {
		log.Event(ctx, "HOMEPAGE_RESPONSE_FAILED. failed to get homepage html", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(homepageHTML)); err != nil {
		log.Event(ctx, "failed to write response for homepage", log.ERROR, log.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

func getTrendInfo(ctx context.Context, userAccessToken, collectionID, lang string, zcli ZebedeeClient, figure MainFigure) mapper.TrendInfo {
	trendResponse := zebedee.TimeseriesMainFigure{}
	var err error
	retrieveTrendFailed := false
	isTimeseriesForTrend := false
	if figure.trendURI != "" {
		isTimeseriesForTrend = true
		trendResponse, err = zcli.GetTimeseriesMainFigure(ctx, userAccessToken, collectionID, lang, figure.trendURI)
		if err != nil {
			// Error getting timeseries, log it but continue to construct rest of main figure tile
			retrieveTrendFailed = true
			log.Event(ctx, "error getting timeseries data for trend indication", log.ERROR, log.Error(err), log.Data{
				"timeseries-data": figure.trendURI,
				"trendResponse":   trendResponse,
			})
		}
	}
	return mapper.TrendInfo{
		TrendFigure:          trendResponse,
		IsTimeseriesForTrend: isTimeseriesForTrend,
		RetrieveTrendFailed:  retrieveTrendFailed,
	}
}

func init() {
	mainFigureMap = make(map[string]MainFigure)

	// Employment
	mainFigureMap["LF24"] = MainFigure{
		uris:               []string{"/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms"},
		trendURI:           "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fux7/lms",
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

	// Unemployment
	mainFigureMap["MGSX"] = MainFigure{
		uris:               []string{"/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms"},
		trendURI:           "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fuu8/lms",
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

	// Inflation (CPIH)
	mainFigureMap["L55O"] = MainFigure{
		uris:               []string{"/economy/inflationandpriceindices/timeseries/l55o/mm23"},
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodMonth,
	}

	// GDP
	mainFigureMap["IHYQ"] = MainFigure{
		uris:               []string{"/economy/grossdomesticproductgdp/timeseries/ihyq/qna", "/economy/grossdomesticproductgdp/timeseries/ihyq/pn2"},
		datePeriod:         mapper.PeriodQuarter,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodQuarter,
	}

	// Population
	mainFigureMap["UKPOP"] = MainFigure{
		uris:               []string{"/peoplepopulationandcommunity/populationandmigration/populationestimates/timeseries/ukpop/pop"},
		datePeriod:         mapper.PeriodYear,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

}

func getLatestTimeSeriesData(ctx context.Context, zts []zebedee.TimeseriesMainFigure) zebedee.TimeseriesMainFigure {
	var latest zebedee.TimeseriesMainFigure

	for _, ts := range zts {
		if latest.URI == "" {
			latest = ts
		}
		releaseDate, err := time.Parse(time.RFC3339, ts.Description.ReleaseDate)
		if err != nil {
			log.Event(ctx, "failed to parse release date", log.ERROR, log.Error(err), log.Data{"release_date": ts.Description.ReleaseDate})
			return ts
		}
		latestReleaseDate, err := time.Parse(time.RFC3339, latest.Description.ReleaseDate)
		if err != nil {
			log.Event(ctx, "failed to parse release date", log.ERROR, log.Error(err), log.Data{"release_date": latest.Description.ReleaseDate})
			return ts
		}
		if releaseDate.After(latestReleaseDate) {
			latest = ts
		}
	}
	return latest
}
