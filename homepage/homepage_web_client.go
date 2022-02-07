package homepage

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-frontend-homepage-controller/dpcache"
	"net/http"
	"time"
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

func (hwc *HomepageWebClient) GetHomePage(ctx context.Context, w http.ResponseWriter, rend RenderClient, userAccessToken, collectionID, lang string) (string, error) {
	homePageCachedHTML, ok := hwc.cache.Get(getCachingKeyForLanguage(lang))
	if ok {
		return fmt.Sprintf("%s", homePageCachedHTML), nil
	}

	return "", fmt.Errorf("failed to read homepage from cache for: %s", getCachingKeyForLanguage(lang))

}

func getCachingKeyForLanguage(lang string) string {
	return fmt.Sprintf("%s___%s", dpcache.HomepageCacheKey, lang)
}

func (hwc *HomepageWebClient) StartBackgroundUpdate(ctx context.Context, w http.ResponseWriter, rend RenderClient, errorChannel chan error) {
	for _, lang := range hwc.languages {
		langKey := getCachingKeyForLanguage(lang)
		hwc.cache.AddUpdateFunc(langKey, hwc.GetHomePageUpdateFor(ctx, w, rend, "", "", lang))
	}

	hwc.cache.StartUpdates(ctx, errorChannel)
}
func (hwc *HomepageWebClient) Close() {
	hwc.cache.Close()
}
