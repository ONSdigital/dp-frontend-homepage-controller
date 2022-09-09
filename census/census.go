package census

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/helper"
	homepage "github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
	"github.com/ONSdigital/log.go/v2/log"
)

// Handler handles requests to census endpoint
func Handler(cfg *config.Config, homepageClient homepage.Clienter, rend RenderClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, cfg, homepageClient, rend, accessToken, collectionID, lang)
	})
}

func handle(w http.ResponseWriter, req *http.Request, cfg *config.Config, homepageClient homepage.Clienter, rend RenderClient, userAccessToken, collectionID, lang string) {
	ctx := req.Context()
	navigationContent, err := homepageClient.GetNavigationData(ctx, lang)
	if err != nil {
		log.Error(ctx, "failed to get navigation data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	homepageContent, err := homepageClient.GetHomePage(ctx, userAccessToken, collectionID, lang)
	if err != nil {
		log.Error(ctx, "HOMEPAGE_RESPONSE_FAILED. failed to get homepage html", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	basePage := rend.NewBasePageModel()
	m := mapper.Census(req, cfg, lang, basePage, navigationContent, homepageContent.EmergencyBanner)

	enableCensusResults := helper.CheckTime(ctx, cfg.CensusFirstResults)

	if enableCensusResults {
		rend.BuildPage(w, m, "census-first-results")
	} else {
		rend.BuildPage(w, m, "census")
	}

	return
}
