package homepage

import (
	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
)

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Renderer *renderer.Renderer
	Zebedee  *zebedee.Client
	Babbage  *release_calendar.Client
	ImageAPI *image.Client
}
