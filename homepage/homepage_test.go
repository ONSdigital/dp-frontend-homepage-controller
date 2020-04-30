package homepage

import (
	"context"
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {
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

	Convey("test homepage handler", t, func() {
		mockZebedeeClient := &ZebedeeClientMock{
			GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken, uri string) (m zebedee.TimeseriesMainFigure, err error) {
				return mockedZebedeeData[0], nil
			},
		}
	})

}
