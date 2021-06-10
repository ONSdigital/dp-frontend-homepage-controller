package routes

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"

	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Renderer *renderer.Renderer
	Zebedee  *zebedee.Client
	Babbage  *release_calendar.Client
	ImageAPI *image.Client
}

// Init initialises routes for the service
func Init(ctx context.Context, r *mux.Router, hcHandler func(http.ResponseWriter, *http.Request), homepageClient homepage.HomepageClienter) {
	log.Event(ctx, "adding routes", log.INFO)
	r.StrictSlash(true).Path("/health").HandlerFunc(hcHandler)
	r.StrictSlash(true).Path("/").Methods("GET").HandlerFunc(homepage.Handler(homepageClient))
}
