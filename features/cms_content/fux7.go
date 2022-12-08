package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// Fux7 Fux7 mocked data
var Fux7 zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "FUX7",
		ReleaseDate: "2021-01-20T00:00:00.000Z",
		Unit:        "",
	},
	Months: []zebedee.TimeseriesDataPoint{
		{
			Label: "2020 AUG",
			Value: "-0.8",
		},
		{
			Label: "2020 SEP",
			Value: "-0.9",
		},
	},
	Quarters: []zebedee.TimeseriesDataPoint{},
	Years:    []zebedee.TimeseriesDataPoint{},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/datasets/labourmarketstatistics",
		},
	},
	URI: "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fux7/lms",
}
