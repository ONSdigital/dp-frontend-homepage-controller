package mapper

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"github.com/shopspring/decimal"
)

const (
	// PeriodYear is the string value for year time period
	PeriodYear = "year"
	// PeriodQuarter is the string value for quarter time period
	PeriodQuarter = "quarter"
	// PeriodMonth is the string value for month time period
	PeriodMonth = "month"
)

// decimalPointDisplayThreshold is a number where we no longer want to
// display numbers with decimals. This number was a product decision
var decimalPointDisplayThreshold = decimal.NewFromInt(1000)

// Homepage maps data to our homepage frontend model
func Homepage(localeCode string, mainFigures map[string]*model.MainFigure, releaseCal *model.ReleaseCalendar, featuredContent *[]model.Feature, serviceMessage string) model.Page {
	var page model.Page
	page.Type = "homepage"
	page.Metadata.Title = "Home"
	page.Data.HasFeaturedContent = hasFeaturedContent(featuredContent)
	page.Data.HasMainFigures = hasMainFigures(mainFigures)
	page.HasJSONLD = true
	page.ServiceMessage = serviceMessage
	page.Language = localeCode
	page.Data.MainFigures = mainFigures
	page.Data.ReleaseCalendar = *releaseCal
	page.Data.Featured = *featuredContent
	return page
}

// MainFigure maps a single main figure object
func MainFigure(ctx context.Context, id, datePeriod, differenceInterval string, figure zebedee.TimeseriesMainFigure) *model.MainFigure {
	var mf model.MainFigure

	mf.ID = id

	mfData := getDataByPeriod(datePeriod, figure)
	previousDataOffset := getDifferenceOffset(datePeriod, differenceInterval) + 1
	if len(mfData) < previousDataOffset {
		log.Event(ctx, "error: too few observations in timeseries array", log.ERROR, log.Error(errors.New("too few observations in timeseries array")))
		return &mf
	}
	latestDataIndex := len(mfData) - 1
	previousDataIndex := len(mfData) - previousDataOffset
	latestData := mfData[latestDataIndex]
	previousData := mfData[previousDataIndex]
	latestFigure, err := decimal.NewFromString(latestData.Value)
	if err != nil {
		log.Event(ctx, "error getting trend description: error converting string to decimal type", log.ERROR, log.Error(err))
		return &mf
	}
	previousFigure, err := decimal.NewFromString(previousData.Value)
	if err != nil {
		log.Event(ctx, "error getting trend description: error converting string to decimal type", log.ERROR, log.Error(err))
		return &mf
	}

	var formattedLatestFigure string
	if latestFigure.GreaterThanOrEqual(decimalPointDisplayThreshold) {
		formattedLatestFigure = latestFigure.String()
	} else {
		formattedLatestFigure = latestFigure.StringFixed(1)
	}

	mf.Figure = formatCommas(formattedLatestFigure)
	mf.Date = latestData.Label
	mf.Unit = figure.Description.Unit
	mf.Trend = getTrend(latestFigure, previousFigure)
	mf.Trend.Difference = getTrendDifference(latestFigure, previousFigure, figure.Description.Unit)
	mf.Trend.Period = differenceInterval
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
func FeaturedContent(homepageData zebedee.HomepageContent, images map[string]image.ImageDownload) []model.Feature {
	var mappedFeaturesContent []model.Feature
	for _, fc := range homepageData.FeaturedContent {
		mappedFeaturesContent = append(mappedFeaturesContent, model.Feature{
			Title:       fc.Title,
			Description: fc.Description,
			URI:         fc.URI,
			ImageURL:    images[fc.ImageID].Href,
		})
	}
	return mappedFeaturesContent
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
	case PeriodYear:
		mf = data.Years
	case PeriodQuarter:
		mf = data.Quarters
	case PeriodMonth:
		mf = data.Months
	default:
		mf = []zebedee.TimeseriesDataPoint{}
	}
	return mf
}

// getTrend returns trend boolean value based on latest and previous figures
func getTrend(latest, previous decimal.Decimal) model.Trend {
	if latest.GreaterThan(previous) {
		return model.Trend{IsUp: true}
	}

	if latest.LessThan(previous) {
		return model.Trend{IsDown: true}
	}

	if latest.Equal(previous) {
		return model.Trend{IsFlat: true}
	}
	return model.Trend{}
}

// getTrendDifference returns string value of the difference between latest and previous
func getTrendDifference(latest, previous decimal.Decimal, unit string) string {
	var trendUnit string
	switch unit {
	case "%":
		trendUnit = "pp"
	default:
		trendUnit = unit
	}
	diff := latest.Sub(previous)
	diffStr := diff.StringFixed(1)
	//formattedDiff := humanize.CommafWithDigits(diff, 2)
	return fmt.Sprintf("%v%v", diffStr, trendUnit)
}

// formats large numbers to contain comma separators e.g. 1000000 => 1,000,000
func formatCommas(str string) string {
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

// getDifferenceOffset works out a numeric value that represents the
// offset of values to compare from data in a timeseries
func getDifferenceOffset(period, interval string) int {
	if period == interval {
		return 1
	}

	if period == PeriodQuarter && interval == PeriodYear {
		return 4
	}

	if period == PeriodMonth {
		if interval == PeriodYear {
			return 12
		}
		if interval == PeriodQuarter {
			return 3
		}
	}
	// only gets here if incomparable options are chosen in code
	panic("unable to get difference offset from choosen period and interval values")
}

func hasFeaturedContent(featuredContent *[]model.Feature) bool {
	if len(*featuredContent) > 0 {
		return true
	}
	return false
}

func hasMainFigures(mainFigures map[string]*model.MainFigure) bool {
	for _, value := range mainFigures {
		if value.Figure != "" {
			return true
		}
	}
	return false
}
