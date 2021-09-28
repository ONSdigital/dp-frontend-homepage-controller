package mapper

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
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

// TrendInfo, stores all trend data for processing
type TrendInfo struct {
	TrendFigure          zebedee.TimeseriesMainFigure
	IsTimeseriesForTrend bool
	RetrieveTrendFailed  bool
}

// decimalPointDisplayThreshold is a number where we no longer want to
// display numbers with decimals. This number was a product decision
var decimalPointDisplayThreshold = decimal.NewFromInt(1000)

// Homepage maps data to our homepage frontend model
func Homepage(localeCode string, mainFigures map[string]*model.MainFigure, releaseCal *model.ReleaseCalendar, featuredContent *[]model.Feature, aroundONS *[]model.Feature,  serviceMessage string) model.Page {
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
	page.Data.AroundONS = *aroundONS
	return page
}

// MainFigure maps a single main figure object
func MainFigure(ctx context.Context, id, datePeriod, differenceInterval string, figure zebedee.TimeseriesMainFigure, trendInfo TrendInfo) *model.MainFigure {
	var mf model.MainFigure
	mf.ShowTrend = !trendInfo.RetrieveTrendFailed
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
	var trend = model.Trend{}
	timeseriesTrendSuccess := trendInfo.IsTimeseriesForTrend && !trendInfo.RetrieveTrendFailed
	if timeseriesTrendSuccess || !trendInfo.IsTimeseriesForTrend {
		trend = resolveTrend(ctx, latestFigure, previousFigure, figure.Description.Unit, differenceInterval, datePeriod, trendInfo)
	}
	mf.Trend = trend
	mf.Figure = formatCommas(formattedLatestFigure)
	mf.Date = latestData.Label
	mf.Unit = figure.Description.Unit

	if len(figure.RelatedDocuments) > 0 {
		mf.FigureURIs.Analysis = figure.RelatedDocuments[0].URI
	}
	mf.FigureURIs.Data = figure.URI
	return &mf
}

func ReleaseCalendar(rawReleaseCalendar release_calendar.ReleaseCalendar) *model.ReleaseCalendar {
	rc := model.ReleaseCalendar{
		Releases:                         []model.Release{},
		NumberOfReleases:                 0,
		NumberOfOtherReleasesInSevenDays: 0,
	}
	// No releases found
	if rawReleaseCalendar.Result.Results == nil {
		return &rc
	}
	releaseResults := *rawReleaseCalendar.Result.Results
	numReleasesScheduled := rawReleaseCalendar.Result.NumberOfResults

	for i := len(releaseResults) - 1; i >= 0; i-- {
		if releaseResults[i].Description.Cancelled || !releaseResults[i].Description.Published {
			numReleasesScheduled--
		}
	}

	latestReleases := getLatestReleases(releaseResults)
	rc.Releases = latestReleases
	rc.NumberOfReleases = numReleasesScheduled
	rc.NumberOfOtherReleasesInSevenDays = numReleasesScheduled - len(latestReleases)
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

// AroundONS takes the homepageContent as returned from the client and returns an array of featured content
func AroundONS(homepageData zebedee.HomepageContent, images map[string]image.ImageDownload) []model.Feature {
	var mappedAroundONS []model.Feature
	if len(homepageData.AroundONS) > 0 {
		for _, fc := range homepageData.AroundONS {
			mappedAroundONS = append(mappedAroundONS, model.Feature{
				Title:       fc.Title,
				Description: fc.Description,
				URI:         fc.URI,
				ImageURL:    images[fc.ImageID].Href,
			})
		}
	}
	return mappedAroundONS
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

// resolveTrend will use the timeseries trend value if present, otherwise will fallback to working it out.
// If there is a timeseries but it fails then it will show no trend indication, this is by design.
func resolveTrend(ctx context.Context, latestMF, previousMF decimal.Decimal, unit, differenceInterval, datePeriod string, trendInfo TrendInfo) model.Trend {
	var trend model.Trend

	// Timeseries data available that explicitly states the trend
	if trendInfo.IsTimeseriesForTrend {
		trendData := getDataByPeriod(datePeriod, trendInfo.TrendFigure)
		latestDataIndex := len(trendData) - 1
		latestData := trendData[latestDataIndex]
		latestTrendFigure, err := decimal.NewFromString(latestData.Value)
		if err != nil {
			log.Event(ctx, "error converting string to decimal type", log.ERROR, log.Error(err))
			return trend
		}
		trend = getTrend(latestTrendFigure, decimal.NewFromInt(0))
		trend.Difference = formatTrend(latestTrendFigure, unit)
		trend.Period = differenceInterval
	} else {
		// No timeseries that explicitly states the trend, therefore work it out
		trend = getTrend(latestMF, previousMF)
		diff := latestMF.Sub(previousMF)
		trend.Difference = formatTrend(diff, unit)
		trend.Period = differenceInterval
	}
	return trend
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

// formatTrend returns correctly formatted string value for a given decimal difference
func formatTrend(diff decimal.Decimal, unit string) string {
	var trendUnit string
	switch unit {
	case "%":
		trendUnit = "pp"
	default:
		trendUnit = unit
	}
	diffStr := diff.StringFixed(1)
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
	return len(*featuredContent) > 0
}

func hasMainFigures(mainFigures map[string]*model.MainFigure) bool {
	for _, value := range mainFigures {
		if value.Figure != "" {
			return true
		}
	}
	return false
}
