package routes

import (
	"context"

	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

type Clients struct {
	Renderer *renderer.Renderer
	Zebedee  *zebedee.Client
	Babbage  *release_calendar.Client
	ImageAPI *image.Client
}

// Init initialises routes for the service
func Init(ctx context.Context, r *mux.Router, hc health.HealthCheck, c Clients) {
	log.Event(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)
	r.StrictSlash(true).Path("/").Methods("GET").HandlerFunc(homepage.Handler(c.Renderer, c.Zebedee, c.Babbage, c.ImageAPI))
}
