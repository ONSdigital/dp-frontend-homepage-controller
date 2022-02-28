package homepage

//go:generate moq -out mocks_homepage_publishing.go -pkg homepage . HomepageClienter

import (
	"context"

	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
)

type HomepageClienter interface {
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error)
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

func (hpc *HomepagePublishingClient) Close() {}
func (hpc *HomepagePublishingClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
}
