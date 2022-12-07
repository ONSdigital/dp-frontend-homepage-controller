package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// Lf24 LF25 mocked data
var Lf24 zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "LF24",
		ReleaseDate: "2022-10-10T23:00:00.000Z",
		Unit:        "%\n",
	},
	Months: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 MAY-JUL",
			Value: "75.4",
		},
		{
			Label: "2022 JUN-AUG",
			Value: "75.5",
		},
	},
	Quarters: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 Q1",
			Value: "75.6",
		},
		{
			Label: "2022 Q2",
			Value: "75.5",
		},
	},
	Years: []zebedee.TimeseriesDataPoint{
		{
			Label: "2020",
			Value: "75.4",
		},
		{
			Label: "2021",
			Value: "75.1",
		},
	},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/bulletins/uklabourmarket/latest",
		},
	},
	URI: "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms",
}
