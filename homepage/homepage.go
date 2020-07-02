package homepage

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/headers"
	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/log.go/log"
)

const (
	// PeriodYear is the string value for year time period
	PeriodYear = "year"
	// PeriodQuarter is the string value for quarter time period
	PeriodQuarter = "quarter"
	// PeriodMonth is the string value for month time period
	PeriodMonth = "month"
	// HomepagePath is the string value which contains the URI to get the homepage's data.json
	HomepagePath = "/"
)

type MainFigure struct {
	uri        string
	datePeriod string
	data       zebedee.TimeseriesMainFigure
}

var mainFigureMap map[string]MainFigure

// Handler handles requests to homepage endpoint
func Handler(rend RenderClient, zcli ZebedeeClient, bcli BabbageClient, icli ImageClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handle(w, req, rend, zcli, bcli, icli)
	}
}

func handle(w http.ResponseWriter, req *http.Request, rend RenderClient, zcli ZebedeeClient, bcli BabbageClient, icli ImageClient) {
	ctx := req.Context()

	userAccessToken, err := headers.GetUserAuthToken(req)
	if err != nil {
		log.Event(ctx, "unable to get user access token from header setting it to empty value", log.WARN, log.Error(err))
		userAccessToken = ""
	}

	var localeCode string
	if ctx.Value(common.LocaleHeaderKey) != nil {
		var ok bool
		localeCode, ok = ctx.Value(common.LocaleHeaderKey).(string)
		if !ok {
			log.Event(ctx, "error retrieving locale code", log.WARN, log.Error(errors.New("error casting locale code to string")))
		}
	}

	mappedMainFigures := make(map[string]*model.MainFigure)
	var wg sync.WaitGroup
	responses := make(chan *model.MainFigure, len(mainFigureMap))
	for id, figure := range mainFigureMap {
		wg.Add(1)
		go func(ctx context.Context, zcli ZebedeeClient, id string, figure MainFigure) {
			defer wg.Done()
			zebResp, err := zcli.GetTimeseriesMainFigure(ctx, userAccessToken, figure.uri)
			if err != nil {
				log.Event(ctx, "error getting timeseries data", log.Error(err), log.Data{"timeseries-data": figure.uri})
				mappedErrorFigure := &model.MainFigure{ID: id}
				responses <- mappedErrorFigure
				return
			}
			mappedMainFigure := mapper.MainFigure(ctx, id, figure.datePeriod, zebResp)
			responses <- mappedMainFigure
			return
		}(ctx, zcli, id, figure)
	}
	wg.Wait()
	close(responses)

	for response := range responses {
		mappedMainFigures[response.ID] = response
	}

	weekAgoTime := time.Now().AddDate(0, 0, -7)
	dateFromDay := weekAgoTime.Format("02")
	dateFromMonth := weekAgoTime.Format("01")
	dateFromYear := weekAgoTime.Format("2006")
	releaseCalResp, err := bcli.GetReleaseCalendar(ctx, userAccessToken, dateFromDay, dateFromMonth, dateFromYear)
	releaseCalModelData := mapper.ReleaseCalendar(releaseCalResp)

	// Get homepage data from Zebedee
	homepageContent, err := zcli.GetHomepageContent(ctx, userAccessToken, HomepagePath)
	if err != nil {
		log.Event(ctx, "error getting homepage data from client", log.Error(err), log.Data{"content-path": HomepagePath})
	}
	imageObjects := []image.Image{}
	for _, fc := range homepageContent.FeaturedContent {
		image, err := icli.GetImage(ctx, userAccessToken, "", "", fc.ImageID)
		if err != nil {
			log.Event(ctx, "error getting image", log.Error(err), log.Data{"featured-content-entry": fc.Title})
		}
		imageObjects = append(imageObjects, image)
	}

	mappedFeaturedContent := mapper.FeaturedContent(homepageContent, imageObjects)

	m := mapper.Homepage(localeCode, mappedMainFigures, releaseCalModelData, &mappedFeaturedContent)

	b, err := json.Marshal(m)
	if err != nil {
		log.Event(ctx, "error marshalling body data to json", log.Error(err))
		http.Error(w, "error marshalling body data to json", http.StatusInternalServerError)
		return
	}

	templateHTML, err := rend.Do("homepage", b)
	if err != nil {
		log.Event(ctx, "error rendering page", log.Error(err))
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

func init() {
	mainFigureMap = make(map[string]MainFigure)

	// Employment
	mainFigureMap["LF24"] = MainFigure{
		uri:        "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms",
		datePeriod: PeriodYear,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Unemployment
	mainFigureMap["MGSX"] = MainFigure{
		uri:        "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms",
		datePeriod: PeriodYear,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Inflation (CPIH)
	mainFigureMap["L55O"] = MainFigure{
		uri:        "/economy/inflationandpriceindices/timeseries/l55o/mm23",
		datePeriod: PeriodMonth,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// GDP
	mainFigureMap["IHYQ"] = MainFigure{
		uri:        "/economy/grossdomesticproductgdp/timeseries/ihyq/qna",
		datePeriod: PeriodQuarter,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Population
	mainFigureMap["UKPOP"] = MainFigure{
		uri:        "/peoplepopulationandcommunity/populationandmigration/populationestimates/timeseries/ukpop/pop",
		datePeriod: PeriodYear,
		data:       zebedee.TimeseriesMainFigure{},
	}

}
