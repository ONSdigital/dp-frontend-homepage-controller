package mapper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	topicModel "github.com/ONSdigital/dp-topic-api/models"

	"github.com/shopspring/decimal"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	mockedTrendData := TrendInfo{
		TrendFigure: zebedee.TimeseriesMainFigure{
			Description: zebedee.TimeseriesDescription{
				CDID:        "",
				Unit:        "%",
				ReleaseDate: "",
			},
			Years:    nil,
			Quarters: nil,
			Months: []zebedee.TimeseriesDataPoint{
				{
					Label: "2020 SEP",
					Value: "1.2",
				},
				{
					Label: "2020 OCT",
					Value: "1.2",
				},
			},
			RelatedDocuments: nil,
			URI:              "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/FUU8/lms",
		},
		IsTimeseriesForTrend: false,
		RetrieveTrendFailed:  false,
	}

	var mockedZebedeeData []zebedee.TimeseriesMainFigure
	mockedZebedeeData = append(mockedZebedeeData, zebedee.TimeseriesMainFigure{
		Months: []zebedee.TimeseriesDataPoint{
			{
				Value: "677.89",
				Label: "Jan 2020",
			},
			{
				Value: "679.56",
				Label: "Feb 2020",
			},
		},
		Years: []zebedee.TimeseriesDataPoint{
			{
				Value: "907.89",
				Label: "2020",
			},
			{
				Value: "1009.56",
				Label: "2021",
			},
		},
		Quarters: []zebedee.TimeseriesDataPoint{
			{
				Value: "13.97",
				Label: "Q1",
			},
			{
				Value: "14.68",
				Label: "Q2",
			},
		},
		RelatedDocuments: []zebedee.Related{
			{
				Title: "Related thing",
				URI:   "test/uri/timeseries/123",
			},
		},
		Description: zebedee.TimeseriesDescription{
			CDID: "LOLZ",
			Unit: "%",
		},
		URI: "test/uri/timeseries/456",
	})
	mockedMainFigures := make(map[string]*model.MainFigure)
	mockedMainFigure := model.MainFigure{
		Date:             "Jun 2020",
		Figure:           "39.9",
		Trend:            model.Trend{IsUp: true},
		TrendDescription: "0.2%% on previous month",
		Unit:             "%",
		FigureURIs:       model.FigureURIs{Analysis: "test/uri/1/2/"},
	}
	mockedMainFigures["test_id"] = &mockedMainFigure

	var mockedHomepageData = zebedee.HomepageContent{
		Intro: zebedee.Intro{
			Title:    "Welcome to the Office for National Statistics",
			Markdown: "Markdown text here",
		},
		FeaturedContent: []zebedee.Featured{
			{
				Title:       "Featured content one",
				Description: "Featured content one description",
				URI:         "Featured content one URI",
				ImageID:     "123",
			},
			{
				Title:       "Featured content two",
				Description: "Featured content two description",
				URI:         "Featured content two URI",
				ImageID:     "456",
			},
			{
				Title:       "Featured content three",
				Description: "Featured content three description",
				URI:         "Featured content three URI",
				ImageID:     "",
			},
		},
		AroundONS: []zebedee.Featured{
			{
				Title:       "Around ONS one",
				Description: "Around ONS one description",
				URI:         "Around ONS one URI",
				ImageID:     "123",
			},
			{
				Title:       "Around ONS two",
				Description: "Around ONS two description",
				URI:         "Around ONS two URI",
				ImageID:     "",
			},
		},
		ServiceMessage: "",
		URI:            "/",
		Type:           "",
		Description: zebedee.HomepageDescription{
			Title:           "Homepage description title",
			Summary:         "Homepage description summary",
			Keywords:        []string{"keyword one", "keyword two"},
			MetaDescription: "",
			Unit:            "",
			PreUnit:         "",
			Source:          "",
		},
	}
	var mockedFeaturedContent = []model.Feature{
		{
			Title:       "Featured content one",
			Description: "Featured content one description",
			URI:         "Featured content one URI",
			ImageURL:    "path/to/123.png",
		},
		{
			Title:       "Featured content two",
			Description: "Featured content two description",
			URI:         "Featured content two URI",
			ImageURL:    "path/to/456.png",
		},
		{
			Title:       "Featured content three",
			Description: "Featured content three description",
			URI:         "Featured content three URI",
			ImageURL:    "",
		},
	}

	var mockedAroundONS = []model.Feature{
		{
			Title:       "Around ONS one",
			Description: "Around ONS one description",
			URI:         "Around ONS one URI",
			ImageURL:    "path/to/123.png",
		},
		{
			Title:       "Around ONS two",
			Description: "Around ONS two description",
			URI:         "Around ONS two URI",
			ImageURL:    "",
		},
	}
	var mockedImageDownloadsData = map[string]image.ImageDownload{
		"123": {
			Size:  111111,
			Type:  "blah",
			Href:  "http://www.example.com/images/123/original.png",
			State: "completed",
		},
		"456": {
			Size:  111111,
			Type:  "blah",
			Href:  "http://www.example.com/images/456/original.png",
			State: "completed",
		},
	}

	testFigure1, _ := decimal.NewFromString("12.345")
	testFigure2, _ := decimal.NewFromString("8.90")
	testFigure3, _ := decimal.NewFromString("100.2")
	testFigure4, _ := decimal.NewFromString("101.423")
	testFigure5, _ := decimal.NewFromString("88.8888")

	serviceMessage := "Test service message"

	emergencyBanner := zebedee.EmergencyBanner{
		Type:        "notable_death",
		Title:       "Emergency banner title",
		Description: "Emergency banner description",
		URI:         "www.google.com",
		LinkText:    "More info",
	}

	mockedNavigationData := &topicModel.Navigation{
		Description: "A list of topical areas and their subtopics in english to generate the website navbar.",
		Links: &topicModel.TopicLinks{
			Self: &topicModel.LinkObject{
				HRef: "/navigation",
			},
		},
		Items: &[]topicModel.TopicNonReferential{
			{
				Title:       "Business, industry and trade",
				Description: "Activities of businesses and industry in the UK, including data on the production and trade of goods and services, sales by retailers, characteristics of businesses, the construction and manufacturing sectors, and international trade.",
				Name:        "business-industry-and-trade",
				Label:       "Business, industry and trade",
				Links: &topicModel.TopicLinks{
					Self: &topicModel.LinkObject{
						ID:   "businessindustryandtrade",
						HRef: "/topics/businessindustryandtrade",
					},
				},
				Uri: "/businessindustryandtrade",
				SubtopicItems: &[]topicModel.TopicNonReferential{
					{
						Title:       "Business",
						Description: "UK businesses registered for VAT and PAYE with regional breakdowns, including data on size (employment and turnover) and activity (type of industry), research and development, and business services.",
						Name:        "business",
						Label:       "Business",
						Links: &topicModel.TopicLinks{
							Self: &topicModel.LinkObject{
								ID:   "business",
								HRef: "/topics/business",
							},
						},
						Uri: "/businessindustryandtrade/business",
					},
				},
			},
		},
	}

	basePage := coreModel.NewPage("path/to/assets", "site-domain")
	Convey("test homepage mapping works", t, func() {
		page := Homepage(config.Config{}, "en", basePage, mockedMainFigures, &mockedFeaturedContent, &mockedAroundONS, serviceMessage, emergencyBanner, mockedNavigationData)
		So(page.SiteDomain, ShouldResemble, basePage.SiteDomain)
		So(page.PatternLibraryAssetsPath, ShouldResemble, basePage.PatternLibraryAssetsPath)
		So(page.Type, ShouldEqual, "homepage")
		So(page.Data.MainFigures["test_id"].Figure, ShouldEqual, mockedMainFigure.Figure)
		So(page.Data.MainFigures["test_id"].TrendDescription, ShouldEqual, mockedMainFigure.TrendDescription)
		So(page.Data.MainFigures["test_id"].Trend, ShouldResemble, mockedMainFigure.Trend)
		So(page.Data.HasFeaturedContent, ShouldEqual, true)
		So(page.Data.HasMainFigures, ShouldEqual, true)
		So(page.Data.Featured, ShouldHaveLength, 3)
		So(page.Data.AroundONS, ShouldHaveLength, 2)
		So(page.EmergencyBanner.Title, ShouldEqual, emergencyBanner.Title)
		So(page.EmergencyBanner.Type, ShouldEqual, "notable-death")
		So(page.EmergencyBanner.Description, ShouldEqual, emergencyBanner.Description)
		So(page.EmergencyBanner.URI, ShouldEqual, emergencyBanner.URI)
		So(page.EmergencyBanner.LinkText, ShouldEqual, emergencyBanner.LinkText)
	})

	Convey("empty emergency banner content, banner does not map", t, func() {
		page := Homepage(config.Config{}, "en", basePage, mockedMainFigures, &mockedFeaturedContent, &mockedAroundONS, serviceMessage, zebedee.EmergencyBanner{}, mockedNavigationData)

		So(page.EmergencyBanner.Title, ShouldBeBlank)
		So(page.EmergencyBanner.Type, ShouldBeBlank)
		So(page.EmergencyBanner.Description, ShouldBeBlank)
		So(page.EmergencyBanner.URI, ShouldBeBlank)
		So(page.EmergencyBanner.LinkText, ShouldBeBlank)
	})

	Convey("test main figures mapping works", t, func() {
		mockedTestData := mockedZebedeeData[0]
		mainFigures := MainFigure(ctx, "cdid", PeriodMonth, PeriodMonth, mockedTestData, mockedTrendData)
		So(mainFigures.Date, ShouldEqual, "Feb 2020")
		So(mainFigures.Figure, ShouldEqual, "679.6")
		So(mainFigures.Trend.IsDown, ShouldEqual, false)
		So(mainFigures.Trend.IsUp, ShouldEqual, true)
		So(mainFigures.Trend.IsFlat, ShouldEqual, false)
		So(mainFigures.FigureURIs.Analysis, ShouldEqual, "test/uri/timeseries/123")
		So(mainFigures.FigureURIs.Data, ShouldEqual, "test/uri/timeseries/456")
		So(mainFigures.Unit, ShouldEqual, "%")
	})

	Convey("test FeaturedContent", t, func() {
		Convey("FeaturedContent handles when no homepage data is passed in", func() {
			mockedTestData := zebedee.HomepageContent{}
			mockedImageTestData := map[string]image.ImageDownload{}
			featuredContent := FeaturedContent(mockedTestData, mockedImageTestData)
			So(featuredContent, ShouldBeEmpty)
		})

		Convey("FeaturedContent maps mock data to page model correctly", func() {
			mockedTestData := mockedHomepageData
			mockedImageTestData := mockedImageDownloadsData
			featuredContent := FeaturedContent(mockedTestData, mockedImageTestData)
			So(len(featuredContent), ShouldEqual, 3)
			for i := 0; i < len(featuredContent); i++ {
				So(featuredContent[i].Title, ShouldEqual, mockedTestData.FeaturedContent[i].Title)
				So(featuredContent[i].Description, ShouldEqual, mockedTestData.FeaturedContent[i].Description)
				So(featuredContent[i].URI, ShouldEqual, mockedTestData.FeaturedContent[i].URI)
				if featuredContent[i].ImageURL != "" {
					So(featuredContent[i].ImageURL, ShouldEqual, mockedImageDownloadsData[mockedTestData.FeaturedContent[i].ImageID].Href)
				}
			}
			So(featuredContent[2].ImageURL, ShouldEqual, "")
		})
	})

	Convey("test AroundONS", t, func() {
		Convey("AroundONS handles when no homepage data is passed in", func() {
			mockedTestData := zebedee.HomepageContent{}
			mockedImageTestData := map[string]image.ImageDownload{}
			aroundONS := AroundONS(mockedTestData, mockedImageTestData)
			So(aroundONS, ShouldBeNil)
		})

		Convey("AroundONS handles when no AroundONS data is passed in", func() {
			mockedTestData := mockedHomepageData
			mockedTestData.AroundONS = nil
			mockedImageTestData := map[string]image.ImageDownload{}
			aroundONS := AroundONS(mockedTestData, mockedImageTestData)
			So(aroundONS, ShouldBeNil)
		})

		Convey("AroundONS handles when AroundONS data with missing fields is passed in", func() {
			mockedTestData := mockedHomepageData
			mockedTestData.AroundONS[1].URI = ""
			mockedTestData.AroundONS[1].Title = ""
			mockedTestData.AroundONS[1].Description = ""
			mockedImageTestData := map[string]image.ImageDownload{}
			aroundONS := AroundONS(mockedTestData, mockedImageTestData)
			So(len(aroundONS), ShouldEqual, 2)
			for i := 0; i < len(aroundONS); i++ {
				So(aroundONS[i].Title, ShouldEqual, mockedTestData.AroundONS[i].Title)
				So(aroundONS[i].Description, ShouldEqual, mockedTestData.AroundONS[i].Description)
				So(aroundONS[i].URI, ShouldEqual, mockedTestData.AroundONS[i].URI)
				if aroundONS[i].ImageURL != "" {
					So(aroundONS[i].ImageURL, ShouldEqual, mockedImageDownloadsData[mockedTestData.AroundONS[i].ImageID].Href)
				}
			}
			So(aroundONS[1].ImageURL, ShouldEqual, "")
		})

		Convey("FeaturedContent maps mock data to page model correctly", func() {
			mockedTestData := mockedHomepageData
			mockedImageTestData := mockedImageDownloadsData
			aroundONS := AroundONS(mockedTestData, mockedImageTestData)
			So(len(aroundONS), ShouldEqual, 2)
			for i := 0; i < len(aroundONS); i++ {
				So(aroundONS[i].Title, ShouldEqual, mockedTestData.AroundONS[i].Title)
				So(aroundONS[i].Description, ShouldEqual, mockedTestData.AroundONS[i].Description)
				So(aroundONS[i].URI, ShouldEqual, mockedTestData.AroundONS[i].URI)
				if aroundONS[i].ImageURL != "" {
					So(aroundONS[i].ImageURL, ShouldEqual, mockedImageDownloadsData[mockedTestData.AroundONS[i].ImageID].Href)
				}
			}
			So(aroundONS[1].ImageURL, ShouldEqual, "")
		})
	})

	Convey("test getDataByPeriod returns correct data struct", t, func() {
		dataForYears := getDataByPeriod("year", mockedZebedeeData[0])
		dataForMonths := getDataByPeriod("month", mockedZebedeeData[0])
		dataForQuarters := getDataByPeriod("quarter", mockedZebedeeData[0])
		So(dataForYears, ShouldResemble, mockedZebedeeData[0].Years)
		So(dataForMonths, ShouldResemble, mockedZebedeeData[0].Months)
		So(dataForQuarters, ShouldResemble, mockedZebedeeData[0].Quarters)
	})

	Convey("test getTrend returns current struct of bools", t, func() {
		trendPositive := getTrend(testFigure1, testFigure2)
		trendNegative := getTrend(testFigure3, testFigure4)
		trendFlat := getTrend(testFigure5, testFigure5)
		So(trendPositive, ShouldResemble, model.Trend{IsUp: true, IsDown: false, IsFlat: false})
		So(trendNegative, ShouldResemble, model.Trend{IsUp: false, IsDown: true, IsFlat: false})
		So(trendFlat, ShouldResemble, model.Trend{IsUp: false, IsDown: false, IsFlat: true})
	})

	Convey("test getTrendDifference returns the current string", t, func() {
		diff := testFigure1.Sub(testFigure2)
		trendDescriptionPositive := formatTrend(diff, "million")
		diff = testFigure3.Sub(testFigure4)
		trendDescriptionNegative := formatTrend(diff, "%")
		So(trendDescriptionPositive, ShouldEqual, "3.4million")
		So(trendDescriptionNegative, ShouldEqual, "-1.2pp")
	})

	Convey("test formatCommas returns correctly formatted numbers as string", t, func() {
		So(formatCommas("12.3"), ShouldEqual, "12.3")
		So(formatCommas("64890980.7"), ShouldEqual, "64,890,980.7")
		So(formatCommas("1000.2"), ShouldEqual, "1,000.2")
		So(formatCommas("88789.1"), ShouldEqual, "88,789.1")
	})

	Convey("test getDifferenceOffset returns correct offset value", t, func() {
		So(getDifferenceOffset(PeriodMonth, PeriodMonth), ShouldEqual, 1)
		So(getDifferenceOffset(PeriodYear, PeriodYear), ShouldEqual, 1)
		So(getDifferenceOffset(PeriodQuarter, PeriodYear), ShouldEqual, 4)
		So(getDifferenceOffset(PeriodMonth, PeriodYear), ShouldEqual, 12)
		So(getDifferenceOffset(PeriodMonth, PeriodQuarter), ShouldEqual, 3)
	})

	Convey("test graceful degradation state is properly mapped", t, func() {
		var mockedNoFeaturedContent []model.Feature
		var mockedNoMainFigures = make(map[string]*model.MainFigure)

		gracefulDegradationPage := Homepage(config.Config{}, "en", basePage, mockedNoMainFigures, &mockedNoFeaturedContent, &mockedAroundONS, serviceMessage, emergencyBanner, mockedNavigationData)

		So(gracefulDegradationPage.Data.HasFeaturedContent, ShouldEqual, false)
		So(gracefulDegradationPage.Data.HasMainFigures, ShouldEqual, false)
	})
}

