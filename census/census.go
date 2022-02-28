package census

import (
	"net/http"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
)

// Handler handles requests to census endpoint
func Handler(cfg *config.Config, rend RenderClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, cfg, rend, accessToken, collectionID, lang)
	})
}

func handle(w http.ResponseWriter, req *http.Request, cfg *config.Config, rend RenderClient, userAccessToken, collectionID, lang string) {
	basePage := rend.NewBasePageModel()
	m := mapper.Census(req, cfg, lang, basePage)

	rend.BuildPage(w, m, "census")

	return
}
