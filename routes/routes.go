package routes

import (
	"context"
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"

	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// Init initialises routes for the service
func Init(ctx context.Context, r *mux.Router, hcHandler func(http.ResponseWriter, *http.Request), homepageClient homepage.HomepageClienter) {
	log.Info(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(hcHandler)
	r.StrictSlash(true).Path("/").Methods("GET").HandlerFunc(homepage.Handler(homepageClient))
}
