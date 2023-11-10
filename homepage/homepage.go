package homepage

import (
	"context"
	"net/http"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	dphandlers "github.com/ONSdigital/dp-net/v2/handlers"
	"github.com/ONSdigital/log.go/v2/log"
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
func Handler(cfg *config.Config, homepageClient Clienter, rend RenderClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, cfg, accessToken, collectionID, lang, homepageClient, rend)
	})
}

func handle(w http.ResponseWriter, req *http.Request, cfg *config.Config, userAccessToken, collectionID, lang string, homepageClient Clienter, rend RenderClient) {
	ctx := req.Context()

	homepageContent, err := homepageClient.GetHomePage(ctx, userAccessToken, collectionID, lang)
	if err != nil {
		log.Error(ctx, "HOMEPAGE_RESPONSE_FAILED. failed to get homepage html content", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	navigationContent, err := homepageClient.GetNavigationData(ctx, lang)
	if err != nil {
		log.Error(ctx, "failed to get navigation data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	basePage := rend.NewBasePageModel()
	m := mapper.Homepage(*cfg, lang, basePage, homepageContent.MainFigures, homepageContent.FeaturedContent, homepageContent.AroundONS, homepageContent.ServiceMessage, homepageContent.EmergencyBanner, navigationContent)

	rend.BuildPage(w, m, "homepage")

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
			log.Error(ctx, "error getting timeseries data for trend indication", err, log.Data{
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
	cfg, err := config.Get()
	if err != nil {
		// do something
	}

	mainFigureMap = make(map[string]MainFigure)

	mainFigureMap["emp"] = MainFigure{
		uris:               []string{"/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms"},
		trendURI:           "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fux7/lms",
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

	// Unemployment
	mainFigureMap["unemp"] = MainFigure{
		uris:               []string{"/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms"},
		trendURI:           "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fuu8/lms",
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

	if cfg.EnableUpdatedMainFigures {
		mainFigureMapCopy := mainFigureMap["emp"]
		mainFigureMapCopy.uris = []string{"/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/s2pw/lms"}
		mainFigureMapCopy.trendURI = "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/s2qm/lms"
		mainFigureMap["emp"] = mainFigureMapCopy

		mainFigureMapCopy = mainFigureMap["unemp"]
		mainFigureMapCopy.uris = []string{"/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/s2pu/lms"}
		mainFigureMapCopy.trendURI = "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/s2qk/lms"
		mainFigureMap["unemp"] = mainFigureMapCopy
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
	for i := 0; i < len(zts); i++ {
		if latest.URI == "" {
			latest = zts[i]
		}
		releaseDate, err := time.Parse(time.RFC3339, zts[i].Description.ReleaseDate)
		if err != nil {
			log.Error(ctx, "failed to parse release date", err, log.Data{"release_date": zts[i].Description.ReleaseDate})
			return zts[i]
		}
		latestReleaseDate, err := time.Parse(time.RFC3339, latest.Description.ReleaseDate)
		if err != nil {
			log.Error(ctx, "failed to parse release date", err, log.Data{"release_date": latest.Description.ReleaseDate})
			return zts[i]
		}
		if releaseDate.After(latestReleaseDate) {
			latest = zts[i]
		}
	}
	return latest
}
