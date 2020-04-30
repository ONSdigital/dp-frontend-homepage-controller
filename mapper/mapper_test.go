package mapper

import (
	"context"
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/davecgh/go-spew/spew"
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

	var mockedMainFigures []model.MainFigure
	mockedMainFigures = append(mockedMainFigures, model.MainFigure{
		Date:             "Jun 2020",
		Figure:           "39.9",
		Trend:            model.Trend{IsUp: true},
		TrendDescription: "0.2%% on previous month",
		Unit:             "%",
		FigureURIs:       model.FigureURIs{Analysis: "test/uri/1/2/"},
	})

	Convey("test homepage mapping works", t, func() {
		page := Homepage(ctx, mockedMainFigures)

		So(page.Type, ShouldEqual, "homepage")
		So(page.Data.MainFigures[0], ShouldResemble, mockedMainFigures[0])
	})

	Convey("test main figures mapping works", t, func() {
		mainFigures := MainFigure(ctx, "months", mockedZebedeeData[0])
		spew.Dump(mainFigures)
		So(mainFigures.Date, ShouldEqual, "Feb 2020")
		So(mainFigures.Figure, ShouldEqual, "679.6")
		So(mainFigures.Trend.IsDown, ShouldEqual, false)
		So(mainFigures.Trend.IsUp, ShouldEqual, true)
		So(mainFigures.Trend.IsFlat, ShouldEqual, false)
		So(mainFigures.FigureURIs.Analysis, ShouldEqual, "test/uri/timeseries/123")
		So(mainFigures.FigureURIs.Data, ShouldEqual, "test/uri/timeseries/456")
		So(mainFigures.Unit, ShouldEqual, "%")
	})

	Convey("test getDataByPeriod returns correct data struct", t, func() {
		dataForYears := getDataByPeriod("years", mockedZebedeeData[0])
		dataForMonths := getDataByPeriod("months", mockedZebedeeData[0])
		dataForQuarters := getDataByPeriod("quarters", mockedZebedeeData[0])
		So(dataForYears, ShouldResemble, mockedZebedeeData[0].Years)
		So(dataForMonths, ShouldResemble, mockedZebedeeData[0].Months)
		So(dataForQuarters, ShouldResemble, mockedZebedeeData[0].Quarters)
	})

	Convey("test getTrend returns current struct of bools", t, func() {
		trendPositive := getTrend(5.8, 4.5672)
		trendNegative := getTrend(9.99, 10.21)
		trendFlat := getTrend(8.88, 8.88)
		spew.Dump(trendPositive, trendNegative, trendFlat)
		So(trendPositive, ShouldResemble, model.Trend{IsUp: true, IsDown: false, IsFlat: false})
		So(trendNegative, ShouldResemble, model.Trend{IsUp: false, IsDown: true, IsFlat: false})
		So(trendFlat, ShouldResemble, model.Trend{IsUp: false, IsDown: false, IsFlat: true})
	})

	Convey("test getTrendDescription returns the current string", t, func() {
		trendDescriptionPositive := getTrendDescription(10.55, 8.568, "million", "month")
		trendDescriptionNegative := getTrendDescription(10.5, 18.7, "%", "year")
		So(trendDescriptionPositive, ShouldEqual, "1.98million on previous month")
		So(trendDescriptionNegative, ShouldEqual, "-8.20% on previous year")
	})

}
