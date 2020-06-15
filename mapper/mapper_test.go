package mapper

import (
	"context"
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"

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
				ImageURL:    "Featured content one imageURL",
			},
			{
				Title:       "Featured content two",
				Description: "Featured content two description",
				URI:         "Featured content two URI",
				ImageURL:    "Featured content two imageURL",
			},
			{
				Title:       "Featured content three",
				Description: "Featured content three description",
				URI:         "Featured content three URI",
				ImageURL:    "Featured content three imageURL",
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
			ImageURL:    "Featured content one imageURL",
		},
		{
			Title:       "Featured content two",
			Description: "Featured content two description",
			URI:         "Featured content two URI",
			ImageURL:    "Featured content two imageURL",
		},
		{
			Title:       "Featured content three",
			Description: "Featured content three description",
			URI:         "Featured content three URI",
			ImageURL:    "Featured content three imageURL",
		},
	}

	Convey("test homepage mapping works", t, func() {
		page := Homepage("en", mockedMainFigures, &mockedReleaseData, &mockedFeaturedContent)

		So(page.Type, ShouldEqual, "homepage")
		So(page.Data.MainFigures["test_id"].Figure, ShouldEqual, mockedMainFigure.Figure)
		So(page.Data.MainFigures["test_id"].TrendDescription, ShouldEqual, mockedMainFigure.TrendDescription)
		So(page.Data.MainFigures["test_id"].Trend, ShouldResemble, mockedMainFigure.Trend)
	})

	Convey("test main figures mapping works", t, func() {
		mockedTestData := mockedZebedeeData[0]
		mainFigures := MainFigure(ctx, "cdid", "month", mockedTestData)
		So(mainFigures.Date, ShouldEqual, "Feb 2020")
		So(mainFigures.Figure, ShouldEqual, "679.56")
		So(mainFigures.Trend.IsDown, ShouldEqual, false)
		So(mainFigures.Trend.IsUp, ShouldEqual, true)
		So(mainFigures.Trend.IsFlat, ShouldEqual, false)
		So(mainFigures.FigureURIs.Analysis, ShouldEqual, "test/uri/timeseries/123")
		So(mainFigures.FigureURIs.Data, ShouldEqual, "test/uri/timeseries/456")
		So(mainFigures.Unit, ShouldEqual, "%")
	})

	Convey("test featured content mapping aligns with expectations", t, func() {
		mockedTestData := mockedHomepageData
		featuredContent := FeaturedContent(mockedTestData)
		So(len(featuredContent), ShouldEqual, 3)
		for i := 0; i < len(featuredContent); i++ {
			So(featuredContent[i].Title, ShouldEqual, mockedTestData.FeaturedContent[i].Title)
			So(featuredContent[i].Description, ShouldEqual, mockedTestData.FeaturedContent[i].Description)
			So(featuredContent[i].URI, ShouldEqual, mockedTestData.FeaturedContent[i].URI)
			So(featuredContent[i].ImageURL, ShouldEqual, mockedTestData.FeaturedContent[i].ImageURL)
		}
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
		trendPositive := getTrend(5.8, 4.5672)
		trendNegative := getTrend(9.99, 10.21)
		trendFlat := getTrend(8.88, 8.88)
		So(trendPositive, ShouldResemble, model.Trend{IsUp: true, IsDown: false, IsFlat: false})
		So(trendNegative, ShouldResemble, model.Trend{IsUp: false, IsDown: true, IsFlat: false})
		So(trendFlat, ShouldResemble, model.Trend{IsUp: false, IsDown: false, IsFlat: true})
	})

	Convey("test getTrendDiffference returns the current string", t, func() {
		trendDescriptionPositive := getTrendDifference(10.55, 8.568, "million")
		trendDescriptionNegative := getTrendDifference(10.5, 18.7, "%")
		So(trendDescriptionPositive, ShouldEqual, "1.98million")
		So(trendDescriptionNegative, ShouldEqual, "-8.2pp")
	})

	Convey("test release calendar maps data correctly", t, func() {
		So(ReleaseCalendar(mockedBabbageRelease), ShouldResemble, &mockedReleaseData)
	})

}
