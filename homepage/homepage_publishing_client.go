package homepage

//go:generate moq -out mocks_homepage_publishing.go -pkg homepage . HomepageClienter

import (
	"context"
	"fmt"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/log.go/v2/log"
)

type Clienter interface {
	AddNavigationCache(ctx context.Context, updateInterval time.Duration) error
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error)
	GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error)
	Close()
	StartBackgroundUpdate(ctx context.Context, errorChannel chan error)
}

type PublishingClient struct {
	Updater
	navigationCache *cache.NavigationCache
	languages       []string
}

func NewHomePagePublishingClient(ctx context.Context, clients *Clients, languages []string) Clienter {
	return &PublishingClient{
		Updater: Updater{
			clients: clients,
		},
		languages: languages,
	}
}

func (hpc *PublishingClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error) {
	return hpc.GetHomePageUpdateFor(ctx, userAccessToken, collectionID, lang)()
}

func (hpc *PublishingClient) AddNavigationCache(ctx context.Context, updateInterval time.Duration) error {
	navigationCache, err := cache.NewNavigationCache(ctx, &updateInterval)
	if err != nil {
		log.Error(ctx, "failed to create navigation cache", err, log.Data{"update_interval": updateInterval})
		return err
	}

	hpc.navigationCache = navigationCache

	return nil
}

func (hpc *PublishingClient) GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error) {
	if hpc.navigationCache == nil {
		log.Warn(ctx, "no-op navigation cache")
		return nil, nil
	}

	navigationData, ok := hpc.navigationCache.Get(getCachingKeyForNavigationLanguage(lang))
	if ok {
		n, ok := navigationData.(*topicModel.Navigation)
		if ok {
			return n, nil
		}
	}

	return nil, fmt.Errorf("failed to read navigation data from cache for: %s", getCachingKeyForNavigationLanguage(lang))
}

func (hpc *PublishingClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
	if hpc.navigationCache == nil {
		return
	}

	for _, lang := range hpc.languages {
		navigationlangKey := getCachingKeyForNavigationLanguage(lang)

		hpc.navigationCache.AddUpdateFunc(navigationlangKey, hpc.UpdateNavigationData(ctx, lang))
	}

	go hpc.navigationCache.StartUpdates(ctx, errorChannel)
}

func (hpc *PublishingClient) Close() {
	if hpc.navigationCache != nil {
		hpc.navigationCache.Close()
	}
}
