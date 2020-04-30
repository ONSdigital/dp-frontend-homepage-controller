package homepage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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

	expectedSuccessResponse := "<html><body><h1>Some HTML from renderer!</h1></body></html>"

	Convey("test homepage handler", t, func() {
		mockZebedeeClient := &ZebedeeClientMock{
			GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken, uri string) (m zebedee.TimeseriesMainFigure, err error) {
				return mockedZebedeeData[0], nil
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
		router.Path("/").HandlerFunc(Handler(mockRenderClient, mockZebedeeClient))

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

	Convey("get main figures list func returns correctly", t, func() {
		list := getMainFiguresList()
		So(list, ShouldNotBeEmpty)
	})

}
