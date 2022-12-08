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
			Label: "2021 JAN",
			Value: "0.9",
		},
		{
			Label: "2021 FEB",
			Value: "0.7",
		},
		{
			Label: "2021 MAR",
			Value: "1.0",
		},
		{
			Label: "2021 APR",
			Value: "1.6",
		},
		{
			Label: "2021 MAY",
			Value: "2.1",
		},
		{
			Label: "2021 JUN",
			Value: "2.4",
		},
		{
			Label: "2021 JUL",
			Value: "2.1",
		},
		{
			Label: "2021 AUG",
			Value: "3.0",
		},
		{
			Label: "2021 SEP",
			Value: "2.9",
		},
		{
			Label: "2021 OCT",
			Value: "3.8",
		},
		{
			Label: "2021 NOV",
			Value: "4.6",
		},
		{
			Label: "2021 DEC",
			Value: "4.8",
		},
		{
			Label: "2022 JAN",
			Value: "4.9",
		},
		{
			Label: "2022 FEB",
			Value: "5.5",
		},
		{
			Label: "2022 MAR",
			Value: "6.2",
		},
		{
			Label: "2022 APR",
			Value: "7.8",
		},
		{
			Label: "2022 MAY",
			Value: "7.9",
		},
		{
			Label: "2022 JUN",
			Value: "8.2",
		},
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
			Label: "2021 Q1",
			Value: "0.9",
		},
		{
			Label: "2021 Q2",
			Value: "2.1",
		},
		{
			Label: "2021 Q3",
			Value: "2.7",
		},
		{
			Label: "2021 Q4",
			Value: "4.4",
		},
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
