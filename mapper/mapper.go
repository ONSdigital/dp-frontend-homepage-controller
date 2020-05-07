package mapper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"github.com/dustin/go-humanize"
)

// Homepage maps data to our homepage frontend model
func Homepage(ctx context.Context, mainFigures map[string]*model.MainFigure) model.Page {
	var page model.Page
	page.Type = "homepage"
	page.Metadata.Title = "Home"
	page.Data.MainFigures = mainFigures
	return page
}

// MainFigure maps a single main figure object
func MainFigure(ctx context.Context, id, datePeriod string, figure zebedee.TimeseriesMainFigure) *model.MainFigure {
	var mf model.MainFigure

	mf.ID = id

	mfData := getDataByPeriod(datePeriod, figure)
	latestDataIndex := len(mfData) - 1
	previousDataIndex := len(mfData) - 2
	latestData := mfData[latestDataIndex]
	previousData := mfData[previousDataIndex]
	latestFigure, err := strconv.ParseFloat(latestData.Value, 64)
	if err != nil {
		log.Event(ctx, "error getting trend description: error parsing float", log.Error(err))
		return &mf
	}
	previousFigure, err := strconv.ParseFloat(previousData.Value, 64)
	if err != nil {
		log.Event(ctx, "error getting trend description: error parsing float", log.Error(err))
		return &mf
	}

	latestFigureFormatted := humanize.CommafWithDigits(latestFigure, 2)

	mf.Figure = latestFigureFormatted
	mf.Date = latestData.Label
	mf.Unit = figure.Description.Unit
	mf.Trend = getTrend(latestFigure, previousFigure)
	mf.TrendDescription = getTrendDescription(latestFigure, previousFigure, figure.Description.Unit, datePeriod)
	if len(figure.RelatedDocuments) > 0 {
		mf.FigureURIs.Analysis = figure.RelatedDocuments[0].URI
	}
	mf.FigureURIs.Data = figure.URI
	return &mf
}

// getDataByPeriod returns the data for the time period set
func getDataByPeriod(datePeriod string, data zebedee.TimeseriesMainFigure) []zebedee.TimeseriesDataPoint {
	var mf []zebedee.TimeseriesDataPoint
	switch datePeriod {
	case "years":
		mf = data.Years
	case "quarters":
		mf = data.Quarters
	case "months":
		mf = data.Months
	default:
		mf = []zebedee.TimeseriesDataPoint{}
	}
	return mf
}

// getTrend returns trend boolean value based on latest and previous figures
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

// getTrendDescription returns a string describing the trend with the difference from the current figure from the previous
func getTrendDescription(latest, previous float64, unit, datePeriod string) string {
	diff := float64(latest - previous)
	//delta := (diff / previous) * 100
	return fmt.Sprintf("%0.2f%v on previous %v", diff, unit, datePeriod)
}
