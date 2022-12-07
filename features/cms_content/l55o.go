package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// L55o L55O mocked data
var L55o zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "L55O",
		ReleaseDate: "2022-09-13T23:00:00.000Z",
		Unit:        "%",
	},
	Months: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 JUL",
			Value: "8.8",
		},
		{
			Label: "2022 AUG",
			Value: "8.6",
		},
	},
	Quarters: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 Q1",
			Value: "5.5",
		},
		{
			Label: "2022 Q2",
			Value: "7.9",
		},
	},
	Years: []zebedee.TimeseriesDataPoint{
		{
			Label: "2020",
			Value: "1.0",
		},
		{
			Label: "2021",
			Value: "2.5",
		},
	},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/economy/inflationandpriceindices/bulletins/consumerpriceinflation/latest",
		},
	},
	URI: "/economy/inflationandpriceindices/timeseries/l55o/mm23",
}
