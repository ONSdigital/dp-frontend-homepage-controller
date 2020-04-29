package mapper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
)

// Homepage maps data to our homepage frontend model
func Homepage(ctx context.Context, mainFigures []model.MainFigure) model.Page {
	var page model.Page
	page.Type = "homepage"
	page.Data.MainFigures = mainFigures
	return page
}

// MainFigure maps a single main figure object
func MainFigure(ctx context.Context, datePeriod string, data zebedee.TimeseriesMainFigure) model.MainFigure {
	var mf model.MainFigure

	mfData := getDataByPeriod(datePeriod, data)
	latestData := mfData[(len(mfData) - 1)]
	previousData := mfData[(len(mfData) - 2)]
	latestFigure, err := strconv.ParseFloat(latestData.Value, 64)
	if err != nil {
		log.Event(ctx, "error getting trend description: error parsing float", log.Error(err))
		return mf
	}
	previousFigure, err := strconv.ParseFloat(previousData.Value, 64)
	if err != nil {
		log.Event(ctx, "error getting trend description: error parsing float", log.Error(err))
		return mf
	}

	mf.Figure = fmt.Sprintf("%0.1f", latestFigure)
	mf.Date = latestData.Label
	mf.Unit = data.Description.Unit
	mf.Trend = getTrend(latestFigure, previousFigure)
	mf.TrendDescription = getTrendDescription(latestFigure, previousFigure, data.Description.Unit, datePeriod)
	if len(data.RelatedDocuments) > 0 {
		mf.FigureURIs.Analysis = data.RelatedDocuments[0].URI
	}
	mf.FigureURIs.Data = data.URI
	return mf
}

func getDataByPeriod(datePeriod string, data zebedee.TimeseriesMainFigure) []zebedee.TimeseriesDataPoint {
	var mf []zebedee.TimeseriesDataPoint
	switch datePeriod {
	case "years":
		mf = data.Years
	case "quarters":
		mf = data.Quarters
	case "months":
		mf = data.Months
	}
	return mf
}

func getTrend(latest, previous float64) model.Trend {
	if latest > previous {
		return model.Trend{IsUp: true}
	}

	if latest < previous {
		return model.Trend{IsDown: true}
	}

	if latest == previous {
		return model.Trend{IsFlat: true}
	}
	return model.Trend{}
}

func getTrendDescription(latest, previous float64, unit, datePeriod string) string {
	diff := float64(latest - previous)
	//delta := (diff / previous) * 100
	return fmt.Sprintf("%0.2f%v on previous %v", diff, unit, datePeriod)
}
