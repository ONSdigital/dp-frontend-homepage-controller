package routes

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/census"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"

	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// Init initialises routes for the service
func Init(ctx context.Context, r *mux.Router, cfg *config.Config, c cache.List, hcHandler func(http.ResponseWriter, *http.Request), homepageClient homepage.Clienter, renderClient homepage.RenderClient) {
	log.Info(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(hcHandler)
	r.StrictSlash(true).Path("/census").Methods("GET").HandlerFunc(census.Handler(cfg, c, homepageClient, renderClient))
	r.StrictSlash(true).Path("/").Methods("GET").HandlerFunc(homepage.Handler(cfg, homepageClient, renderClient))
}
