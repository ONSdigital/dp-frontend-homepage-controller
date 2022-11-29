package homepage

//go:generate moq -out mock_clients.go -pkg homepage . ZebedeeClient ImageClient RenderClient
//go:generate moq -out mock_homepage_clienter.go -pkg homepage . Clienter

import (
	"context"
	"io"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	rendModel "github.com/ONSdigital/dp-renderer/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
)

// Clienter is an interface with methods required for Homepage client
type Clienter interface {
	AddNavigationCache(ctx context.Context, updateInterval time.Duration) error
	GetHomePage(ctx context.Context, userAccessToken, collectionID, lang string) (*model.HomepageData, error)
	GetNavigationData(ctx context.Context, lang string) (*topicModel.Navigation, error)
	Close()
	StartBackgroundUpdate(ctx context.Context, errorChannel chan error)
}

// ZebedeeClient is an interface with methods required for a zebedee client
type ZebedeeClient interface {
	GetTimeseriesMainFigure(ctx context.Context, userAuthToken, collectionID, lang, uri string) (m zebedee.TimeseriesMainFigure, err error)
	GetHomepageContent(ctx context.Context, userAccessToken, collectionID, lang, path string) (m zebedee.HomepageContent, err error)
	Checker(ctx context.Context, check *health.CheckState) error
}

// ImageClient is an interface with methods required for the Image API service client
type ImageClient interface {
	GetDownloadVariant(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, imageID, variant string) (m image.ImageDownload, err error)
	Checker(ctx context.Context, check *health.CheckState) error
}

// RenderClient is an interface with methods required for rendering a template from a page model
type RenderClient interface {
	BuildPage(w io.Writer, pageModel interface{}, templateName string)
	NewBasePageModel() rendModel.Page
}

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Zebedee  ZebedeeClient
	ImageAPI ImageClient
	Renderer RenderClient
	Topic    topicCli.Clienter
}
