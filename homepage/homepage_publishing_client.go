package homepage

//go:generate moq -out mocks_homepage_publishing.go -pkg homepage . HomepageClienter

import (
	"context"
	"time"

	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
)

type HomepageClienter interface {
	AddNavigationCache(ctx context.Context, updateInterval time.Duration) error
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error)
	GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error)
	Close()
	StartBackgroundUpdate(ctx context.Context, errorChannel chan error)
}

type HomepagePublishingClient struct {
	HomepageUpdater
}

func NewHomePagePublishingClient(clients *Clients) HomepageClienter {
	return &HomepagePublishingClient{
		HomepageUpdater: HomepageUpdater{
			clients: clients,
		},
	}
}

func (hpc *HomepagePublishingClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error) {
	return hpc.GetHomePageUpdateFor(ctx, userAccessToken, collectionID, lang)()
}
func (hpc *HomepagePublishingClient) GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error) {
	return hpc.UpdateNavigationData(ctx, lang)(), nil
}
func (hpc *HomepagePublishingClient) AddNavigationCache(ctx context.Context, updateInterval time.Duration) error {
	return nil
}
func (hpc *HomepagePublishingClient) Close() {}
func (hpc *HomepagePublishingClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
}
