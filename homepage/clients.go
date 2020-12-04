package homepage

import (
	"context"

	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
)

//go:generate moq -out mocks_test.go -pkg homepage . ZebedeeClient RenderClient BabbageClient ImageClient

// ZebedeeClient is an interface with methods required for a zebedee client
type ZebedeeClient interface {
	GetTimeseriesMainFigure(ctx context.Context, userAuthToken, collectionID, lang, uri string) (m zebedee.TimeseriesMainFigure, err error)
	GetHomepageContent(ctx context.Context, userAccessToken, collectionID, lang, path string) (m zebedee.HomepageContent, err error)
}

// BabbageClient is an interface with methods required for a babbage client
type BabbageClient interface {
	GetReleaseCalendar(ctx context.Context, userAccessToken, fromDay, fromMonth, fromYear string) (m release_calendar.ReleaseCalendar, err error)
}

// ImageClient is an interface with methods required for the Image API service client
type ImageClient interface {
	GetDownloadVariant(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, imageID, variant string) (m image.ImageDownload, err error)
}

// RenderClient is an interface with methods required for rendering a template
type RenderClient interface {
	Do(string, []byte) ([]byte, error)
}
