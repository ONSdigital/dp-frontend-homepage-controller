package census

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	homepage "github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	dphandlers "github.com/ONSdigital/dp-net/v3/handlers"
	"github.com/ONSdigital/log.go/v2/log"
)

// Handler handles requests to census endpoint
func Handler(cfg *config.Config, c cache.List, homepageClient homepage.Clienter, rend RenderClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handle(w, r, cfg, c, homepageClient, rend, accessToken, collectionID, lang)
	})
}

func handle(w http.ResponseWriter, req *http.Request, cfg *config.Config, c cache.List, homepageClient homepage.Clienter, rend RenderClient, userAccessToken, collectionID, lang string) {
	ctx := req.Context()
	navigationContent, err := homepageClient.GetNavigationData(ctx, lang)
	if err != nil {
		log.Error(ctx, "failed to get navigation data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cfg.EnableCensusTopicSubsection {
		var censusTopics *cache.Topic
		censusTopics, err = c.CensusTopic.GetCensusData(ctx)
		if err != nil {
			log.Error(ctx, "failed to get census topic data", err)
		}

		items := censusTopics.List.GetSubtopicItems()

		var availableItems []model.Topics

		if len(items) > 0 {
			for _, subTopics := range items {
				// do not map "Equalities" or "Historic census" since there are no results for these topics
				if subTopics.Title == "Equalities" || subTopics.Title == "Historic census" {
					continue
				}
				availableItems = append(availableItems, model.Topics{
					Topic: subTopics.Title,
					URL:   fmt.Sprintf("/search?topics=%s", subTopics.ID),
					ID:    subTopics.ID,
				})
			}
			// sort available items alphabetically
			sort.Slice(availableItems, func(i, j int) bool {
				return availableItems[i].Topic < availableItems[j].Topic
			})
		}

		log.Info(ctx, "census topics", log.Data{"census_topics": censusTopics, "items": items})

		homepageContent, err := homepageClient.GetHomePage(ctx, userAccessToken, collectionID, lang)
		if err != nil {
			log.Error(ctx, "HOMEPAGE_RESPONSE_FAILED. failed to get homepage html", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		basePage := rend.NewBasePageModel()
		m := mapper.Census(req, cfg, lang, basePage, navigationContent, homepageContent.EmergencyBanner, availableItems)

		rend.BuildPage(w, m, "census-topics")
	} else {
		homepageContent, err := homepageClient.GetHomePage(ctx, userAccessToken, collectionID, lang)
		if err != nil {
			log.Error(ctx, "HOMEPAGE_RESPONSE_FAILED. failed to get homepage html", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		basePage := rend.NewBasePageModel()
		m := mapper.CensusLegacy(req, cfg, lang, basePage, navigationContent, homepageContent.EmergencyBanner)

		rend.BuildPage(w, m, "census-first-results")
	}
	return
}
