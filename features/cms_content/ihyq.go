package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// IhyqQna IHYQ QNA mocked data
var IhyqQna zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "IHYQ",
		ReleaseDate: "2022-09-29T23:00:00.000Z",
		Unit:        "%",
	},
	Months: []zebedee.TimeseriesDataPoint{},
	Quarters: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 Q1",
			Value: "0.7",
		},
		{
			Label: "2022 Q2",
			Value: "0.2",
		},
	},
	Years: []zebedee.TimeseriesDataPoint{},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/economy/grossdomesticproductgdp/bulletins/quarterlynationalaccounts/latest",
		},
	},
	URI: "/economy/grossdomesticproductgdp/timeseries/ihyq/qna",
}

// IhyqPn2 IHYQ PN2 mocked data
var IhyqPn2 zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "IHYQ",
		ReleaseDate: "2022-08-11T23:00:00.000Z",
		Unit:        "%",
	},
	Months: []zebedee.TimeseriesDataPoint{},
	Quarters: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 Q1",
			Value: "0.8",
		},
		{
			Label: "2022 Q2",
			Value: "-0.1",
		},
	},
	Years: []zebedee.TimeseriesDataPoint{},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/economy/grossdomesticproductgdp/bulletins/gdpfirstquarterlyestimateuk/latest",
		},
	},
	URI: "/economy/grossdomesticproductgdp/timeseries/ihyq/pn2",
}
