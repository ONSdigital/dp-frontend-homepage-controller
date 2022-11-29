package feature

import (
	"context"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/dp-topic-api/sdk"
	apiError "github.com/ONSdigital/dp-topic-api/sdk/errors"
)

func GetHomepageContentFuncMock(ctx context.Context, userAccessToken, collectionID, lang, path string) (zebedee.HomepageContent, error) {
	return zebedee.HomepageContent{
		Intro: zebedee.Intro{
			Title:    "Welcome to the Office for National Statistics",
			Markdown: "Markdown text here",
		},
		FeaturedContent: []zebedee.Featured{
			{
				Title:       "Featured content one",
				Description: "Featured content one description",
				URI:         "Featured content one URI",
				ImageID:     "123",
			},
			{
				Title:       "Featured content two",
				Description: "Featured content two description",
				URI:         "Featured content two URI",
				ImageID:     "456",
			},
			{
				Title:       "Featured content three",
				Description: "Featured content three description",
				URI:         "Featured content three URI",
				ImageID:     "",
			},
		},
		AroundONS: []zebedee.Featured{
			{
				Title:       "Around ONS one",
				Description: "Around ONS one description",
				URI:         "Around ONS one URI",
				ImageID:     "123",
			},
			{
				Title:       "Around ONS two",
				Description: "Around ONS two description",
				URI:         "Around ONS two URI",
				ImageID:     "",
			},
		},
		ServiceMessage: "",
		URI:            "",
		Type:           "",
		Description: zebedee.HomepageDescription{
			Title:           "Homepage description title",
			Summary:         "Homepage description summary",
			Keywords:        []string{"keyword one", "keyword two"},
			MetaDescription: "UKPOP",
			Unit:            "",
			PreUnit:         "",
			Source:          "UKPOP",
		},
		EmergencyBanner: zebedee.EmergencyBanner{
			Type:        "notable_death",
			Title:       "Emergency banner title",
			Description: "Emergency banner description",
			URI:         "www.google.com",
			LinkText:    "More info",
		},
	}, nil
}

func GetTimeseriesMainFigureFuncMock(ctx context.Context, userAuthToken, collectionID, lang, uri string) (zebedee.TimeseriesMainFigure, error) {
	return zebedee.TimeseriesMainFigure{
		Months: []zebedee.TimeseriesDataPoint{
			{
				Value: "677.89",
				Label: "Jan 2020",
			},
			{
				Value: "679.56",
				Label: "Feb 2020",
			},
		},
		Years: []zebedee.TimeseriesDataPoint{
			{
				Value: "907.89",
				Label: "2020",
			},
			{
				Value: "1009.56",
				Label: "2021",
			},
		},
		Quarters: []zebedee.TimeseriesDataPoint{
			{
				Value: "13.97",
				Label: "Q1",
			},
			{
				Value: "14.68",
				Label: "Q2",
			},
		},
		RelatedDocuments: []zebedee.Related{
			{
				Title: "Related thing",
				URI:   "test/uri/timeseries/123",
			},
		},
		Description: zebedee.TimeseriesDescription{
			CDID: "LOLZ",
			Unit: "%",
		},
		URI: "test/uri/timeseries/456",
	}, nil
}

func GetDownloadVariantFuncMock(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, imageID, variant string) (image.ImageDownload, error) {
	return image.ImageDownload{}, nil
}

func GetSubtopicsPublicFuncMock(ctx context.Context, reqHeaders sdk.Headers, id string) (*models.PublicSubtopics, apiError.Error) {
	// TODO Extend data setup when topic summaries work is completed, can use the
	// id to determine different responses

	return &models.PublicSubtopics{
		Count:       0,
		Offset:      0,
		Limit:       0,
		TotalCount:  0,
		PublicItems: nil,
	}, nil
}

func GetTopicPublicFuncMock(ctx context.Context, reqHeaders sdk.Headers, id string) (*models.Topic, apiError.Error) {
	// TODO Extend data setup when topic summaries work is completed, can use the
	// id to determine different responses

	releaseDate := time.Date(2022, time.November, 23, 9, 30, 0, 0, time.Local)

	return &models.Topic{
		ID:          id,
		ReleaseDate: &releaseDate,
		Title:       "",
	}, nil
}
