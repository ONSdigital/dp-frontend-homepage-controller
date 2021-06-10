package homepage

import (
	"context"
	"github.com/ONSdigital/dp-frontend-homepage-controller/routes"
)

type HomepageClienter interface {
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (string, error)
	Close()
	StartBackgroundUpdate(ctx context.Context)
}

type HomepagePublishingClient struct {
	HomepageUpdater
}

func NewHomePagePublishingClient(clients *routes.Clients) HomepageClienter {
	return &HomepagePublishingClient{
		HomepageUpdater: HomepageUpdater{
			clients: clients,
		},
	}
}

func (hpc *HomepagePublishingClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (string, error) {
	return hpc.GetHomePageUpdateFor(ctx, userAccessToken, collectionID, lang)()
}

func (hpc *HomepagePublishingClient) Close() {}
func (hpc *HomepagePublishingClient) StartBackgroundUpdate(ctx context.Context) {}