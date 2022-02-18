package homepage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ReneKroon/ttlcache"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
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
