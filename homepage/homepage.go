package homepage

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
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
func Handler(rend RenderClient, zcli ZebedeeClient, bcli BabbageClient, icli ImageClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, rend, zcli, bcli, icli, accessToken, collectionID, lang)
	})

}

func handle(w http.ResponseWriter, req *http.Request, rend RenderClient, zcli ZebedeeClient, bcli BabbageClient, icli ImageClient, userAccessToken, collectionID, lang string) {
	ctx := req.Context()

	mappedMainFigures := make(map[string]*model.MainFigure)
	var wg sync.WaitGroup
	responses := make(chan *model.MainFigure, len(mainFigureMap))
	for id, figure := range mainFigureMap {
		wg.Add(1)
		go func(ctx context.Context, zcli ZebedeeClient, id string, figure MainFigure) {
			defer wg.Done()
			zebResponses := []zebedee.TimeseriesMainFigure{}
			for _, uri := range figure.uris {
				zebResponse, err := zcli.GetTimeseriesMainFigure(ctx, userAccessToken, collectionID, lang, uri)
				if err != nil {
					log.Event(ctx, "error getting timeseries data", log.ERROR, log.Error(err), log.Data{"timeseries-data": uri})
					mappedErrorFigure := &model.MainFigure{ID: id}
					responses <- mappedErrorFigure
					return
				}
				zebResponses = append(zebResponses, zebResponse)
			}
			trendInfo := getTrendInfo(ctx, userAccessToken, collectionID, lang, zcli, figure)
			latestMainFigure := getLatestTimeSeriesData(ctx, zebResponses)
			mappedMainFigure := mapper.MainFigure(ctx, id, figure.datePeriod, figure.differenceInterval, latestMainFigure, trendInfo)
			responses <- mappedMainFigure
			return
		}(ctx, zcli, id, figure)
	}
	wg.Wait()
	close(responses)

	for response := range responses {
		log.Event(ctx, "the response of this request was", log.ERROR, log.Data{"response": response})
		mappedMainFigures[response.ID] = response
	}

	weekAgoTime := time.Now().AddDate(0, 0, -7)
	dateFromDay := weekAgoTime.Format("02")
	dateFromMonth := weekAgoTime.Format("01")
	dateFromYear := weekAgoTime.Format("2006")
	releaseCalResp, err := bcli.GetReleaseCalendar(ctx, userAccessToken, dateFromDay, dateFromMonth, dateFromYear)
	if err != nil {
		log.Event(ctx, "error failed to get release calendar data from babbage ", log.ERROR, log.Error(err))
	}
	releaseCalModelData := mapper.ReleaseCalendar(releaseCalResp)

	// Get homepage data from Zebedee
	homepageContent, err := zcli.GetHomepageContent(ctx, userAccessToken, collectionID, lang, HomepagePath)
	if err != nil {
		log.Event(ctx, "error getting homepage data from client", log.ERROR, log.Error(err), log.Data{"content-path": HomepagePath})
	}

	var mappedFeaturedContent []model.Feature
	if len(homepageContent.FeaturedContent) > 0 {
		imageObjects := map[string]image.ImageDownload{}
		for _, fc := range homepageContent.FeaturedContent {
			if fc.ImageID != "" {
				image, err := icli.GetDownloadVariant(ctx, userAccessToken, "", "", fc.ImageID, ImageVariant)
				if err != nil {
					log.Event(ctx, "error getting image download variant", log.ERROR, log.Error(err), log.Data{"featured-content-entry": fc.Title})
				}
				imageObjects[fc.ImageID] = image
			}
		}
		mappedFeaturedContent = mapper.FeaturedContent(homepageContent, imageObjects)
	}

	m := mapper.Homepage(lang, mappedMainFigures, releaseCalModelData, &mappedFeaturedContent, homepageContent.ServiceMessage)

	b, err := json.Marshal(m)
	if err != nil {
		log.Event(ctx, "error marshalling body data to json", log.ERROR, log.Error(err))
		http.Error(w, "error marshalling body data to json", http.StatusInternalServerError)
		return
	}

	templateHTML, err := rend.Do("homepage", b)
	if err != nil {
		log.Event(ctx, "error rendering page", log.ERROR, log.Error(err))
		http.Error(w, "error rendering page", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(templateHTML); err != nil {
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
		trendURI:           "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/FUX7/lms",
		datePeriod:         mapper.PeriodMonth,
		data:               zebedee.TimeseriesMainFigure{},
		differenceInterval: mapper.PeriodYear,
	}

	// Unemployment
	mainFigureMap["MGSX"] = MainFigure{
		uris:               []string{"/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms"},
		trendURI:           "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/FUU8/lms",
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
