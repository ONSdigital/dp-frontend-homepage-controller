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
			Label: "2021 JAN-MAR",
			Value: "74.7",
		},
		{
			Label: "2021 FEB-APR",
			Value: "74.7",
		},
		{
			Label: "2021 MAR-MAY",
			Value: "74.8",
		},
		{
			Label: "2021 APR-JUN",
			Value: "75.0",
		},
		{
			Label: "2021 MAY-JUL",
			Value: "75.1",
		},
		{
			Label: "2021 JUN-AUG",
			Value: "75.2",
		},
		{
			Label: "2021 JUL-SEP",
			Value: "75.3",
		},
		{
			Label: "2021 AUG-OCT",
			Value: "75.4",
		},
		{
			Label: "2021 SEP-NOV",
			Value: "75.4",
		},
		{
			Label: "2021 OCT-DEC",
			Value: "75.5",
		},
		{
			Label: "2021 NOV-JAN",
			Value: "75.4",
		},
		{
			Label: "2021 DEC-FEB",
			Value: "75.5",
		},
		{
			Label: "2022 JAN-MAR",
			Value: "75.6",
		},
		{
			Label: "2022 FEB-APR",
			Value: "75.6",
		},
		{
			Label: "2022 MAR-MAY",
			Value: "75.9",
		},
		{
			Label: "2022 APR-JUN",
			Value: "75.5",
		},
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
			Label: "2021 Q1",
			Value: "74.7",
		},
		{
			Label: "2021 Q2",
			Value: "75.0",
		},
		{
			Label: "2021 Q3",
			Value: "75.3",
		},
		{
			Label: "2021 Q4",
			Value: "75.5",
		},
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
