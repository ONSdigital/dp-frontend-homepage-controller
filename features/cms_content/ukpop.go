package content

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
)

// UkPop UK POP mocked data
var UkPop zebedee.TimeseriesMainFigure = zebedee.TimeseriesMainFigure{
	Description: zebedee.TimeseriesDescription{
		CDID:        "UKPOP",
		ReleaseDate: "2021-06-24T23:00:00.000Z",
		Unit:        "",
	},
	Months:   []zebedee.TimeseriesDataPoint{},
	Quarters: []zebedee.TimeseriesDataPoint{},
	Years: []zebedee.TimeseriesDataPoint{
		{
			Label: "2019",
			Value: "66796800",
		},
		{
			Label: "2020",
			Value: "67081000",
		},
	},
	RelatedDocuments: []zebedee.Link{
		{
			Title: "",
			URI:   "/peoplepopulationandcommunity/populationandmigration/populationestimates/bulletins/annualmidyearpopulationestimates/latest",
		},
	},
	URI: "/peoplepopulationandcommunity/populationandmigration/populationestimates/timeseries/ukpop/pop",
}
