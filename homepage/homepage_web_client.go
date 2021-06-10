package homepage

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-frontend-homepage-controller/dpcache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/routes"
	"time"
)

type HomepageWebClient struct {
	HomepageUpdater
	cache     dpcache.DpCacher
	languages []string
}

func NewHomePageWebClient(clients *routes.Clients, updateInterval time.Duration, languages []string) HomepageClienter {
	return &HomepageWebClient{
		HomepageUpdater: HomepageUpdater{
			clients: clients,
		},
		cache: dpcache.NewDpCache(updateInterval),
	}
}

func (hwc *HomepageWebClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (string, error) {
	homePageCachedHTML, ok := hwc.cache.Get(getCachingKeyForLanguage(lang))
	if ok {
		homePageString, convertible := homePageCachedHTML.([]byte)
		if convertible {
			return string(homePageString), nil
		}
	}

	return "", fmt.Errorf("failed to read homepage from cache")

}

func getCachingKeyForLanguage(lang string) string {
	return fmt.Sprintf("%-%s", dpcache.HomepageCacheKey, lang)
}

func (hwc *HomepageWebClient) StartBackgroundUpdate(ctx context.Context) {
	for _, lang := range hwc.languages {
		langKey := getCachingKeyForLanguage(lang)
		hwc.cache.AddUpdateFunc(langKey, hwc.GetHomePageUpdateFor(ctx, "", "", lang))
	}

	go hwc.cache.StartUpdates(ctx)
}
func (hwc *HomepageWebClient) Close() {
	hwc.cache.Close()
}