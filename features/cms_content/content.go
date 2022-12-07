package content

import "github.com/ONSdigital/dp-api-clients-go/v2/zebedee"

// Zebedee is a map of mocked zebedee content,
// where keys are zebedee URLs
// and values are the corresponding mocked response
var Zebedee map[string]zebedee.TimeseriesMainFigure = map[string]zebedee.TimeseriesMainFigure{
	UkPop.URI:   UkPop,
	IhyqQna.URI: IhyqQna,
	IhyqPn2.URI: IhyqPn2,
	L55o.URI:    L55o,
}
