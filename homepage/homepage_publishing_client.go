package homepage

import (
	"context"
)

type HomepageClienter interface {
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (string, error)
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

func (hpc *HomepagePublishingClient) GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (string, error) {
	return hpc.GetHomePageUpdateFor(ctx, userAccessToken, collectionID, lang)()
}

func (hpc *HomepagePublishingClient) Close() {}
func (hpc *HomepagePublishingClient) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
}
