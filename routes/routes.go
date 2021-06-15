package routes

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"

	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

// Init initialises routes for the service
func Init(ctx context.Context, r *mux.Router, hcHandler func(http.ResponseWriter, *http.Request), homepageClient homepage.HomepageClienter) {
	log.Event(ctx, "adding routes", log.INFO)
	r.StrictSlash(true).Path("/health").HandlerFunc(hcHandler)
	r.StrictSlash(true).Path("/").Methods("GET").HandlerFunc(homepage.Handler(homepageClient))
}
