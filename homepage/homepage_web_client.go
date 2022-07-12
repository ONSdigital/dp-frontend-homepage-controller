package homepage

import (
	"context"
	"fmt"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/log.go/v2/log"
)

type HomepageWebClient struct {
	HomepageUpdater
	cache           *cache.HomepageCache
	navigationCache *cache.NavigationCache
	languages       []string
}

func NewHomePageWebClient(ctx context.Context, clients *Clients, updateInterval time.Duration, languages []string) (HomepageClienter, error) {
	homepageCache, err := cache.NewHomepageCache(ctx, &updateInterval)
	if err != nil {
		log.Error(ctx, "failed to create new homepage cache", err, log.Data{"update_interval": updateInterval})
		return nil, err
	}

	return &HomepageWebClient{
		HomepageUpdater: HomepageUpdater{
			clients: clients,
		},
		cache:     homepageCache,
		languages: languages,
	}, nil
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
	return fmt.Sprintf("%s___%s", cache.HomepageCacheKey, lang)
}

func (hwc *HomepageWebClient) AddNavigationCache(ctx context.Context, updateInterval time.Duration) error {
	navigationCache, err := cache.NewNavigationCache(ctx, &updateInterval)
	if err != nil {
		log.Error(ctx, "failed to create navigation cache", err, log.Data{"update_interval": updateInterval})
		return err
	}

	hwc.navigationCache = navigationCache

	return nil
}

func (hwc *HomepageWebClient) GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error) {
	navigationData, ok := hwc.navigationCache.Get(getCachingKeyForNavigationLanguage(lang))
	if ok {
		n, ok := navigationData.(*topicModel.Navigation)
		if ok {
			return n, nil
		}
	}

	return nil, fmt.Errorf("failed to read navigation data from cache for: %s", getCachingKeyForNavigationLanguage(lang))
}

func getCachingKeyForNavigationLanguage(lang string) string {
	return fmt.Sprintf("%s___%s", cache.NavigationCacheKey, lang)
}

func (hwc *HomepageWebClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
	for _, lang := range hwc.languages {
		langKey := getCachingKeyForLanguage(lang)
		hwc.cache.AddUpdateFunc(langKey, hwc.GetHomePageUpdateFor(ctx, "", "", lang))

		navigationlangKey := getCachingKeyForNavigationLanguage(lang)
		hwc.navigationCache.AddUpdateFunc(navigationlangKey, hwc.UpdateNavigationData(ctx, lang))
	}

	go hwc.cache.StartUpdates(ctx, errorChannel)
	go hwc.navigationCache.StartUpdates(ctx, errorChannel)
}

func (hwc *HomepageWebClient) Close() {
	hwc.cache.Close()
	hwc.navigationCache.Close()
}
