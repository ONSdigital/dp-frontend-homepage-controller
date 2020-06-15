package homepage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {

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
	mockedHomepageData := zebedee.HomepageContent{
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
	expectedSuccessResponse := "<html><body><h1>Some HTML from renderer!</h1></body></html>"

	Convey("test homepage handler", t, func() {
		mockZebedeeClient := &ZebedeeClientMock{
			GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken, uri string) (m zebedee.TimeseriesMainFigure, err error) {
				return mockedZebedeeData[0], nil
			},
			GetHomepageContentFunc: func(ctx context.Context, userAuthToken, uri string) (m zebedee.HomepageContent, err error) {
				return mockedHomepageData, nil
			},
		}

		mockBabbageClient := &BabbageClientMock{
			GetReleaseCalendarFunc: func(ctx context.Context, userAuthToken, fromDay, fromMonth, fromYear string) (m release_calendar.ReleaseCalendar, err error) {
				return mockedBabbageRelease, nil
			},
		}

		mockRenderClient := &RenderClientMock{
			DoFunc: func(string, []byte) ([]byte, error) {
				return []byte(expectedSuccessResponse), nil
			},
		}

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("X-Florence-Token", "testuser")
		rec := httptest.NewRecorder()
		router := mux.NewRouter()
		router.Path("/").HandlerFunc(Handler(mockRenderClient, mockZebedeeClient, mockBabbageClient))

		Convey("returns 200 response", func() {
			router.ServeHTTP(rec, req)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Convey("renderer returns HTML body", func() {
			router.ServeHTTP(rec, req)
			response := rec.Body.String()
			So(response, ShouldEqual, expectedSuccessResponse)
		})
	})
}
