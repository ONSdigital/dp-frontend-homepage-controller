package homepage

//go:generate moq -out client_mock.go -pkg client . ZebedeeClient BabbageClient ImageClient RenderClient

import (
	"context"
<<<<<<< HEAD
	"io"
=======
>>>>>>> develop

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-renderer/model"
)

//go:generate moq -out mocks_test.go -pkg homepage . ZebedeeClient ImageClient RenderClient

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
	NewBasePageModel() model.Page
}

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Zebedee  ZebedeeClient
	ImageAPI ImageClient
	Renderer RenderClient
}
