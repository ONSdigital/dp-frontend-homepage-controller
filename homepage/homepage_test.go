package homepage

import (
	"context"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ReneKroon/ttlcache"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"

	"github.com/ONSdigital/dp-api-clients-go/image"
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
		Description:    zebedee.HomepageDescription{},
	}
	var mockedImageDownloadData = map[string]image.ImageDownload{
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
	expectedSuccessResponse := "<html><body><h1>Some HTML from renderer!</h1></body></html>"

	Convey("test homepage handler", t, func() {
		mockZebedeeClient := &ZebedeeClientMock{
			GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken, collectionID, lang, uri string) (m zebedee.TimeseriesMainFigure, err error) {
				return mockedZebedeeData[0], nil
			},
			GetHomepageContentFunc: func(ctx context.Context, userAuthToken, collectionID, lang, uri string) (m zebedee.HomepageContent, err error) {
				return mockedHomepageData, nil
			},
			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
				return nil
			},
		}

		mockBabbageClient := &BabbageClientMock{
			GetReleaseCalendarFunc: func(ctx context.Context, userAuthToken, fromDay, fromMonth, fromYear string) (m release_calendar.ReleaseCalendar, err error) {
				return mockedBabbageRelease, nil
			},
			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
				return nil
			},
		}

		mockImageClient := &ImageClientMock{
			GetDownloadVariantFunc: func(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, imageID, variant string) (image.ImageDownload, error) {
				return mockedImageDownloadData[imageID], nil
			},
			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
				return nil
			},
		}

		mockRenderClient := &RenderClientMock{
			DoFunc: func(string, []byte) ([]byte, error) {
				return []byte(expectedSuccessResponse), nil
			},
			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
				return nil
			},
		}

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Florence-Token", "testuser")
	rec := httptest.NewRecorder()
	router := mux.NewRouter()
	cache := ttlcache.NewCache()
	cache.SetTTL(10 * time.Millisecond)
	clients := &Clients{
		Renderer: mockRenderClient,
		Zebedee:  mockZebedeeClient,
		Babbage:  mockBabbageClient,
		ImageAPI: mockImageClient,
	}
	homepageClient := NewHomePagePublishingClient(clients)
	router.Path("/").HandlerFunc(Handler(homepageClient))

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
