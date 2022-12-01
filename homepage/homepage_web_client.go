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

type WebClient struct {
	Updater
	cache            *cache.HomepageCache
	censusTopicCache *cache.TopicCache
	navigationCache  *cache.NavigationCache
	languages        []string
}

func NewWebClient(ctx context.Context, clients *Clients, updateInterval time.Duration, languages []string) (Clienter, error) {
	homepageCache, err := cache.NewHomepageCache(ctx, &updateInterval)
	if err != nil {
		log.Error(ctx, "failed to create new homepage cache", err, log.Data{"update_interval": updateInterval})
		return nil, err
	}

	return &WebClient{
		Updater: Updater{
			clients: clients,
		},
		cache:     homepageCache,
		languages: languages,
	}, nil
}

func (hwc *WebClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error) {
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

func (hwc *WebClient) AddNavigationCache(ctx context.Context, updateInterval time.Duration) error {
	navigationCache, err := cache.NewNavigationCache(ctx, &updateInterval)
	if err != nil {
		log.Error(ctx, "failed to create navigation cache", err, log.Data{"update_interval": updateInterval})
		return err
	}

	hwc.navigationCache = navigationCache

	return nil
}

func (hwc *WebClient) GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error) {
	if hwc.navigationCache == nil {
		log.Warn(ctx, "no-op navigation cache")
		return nil, nil
	}

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

func (hwc *WebClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
	for _, lang := range hwc.languages {
		langKey := getCachingKeyForLanguage(lang)
		hwc.cache.AddUpdateFunc(langKey, hwc.GetHomePageUpdateFor(ctx, "", "", lang))

		navigationlangKey := getCachingKeyForNavigationLanguage(lang)

		if hwc.navigationCache != nil {
			hwc.navigationCache.AddUpdateFunc(navigationlangKey, hwc.UpdateNavigationData(ctx, lang))
		}
	}

	go hwc.cache.StartUpdates(ctx, errorChannel)

	if hwc.navigationCache != nil {
		go hwc.navigationCache.StartUpdates(ctx, errorChannel)
	}
}

func (hwc *WebClient) Close() {
	hwc.cache.Close()

	if hwc.navigationCache != nil {
		hwc.navigationCache.Close()
	}
}
