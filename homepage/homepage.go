package homepage

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/headers"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
)

type MainFigure struct {
	uri        string
	datePeriod string
	data       zebedee.TimeseriesMainFigure
}

//go:generate moq -out mocks_test.go -pkg homepage . ZebedeeClient RenderClient

// ZebedeeClient is an interface with methods required for a zebedee client
type ZebedeeClient interface {
	GetTimeseriesMainFigure(ctx context.Context, userAuthToken, uri string) (m zebedee.TimeseriesMainFigure, err error)
}

// RenderClient is an interface with methods for require for rendering a template
type RenderClient interface {
	Do(string, []byte) ([]byte, error)
}

// Handler handles requests to homepage endpoint
func Handler(rend RenderClient, zcli ZebedeeClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handle(w, req, rend, zcli)
	}
}

func handle(w http.ResponseWriter, req *http.Request, rend RenderClient, zcli ZebedeeClient) {
	ctx := req.Context()

	userAccessToken, err := headers.GetUserAuthToken(req)
	if err == headers.ErrHeaderNotFound {
		log.Event(ctx, "no user access token header set", log.ERROR, log.Error(err))
		userAccessToken = ""
	}
	if err != nil {
		log.Event(ctx, "error getting user access token from header", log.ERROR, log.Error(err))
		userAccessToken = ""
	}

	mainFiguresList := getMainFiguresList()

	var mappedMainFigures []model.MainFigure
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for _, value := range mainFiguresList {
		wg.Add(1)
		go func(ctx context.Context, zcli ZebedeeClient, value MainFigure) {
			defer wg.Done()
			zebResp, err := zcli.GetTimeseriesMainFigure(ctx, userAccessToken, value.uri)
			if err != nil {
				log.Event(ctx, "error getting timeseries data", log.Error(err))
				http.Error(w, "error getting timeseries data", http.StatusBadRequest)
			}
			mappedMainFigure := mapper.MainFigure(ctx, value.datePeriod, zebResp)
			mutex.Lock()
			defer mutex.Unlock()
			mappedMainFigures = append(mappedMainFigures, mappedMainFigure)
			return
		}(ctx, zcli, value)
	}
	wg.Wait()

	m := mapper.Homepage(ctx, mappedMainFigures)

	b, err := json.Marshal(m)
	if err != nil {
		log.Event(ctx, "error marsahlling body data to json", log.Error(err))
		http.Error(w, "error marsahlling body data to json", http.StatusBadRequest)
		return
	}

	templateHTML, err := rend.Do("homepage", b)
	if err != nil {
		log.Event(ctx, "error rendering page", log.Error(err))
		http.Error(w, "error rendering page", http.StatusInternalServerError)
		return
	}

	w.Write(templateHTML)
	return
}

const (
	// PeriodYears is the string value for years time period
	PeriodYears = "years"
	// PeriodQuarters is the string value for quarters time period
	PeriodQuarters = "quarters"
	// PeriodMonths is the string value for months time period
	PeriodMonths = "months"
)

func getMainFiguresList() map[string]MainFigure {
	mainFigureMap := make(map[string]MainFigure)

	// Employment
	mainFigureMap["LF24"] = MainFigure{
		uri:        "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms",
		datePeriod: PeriodMonths,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Unemployment
	mainFigureMap["MGSX"] = MainFigure{
		uri:        "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms",
		datePeriod: PeriodMonths,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Inflation (CPIH)
	mainFigureMap["L55O"] = MainFigure{
		uri:        "/economy/inflationandpriceindices/timeseries/l55o/mm23",
		datePeriod: PeriodMonths,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// GDP
	mainFigureMap["IHYQ"] = MainFigure{
		uri:        "/economy/grossdomesticproductgdp/timeseries/ihyq/qna",
		datePeriod: PeriodQuarters,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Population
	mainFigureMap["UKPOP"] = MainFigure{
		uri:        "/peoplepopulationandcommunity/populationandmigration/populationestimates/timeseries/ukpop/pop",
		datePeriod: PeriodYears,
		data:       zebedee.TimeseriesMainFigure{},
	}

	return mainFigureMap
}
