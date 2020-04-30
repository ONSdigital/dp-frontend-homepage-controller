package homepage

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"github.com/davecgh/go-spew/spew"
)

type MainFigure struct {
	uri        string
	datePeriod string
	data       zebedee.TimeseriesMainFigure
}

// Handler handles requests to homepage endpoint
func Handler(rend renderer.Renderer, zcli *zebedee.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handle(w, req, rend, zcli)
	}
}

func handle(w http.ResponseWriter, req *http.Request, rend renderer.Renderer, zcli *zebedee.Client) {
	ctx := req.Context()
	mainFiguresList := getMainFiguresList()

	var mappedMainFigures []model.MainFigure
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for _, value := range mainFiguresList {
		wg.Add(1)
		go func(ctx context.Context, zcli *zebedee.Client, value MainFigure) {
			defer wg.Done()
			zebResp, err := zcli.GetTimeseriesMainFigure(ctx, "", value.uri)
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

	spew.Dump(m)

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
	PERIOD_YEARS    = "years"
	PERIOD_QUARTERS = "quarters"
	PERIOD_MONTHS   = "months"
)

func getMainFiguresList() map[string]MainFigure {
	mainFigureMap := make(map[string]MainFigure)

	// Employment
	// mainFigureMap["LF24"] = MainFigure{
	// 	uri:        "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms",
	// 	datePeriod: PERIOD_MONTHS,
	// 	data:       zebedee.TimeseriesMainFigure{},
	// }

	// Unemployment
	// mainFigureMap["MGSX"] = MainFigure{
	// 	uri:        "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms",
	// 	datePeriod: PERIOD_MONTHS,
	// 	data:       zebedee.TimeseriesMainFigure{},
	// }

	// Inflation (CPIH)
	mainFigureMap["L55O"] = MainFigure{
		uri:        "/economy/inflationandpriceindices/timeseries/l55o/mm23",
		datePeriod: PERIOD_MONTHS,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// GDP
	mainFigureMap["IHYQ"] = MainFigure{
		uri:        "/economy/grossdomesticproductgdp/timeseries/ihyq/qna",
		datePeriod: PERIOD_QUARTERS,
		data:       zebedee.TimeseriesMainFigure{},
	}

	// Population
	mainFigureMap["UKPOP"] = MainFigure{
		uri:        "/peoplepopulationandcommunity/populationandmigration/populationestimates/timeseries/ukpop/pop",
		datePeriod: PERIOD_YEARS,
		data:       zebedee.TimeseriesMainFigure{},
	}

	return mainFigureMap
}
