package homepage

//go:generate moq -out mocks/mocks_test.go -pkg mock . ZebedeeClient BabbageClient ImageClient RenderClient

import (
	"context"
	"io"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-renderer/model"
)

// ZebedeeClient is an interface with methods required for a zebedee client
type ZebedeeClient interface {
	GetTimeseriesMainFigure(ctx context.Context, userAuthToken, collectionID, lang, uri string) (m zebedee.TimeseriesMainFigure, err error)
	GetHomepageContent(ctx context.Context, userAccessToken, collectionID, lang, path string) (m zebedee.HomepageContent, err error)
	Checker(ctx context.Context, check *health.CheckState) error
}

// BabbageClient is an interface with methods required for a babbage client
type BabbageClient interface {
	GetReleaseCalendar(ctx context.Context, userAccessToken, fromDay, fromMonth, fromYear string) (m release_calendar.ReleaseCalendar, err error)
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
	Babbage  BabbageClient
	ImageAPI ImageClient
	Renderer RenderClient
}
