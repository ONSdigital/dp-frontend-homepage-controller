package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// Fuu8 Fuu8 mocked data
var Fuu8 zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "FUU8",
		ReleaseDate: "2021-01-20T00:00:00.000Z",
		Unit:        "",
	},
	Months: []zebedee.TimeseriesDataPoint{
		{
			Label: "2020 AUG",
			Value: "0.9",
		},
		{
			Label: "2020 SEP",
			Value: "1.2",
		},
	},
	Quarters: []zebedee.TimeseriesDataPoint{},
	Years:    []zebedee.TimeseriesDataPoint{},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/bulletins/wrong",
		},
	},
	URI: "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/fuu8/lms",
}
