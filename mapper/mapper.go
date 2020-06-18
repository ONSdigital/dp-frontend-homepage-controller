package mapper

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"github.com/dustin/go-humanize"
)

// Homepage maps data to our homepage frontend model
func Homepage(localeCode string, mainFigures map[string]*model.MainFigure, releaseCal *model.ReleaseCalendar, featuredContent *[]model.Feature) model.Page {
	var page model.Page
	page.Type = "homepage"
	page.Metadata.Title = "Home"
	page.Language = localeCode
	page.Data.MainFigures = mainFigures
	page.Data.ReleaseCalendar = *releaseCal
	page.Data.Featured = *featuredContent
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

func ReleaseCalendar(rawReleaseCalendar release_calendar.ReleaseCalendar) *model.ReleaseCalendar {
	releaseResults := *rawReleaseCalendar.Result.Results
	numReleasesScheduled := rawReleaseCalendar.Result.NumberOfResults

	for i := len(releaseResults) - 1; i >= 0; i-- {
		if releaseResults[i].Description.Cancelled || !releaseResults[i].Description.Published {
			numReleasesScheduled--
		}
	}

	latestReleases := getLatestReleases(releaseResults)
	rc := model.ReleaseCalendar{
		Releases:                         latestReleases,
		NumberOfReleases:                 numReleasesScheduled,
		NumberOfOtherReleasesInSevenDays: numReleasesScheduled - len(latestReleases),
	}
	return &rc
}

// FeaturedContent takes the homepageContent as returned from the client and returns an array of featured content
func FeaturedContent(homepageData zebedee.HomepageContent) []model.Feature {
	var mappedFeaturesArray []model.Feature
	for _, fc := range homepageData.FeaturedContent {
		mappedFeaturesArray = append(mappedFeaturesArray, model.Feature{
			Title:       fc.Title,
			Description: fc.Description,
			URI:         fc.URI,
		})
	}
	return mappedFeaturesArray
}

func getLatestReleases(rawReleases []release_calendar.Results) []model.Release {
	var latestReleases []model.Release

	// Removed canceled releases or unpublished releases
	for i := len(rawReleases) - 1; i >= 0; i-- {
		if rawReleases[i].Description.Cancelled || !rawReleases[i].Description.Published {
			rawReleases = append(rawReleases[:i], rawReleases[i+1:]...)
		}
	}

	// Reverse order
	sort.Slice(rawReleases, func(i, j int) bool {
		return rawReleases[i].Description.ReleaseDate.After(rawReleases[j].Description.ReleaseDate)
	})
	displayedReleases := 3
	for i := 0; i < displayedReleases; i++ {
		if len(rawReleases)-1 >= i {
			latestReleases = append(latestReleases, model.Release{
				Title:       rawReleases[i].Description.Title,
				URI:         rawReleases[i].URI,
				ReleaseDate: rawReleases[i].Description.ReleaseDate.Format("2 January 2006"),
			})
		}
	}
	return latestReleases
}

// getDataByPeriod returns the data for the time period set
func getDataByPeriod(datePeriod string, data zebedee.TimeseriesMainFigure) []zebedee.TimeseriesDataPoint {
	var mf []zebedee.TimeseriesDataPoint
	switch datePeriod {
	case "year":
		mf = data.Years
	case "quarter":
		mf = data.Quarters
	case "month":
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
