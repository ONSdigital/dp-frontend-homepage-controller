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
			Label: "2021 JAN-MAR",
			Value: "4.9",
		},
		{
			Label: "2021 FEB-APR",
			Value: "4.9",
		},
		{
			Label: "2021 MAR-MAY",
			Value: "4.9",
		},
		{
			Label: "2021 APR-JUN",
			Value: "4.7",
		},
		{
			Label: "2021 MAY-JUL",
			Value: "4.6",
		},
		{
			Label: "2021 JUN-AUG",
			Value: "4.4",
		},
		{
			Label: "2021 JUL-SEP",
			Value: "4.3",
		},
		{
			Label: "2021 AUG-OCT",
			Value: "4.2",
		},
		{
			Label: "2021 SEP-NOV",
			Value: "4.1",
		},
		{
			Label: "2021 OCT-DEC",
			Value: "4.0",
		},
		{
			Label: "2021 NOV-JAN",
			Value: "4.0",
		},
		{
			Label: "2021 DEC-FEB",
			Value: "3.8",
		},
		{
			Label: "2022 JAN-MAR",
			Value: "3.7",
		},
		{
			Label: "2022 FEB-APR",
			Value: "3.8",
		},
		{
			Label: "2022 MAR-MAY",
			Value: "3.8",
		},
		{
			Label: "2022 APR-JUN",
			Value: "3.8",
		},
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
			Label: "2021 Q1",
			Value: "4.9",
		},
		{
			Label: "2021 Q2",
			Value: "4.7",
		},
		{
			Label: "2021 Q3",
			Value: "4.3",
		},
		{
			Label: "2021 Q4",
			Value: "4.0",
		},
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
