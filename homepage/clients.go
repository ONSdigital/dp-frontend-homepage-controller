package homepage

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
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

// RenderClient is an interface with methods required for rendering a template
type RenderClient interface {
	Do(string, []byte) ([]byte, error)
	Checker(ctx context.Context, check *health.CheckState) error
}

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Renderer RenderClient
	Zebedee  ZebedeeClient
	Babbage  BabbageClient
	ImageAPI ImageClient
}