func TestUnitMapCookiesPreferences(t *testing.T) {
	req := httptest.NewRequest("", "/", nil)
	pageModel := coreModel.Page{
		CookiesPreferencesSet: false,
		CookiesPolicy: coreModel.CookiesPolicy{
			Essential: false,
			Usage:     false,
		},
	}

	Convey("cookies preferences initialise as false", t, func() {
		So(pageModel.CookiesPreferencesSet, ShouldBeFalse)
		So(pageModel.CookiesPolicy.Essential, ShouldBeFalse)
		So(pageModel.CookiesPolicy.Usage, ShouldBeFalse)
	})

	Convey("cookie preferences map to page model", t, func() {
		req.AddCookie(&http.Cookie{Name: "cookies_preferences_set", Value: "true"})
		req.AddCookie(&http.Cookie{Name: "cookies_policy", Value: "%7B%22essential%22%3Atrue%2C%22usage%22%3Atrue%7D"})
		mapCookiePreferences(req, &pageModel.CookiesPreferencesSet, &pageModel.CookiesPolicy)
		So(pageModel.CookiesPreferencesSet, ShouldBeTrue)
		So(pageModel.CookiesPolicy.Essential, ShouldBeTrue)
		So(pageModel.CookiesPolicy.Usage, ShouldBeTrue)
	})
}
