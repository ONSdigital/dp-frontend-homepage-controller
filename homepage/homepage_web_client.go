package homepage

import (
	"context"
	"fmt"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/dpcache"
	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
)

type HomepageWebClient struct {
	HomepageUpdater
	cache     dpcache.DpCacher
	languages []string
}

func NewHomePageWebClient(clients *Clients, updateInterval time.Duration, languages []string) HomepageClienter {
	return &HomepageWebClient{
		HomepageUpdater: HomepageUpdater{
			clients: clients,
		},
		cache:     dpcache.NewDpCache(updateInterval),
		languages: languages,
	}
}

func (hwc *HomepageWebClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error) {
	homepageData, ok := hwc.cache.Get(getCachingKeyForLanguage(lang))
	if ok {
		h, ok := homepageData.(*model.HomepageData)
		if ok {
			return h, nil
		}
	}

	return nil, fmt.Errorf("failed to read homepage from cache for: %s", getCachingKeyForLanguage(lang))

}

func getCachingKeyForLanguage(lang string) string {
	return fmt.Sprintf("%s___%s", dpcache.HomepageCacheKey, lang)
}

func (hwc *HomepageWebClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
	for _, lang := range hwc.languages {
		langKey := getCachingKeyForLanguage(lang)
		hwc.cache.AddUpdateFunc(langKey, hwc.GetHomePageUpdateFor(ctx, "", "", lang))
	}

	hwc.cache.StartUpdates(ctx, errorChannel)
}
func (hwc *HomepageWebClient) Close() {
	hwc.cache.Close()
}
