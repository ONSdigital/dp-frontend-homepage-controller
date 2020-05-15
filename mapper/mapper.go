package mapper

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	"strconv"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"github.com/dustin/go-humanize"
)

// Homepage maps data to our homepage frontend model
func Homepage(ctx context.Context, localeCode string, mainFigures map[string]*model.MainFigure, releaseCal *model.ReleaseCalendar) model.Page {
	var page model.Page
	page.Type = "homepage"
	page.Metadata.Title = "Home"
	page.Language = localeCode
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
	mf.Trend.Difference = getTrendDifference(latestFigure, previousFigure, figure.Description.Unit)
	mf.Trend.Period = datePeriod
	if len(figure.RelatedDocuments) > 0 {
		mf.FigureURIs.Analysis = figure.RelatedDocuments[0].URI
	}
	mf.FigureURIs.Data = figure.URI
	return &mf
}

func ReleaseCalendar(ctx context.Context, rawReleaseCalendar release_calendar.ReleaseCalendar) *model.ReleaseCalendar{
	var rc model.ReleaseCalendar
	return &rc
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

// getTrendDifference returns string value of the difference between latest and previous
func getTrendDifference(latest, previous float64, unit string) string {
	var trendUnit string
	switch unit {
	case "%":
		trendUnit = "pp"
	default:
		trendUnit = unit
	}
	diff := float64(latest - previous)
	formattedDiff := humanize.CommafWithDigits(diff, 2)
	return fmt.Sprintf("%v%v", formattedDiff, trendUnit)
}
