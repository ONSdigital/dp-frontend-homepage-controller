package mapper

import (
	"context"
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	"github.com/shopspring/decimal"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

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
	mockedDescriptions := [5]*release_calendar.Description{
		{
			ReleaseDate: time.Now().AddDate(0, 0, -1),
			Cancelled:   false,
			Published:   true,
			Title:       "Foo",
		}, {
			ReleaseDate: time.Now().AddDate(0, 0, -2),
			Cancelled:   false,
			Published:   true,
			Title:       "bAr",
		}, {
			ReleaseDate: time.Now().AddDate(0, 0, -3),
			Cancelled:   false,
			Published:   true,
			Title:       "BAZ",
		}, {
			ReleaseDate: time.Now().AddDate(0, 0, -4),
			Cancelled:   false,
			Published:   true,
			Title:       "qux",
		}, {
			ReleaseDate: time.Now().AddDate(0, 0, -5),
			Cancelled:   true,
			Published:   false,
			Title:       "Qu ux",
		},
	}
	mockedResults := []release_calendar.Results{
		{
			Type:        "release",
			Description: mockedDescriptions[0],
			SearchBoost: nil,
			URI:         "/releases/foo",
		},
		{
			Type:        "release",
			Description: mockedDescriptions[1],
			SearchBoost: nil,
			URI:         "/releases/bar",
		},
		{
			Type:        "release",
			Description: mockedDescriptions[2],
			SearchBoost: nil,
			URI:         "/releases/baz",
		},
		{
			Type:        "release",
			Description: mockedDescriptions[3],
			SearchBoost: nil,
			URI:         "/releases/qux",
		},
		{
			Type:        "release",
			Description: mockedDescriptions[4],
			SearchBoost: nil,
			URI:         "/releases/quux",
		},
	}
	mockedBabbageRelease := release_calendar.ReleaseCalendar{
		Type:     "list",
		ListType: "releasecalendar",
		URI:      "/releasecalendar/data",
		Result: release_calendar.Result{
			NumberOfResults: 5,
			Took:            3,
			Results:         &mockedResults,
			Suggestions:     nil,
			DocCounts:       struct{}{},
			SortBy:          "release_date",
		},
	}
	var mockedReleaseData = model.ReleaseCalendar{
		Releases: []model.Release{
			{
				Title:       "Foo",
				URI:         "/releases/foo",
				ReleaseDate: time.Now().AddDate(0, 0, -1).Format("2 January 2006"),
			},
			{
				Title:       "bAr",
				URI:         "/releases/bar",
				ReleaseDate: time.Now().AddDate(0, 0, -2).Format("2 January 2006"),
			},
			{
				Title:       "BAZ",
				URI:         "/releases/baz",
				ReleaseDate: time.Now().AddDate(0, 0, -3).Format("2 January 2006"),
			},
		},
		NumberOfReleases:                 4,
		NumberOfOtherReleasesInSevenDays: 1,
	}
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
	var mockedImageData = []image.Image{
		{
			Id:           "123",
			CollectionId: "",
			Filename:     "123.png",
			Downloads: map[string]image.ImageDownload{
				"png": {
					Href: "path/to/123.png",
				},
			},
		},
		{
			Id:           "456",
			CollectionId: "",
			Filename:     "456.png",
			Downloads: map[string]image.ImageDownload{
				"png": {
					Href: "path/to/456.png",
				},
			},
		},
	}

	testFigure1, _ := decimal.NewFromString("12.345")
	testFigure2, _ := decimal.NewFromString("8.90")
	testFigure3, _ := decimal.NewFromString("100.2")
	testFigure4, _ := decimal.NewFromString("101.423")
	testFigure5, _ := decimal.NewFromString("88.8888")

	Convey("test homepage mapping works", t, func() {
		page := Homepage("en", mockedMainFigures, &mockedReleaseData, &mockedFeaturedContent)

		So(page.Type, ShouldEqual, "homepage")
		So(page.Data.MainFigures["test_id"].Figure, ShouldEqual, mockedMainFigure.Figure)
		So(page.Data.MainFigures["test_id"].TrendDescription, ShouldEqual, mockedMainFigure.TrendDescription)
		So(page.Data.MainFigures["test_id"].Trend, ShouldResemble, mockedMainFigure.Trend)
		So(len(page.Data.Featured), ShouldEqual, 3)
	})

	Convey("test main figures mapping works", t, func() {
		mockedTestData := mockedZebedeeData[0]
		mainFigures := MainFigure(ctx, "cdid", "month", mockedTestData)
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
			mockedImageTestData := []image.Image{}
			featuredContent := FeaturedContent(mockedTestData, mockedImageTestData)
			So(featuredContent, ShouldBeNil)
		})

		Convey("FeaturedContent maps mock data to page model correctly", func() {
			mockedTestData := mockedHomepageData
			mockedImageTestData := mockedImageData
			featuredContent := FeaturedContent(mockedTestData, mockedImageTestData)
			So(len(featuredContent), ShouldEqual, 3)
			for i := 0; i < len(featuredContent); i++ {
				So(featuredContent[i].Title, ShouldEqual, mockedTestData.FeaturedContent[i].Title)
				So(featuredContent[i].Description, ShouldEqual, mockedTestData.FeaturedContent[i].Description)
				So(featuredContent[i].URI, ShouldEqual, mockedTestData.FeaturedContent[i].URI)
				if featuredContent[i].ImageURL != "" {
					So(featuredContent[i].ImageURL, ShouldEqual, mockedImageData[i].Downloads["png"].Href)
				}
			}
			So(featuredContent[2].ImageURL, ShouldEqual, "")
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

	Convey("test getTrendDiffference returns the current string", t, func() {
		trendDescriptionPositive := getTrendDifference(testFigure1, testFigure2, "million")
		trendDescriptionNegative := getTrendDifference(testFigure3, testFigure4, "%")
		So(trendDescriptionPositive, ShouldEqual, "3.4million")
		So(trendDescriptionNegative, ShouldEqual, "-1.2pp")
	})

	Convey("test release calendar maps data correctly", t, func() {
		So(ReleaseCalendar(mockedBabbageRelease), ShouldResemble, &mockedReleaseData)
	})

	Convey("test formatCommas returns correctly formatted numbers as string", t, func() {
		So(formatCommas("12.3"), ShouldEqual, "12.3")
		So(formatCommas("64890980.7"), ShouldEqual, "64,890,980.7")
		So(formatCommas("1000.2"), ShouldEqual, "1,000.2")
		So(formatCommas("88789.1"), ShouldEqual, "88,789.1")
	})

}
