package mapper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-cookies/cookies"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/shopspring/decimal"
)

const (
	// PeriodYear is the string value for year time period
	PeriodYear = "year"
	// PeriodQuarter is the string value for quarter time period
	PeriodQuarter = "quarter"
	// PeriodMonth is the string value for month time period
	PeriodMonth    = "month"
	CensusURI      = "/census"
	CensusPageType = "census"
	CensusTitle    = "Census"
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
func Homepage(cfg config.Config, localeCode string, basePage coreModel.Page, mainFigures map[string]*model.MainFigure, featuredContent, aroundONS *[]model.Feature, serviceMessage string, emergencyBannerContent zebedee.EmergencyBanner, navigationContent *topicModel.Navigation) model.Page {
	page := model.Page{
		Data: model.Homepage{},
		Page: basePage,
	}

	page.Type = "homepage"
	page.Metadata.Title = "Home"
	page.Data.HasFeaturedContent = hasFeaturedContent(featuredContent)
	page.Data.HasMainFigures = hasMainFigures(mainFigures)
	page.HasJSONLD = true
	page.ServiceMessage = serviceMessage
	page.Language = localeCode
	page.Data.MainFigures = mainFigures
	page.EmergencyBanner = mapEmergencyBanner(emergencyBannerContent)
	page.FeatureFlags.SixteensVersion = "8466e33"
	if navigationContent != nil {
		page.NavigationContent = mapNavigationContent(*navigationContent)
	}

	if aroundONS != nil {
		page.Data.AroundONS = *aroundONS
	}

	if featuredContent != nil {
		page.Data.Featured = *featuredContent
	}

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
		log.Error(ctx, "error: too few observations in timeseries array", errors.New("too few observations in timeseries array"))
		return &mf
	}

	latestDataIndex := len(mfData) - 1
	previousDataIndex := len(mfData) - previousDataOffset
	latestData := mfData[latestDataIndex]
	previousData := mfData[previousDataIndex]
	latestFigure, err := decimal.NewFromString(latestData.Value)
	if err != nil {
		log.Error(ctx, "error getting trend description: error converting string to decimal type", err)
		return &mf
	}
	previousFigure, err := decimal.NewFromString(previousData.Value)
	if err != nil {
		log.Error(ctx, "error getting trend description: error converting string to decimal type", err)
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

// FeaturedContent takes the homepageContent as returned from the client and returns an array of featured content
func FeaturedContent(homepageData zebedee.HomepageContent, images map[string]image.ImageDownload) []model.Feature {
	mappedFeaturesContent := make([]model.Feature, len(homepageData.FeaturedContent))
	for index, fc := range homepageData.FeaturedContent {
		mappedFeaturesContent[index] = model.Feature{
			Title:       fc.Title,
			Description: fc.Description,
			URI:         fc.URI,
			ImageURL:    images[fc.ImageID].Href,
		}
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

// mapNavigationContent takes navigationContent as returned from the client and returns information needed for the navigation bar
func mapNavigationContent(navigationContent topicModel.Navigation) []coreModel.NavigationItem {
	var mappedNavigationContent []coreModel.NavigationItem
	if navigationContent.Items != nil {
		for _, rootContent := range *navigationContent.Items {
			var subItems []coreModel.NavigationItem
			if rootContent.SubtopicItems != nil {
				for _, subtopicContent := range *rootContent.SubtopicItems {
					subItems = append(subItems, coreModel.NavigationItem{
						Uri:   subtopicContent.Uri,
						Label: subtopicContent.Label,
					})
				}
			}
			mappedNavigationContent = append(mappedNavigationContent, coreModel.NavigationItem{
				Uri:      rootContent.Uri,
				Label:    rootContent.Label,
				SubItems: subItems,
			})
		}
	}
	return mappedNavigationContent
}

func mapEmergencyBanner(bannerData zebedee.EmergencyBanner) coreModel.EmergencyBanner {
	var mappedEmergencyBanner coreModel.EmergencyBanner
	emptyBannerObj := zebedee.EmergencyBanner{}
	if bannerData != emptyBannerObj {
		mappedEmergencyBanner.Title = bannerData.Title
		mappedEmergencyBanner.Type = strings.Replace(bannerData.Type, "_", "-", -1)
		mappedEmergencyBanner.Description = bannerData.Description
		mappedEmergencyBanner.URI = bannerData.URI
		mappedEmergencyBanner.LinkText = bannerData.LinkText
	}
	return mappedEmergencyBanner
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
			log.Error(ctx, "error converting string to decimal type", err)
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
	if featuredContent == nil {
		return false
	}

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

// Census maps data to our census frontend model
func Census(req *http.Request, cfg *config.Config, localeCode string, basePage coreModel.Page, navigationContent *topicModel.Navigation, emergencyBannerContent zebedee.EmergencyBanner, censusSubTopics []model.Topics) model.CensusPage {
	page := model.CensusPage{
		Page: basePage,
		Data: model.Census{},
	}

	mapCookiePreferences(req, &page.Page.CookiesPreferencesSet, &page.Page.CookiesPolicy)
	page.URI = CensusURI
	page.Type = CensusPageType
	page.Metadata.Title = CensusTitle
	page.Language = localeCode
	page.PatternLibraryAssetsPath = cfg.PatternLibraryAssetsPath
	page.EmergencyBanner = mapEmergencyBanner(emergencyBannerContent)
	page.Data.EnableCensusTopicSubsection = cfg.EnableCensusTopicSubsection
	page.Data.CensusSearchTopicID = cfg.CensusTopicID
	page.Data.EnableGetDataCard = cfg.EnableGetDataCard
	page.Data.AvailableTopics = censusSubTopics

	if navigationContent != nil {
		page.NavigationContent = mapNavigationContent(*navigationContent)
	}

	return page
}

func CensusLegacy(req *http.Request, cfg *config.Config, localeCode string, basePage coreModel.Page, navigationContent *topicModel.Navigation, emergencyBannerContent zebedee.EmergencyBanner) model.CensusPage {
	page := model.CensusPage{
		Page: basePage,
		Data: model.Census{},
	}

	mapCookiePreferences(req, &page.Page.CookiesPreferencesSet, &page.Page.CookiesPolicy)
	page.URI = CensusURI
	page.Type = CensusPageType
	page.Metadata.Title = CensusTitle
	page.Language = localeCode
	page.PatternLibraryAssetsPath = cfg.PatternLibraryAssetsPath
	page.EmergencyBanner = mapEmergencyBanner(emergencyBannerContent)
	page.Data.EnableCensusTopicSubsection = cfg.EnableCensusTopicSubsection
	page.Data.CensusSearchTopicID = cfg.CensusTopicID
	page.Data.EnableGetDataCard = cfg.EnableGetDataCard

	if navigationContent != nil {
		page.NavigationContent = mapNavigationContent(*navigationContent)
	}

	return page
}

// mapCookiePreferences reads cookie policy and preferences cookies and then maps the values to the page model
func mapCookiePreferences(req *http.Request, preferencesIsSet *bool, policy *coreModel.CookiesPolicy) {
	preferencesCookie := cookies.GetCookiePreferences(req)
	*preferencesIsSet = preferencesCookie.IsPreferenceSet
	*policy = coreModel.CookiesPolicy{
		Essential: preferencesCookie.Policy.Essential,
		Usage:     preferencesCookie.Policy.Usage,
	}
}
