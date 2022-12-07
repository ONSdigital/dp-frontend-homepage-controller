package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// Mgsx MGSX mocked data
var Mgsx zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "MGSX",
		ReleaseDate: "2022-10-10T23:00:00.000Z",
		Unit:        "%",
	},
	Months: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 MAY-JUL",
			Value: "3.6",
		},
		{
			Label: "2022 JUN-AUG",
			Value: "3.5",
		},
	},
	Quarters: []zebedee.TimeseriesDataPoint{
		{
			Label: "2022 Q1",
			Value: "3.7",
		},
		{
			Label: "2022 Q2",
			Value: "3.8",
		},
	},
	Years: []zebedee.TimeseriesDataPoint{
		{
			Label: "2020",
			Value: "4.6",
		},
		{
			Label: "2021",
			Value: "4.5",
		},
	},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/bulletins/uklabourmarket/latest",
		},
	},
	URI: "/employmentandlabourmarket/peoplenotinwork/unemployment/timeseries/mgsx/lms",
}
