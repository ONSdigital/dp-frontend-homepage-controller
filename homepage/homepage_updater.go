package homepage

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-homepage-controller/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
	"github.com/ONSdigital/log.go/v2/log"
)

type Updater struct {
	clients *Clients
}

func (hu *Updater) GetHomePageUpdateFor(ctx context.Context, userAccessToken, collectionID, lang string) func() (*model.HomepageData, error) {
	return func() (*model.HomepageData, error) {
		mappedMainFigures := make(map[string]*model.MainFigure)
		var wg sync.WaitGroup
		responses := make(chan *model.MainFigure, len(mainFigureMap))
		// nolint: gocritic
		for id, figure := range mainFigureMap {
			wg.Add(1)
			go func(ctx context.Context, zcli ZebedeeClient, id string, figure MainFigure) {
				defer wg.Done()
				zebResponses := []zebedee.TimeseriesMainFigure{}
				for _, uri := range figure.uris {
					zebResponse, err := zcli.GetTimeseriesMainFigure(ctx, userAccessToken, collectionID, lang, uri)
					if err != nil {
						log.Error(ctx, "error getting timeseries data", err, log.Data{"timeseries-data": uri})
						mappedErrorFigure := &model.MainFigure{ID: id}
						responses <- mappedErrorFigure
						return
					}
					zebResponses = append(zebResponses, zebResponse)
				}
				trendInfo := getTrendInfo(ctx, userAccessToken, collectionID, lang, zcli, figure)
				latestMainFigure := getLatestTimeSeriesData(ctx, zebResponses)
				mappedMainFigure := mapper.MainFigure(ctx, id, figure.datePeriod, figure.differenceInterval, latestMainFigure, trendInfo)
				responses <- mappedMainFigure
				return
			}(ctx, hu.clients.Zebedee, id, figure)
		}
		wg.Wait()
		close(responses)

		for response := range responses {
			log.Info(ctx, "the response of this request was", log.Data{"response": response})
			mappedMainFigures[response.ID] = response
		}

		// Get homepage data from Zebedee
		homepageContent, err := hu.clients.Zebedee.GetHomepageContent(ctx, userAccessToken, collectionID, lang, HomepagePath)
		if err != nil {
			log.Error(ctx, "error getting homepage data from client", err, log.Data{"content-path": HomepagePath})
		}

		var mappedFeaturedContent []model.Feature
		if len(homepageContent.FeaturedContent) > 0 {
			imageObjects := map[string]image.ImageDownload{}
			for _, fc := range homepageContent.FeaturedContent {
				if fc.ImageID != "" {
					imageDetails, err := hu.clients.ImageAPI.GetDownloadVariant(ctx, userAccessToken, "", "", fc.ImageID, ImageVariant)
					if err != nil {
						log.Error(ctx, "error getting image download variant", err, log.Data{"featured-content-entry": fc.Title})
					}
					imageObjects[fc.ImageID] = imageDetails
				}
			}
			mappedFeaturedContent = mapper.FeaturedContent(homepageContent, imageObjects)
		}

		var mappedAroundONS []model.Feature
		if len(homepageContent.AroundONS) > 0 {
			imageObjects := map[string]image.ImageDownload{}
			for _, fc := range homepageContent.AroundONS {
				if fc.ImageID != "" {
					imageVariant, err := hu.clients.ImageAPI.GetDownloadVariant(ctx, userAccessToken, "", "", fc.ImageID, ImageVariant)
					if err != nil {
						log.Error(ctx, "error getting image download variant", err, log.Data{"around-ons-entry": fc.Title})
					}
					imageObjects[fc.ImageID] = imageVariant
				}
			}
			mappedAroundONS = mapper.AroundONS(homepageContent, imageObjects)
		}

		homepageData := &model.HomepageData{
			AroundONS:       &mappedAroundONS,
			EmergencyBanner: homepageContent.EmergencyBanner,
			FeaturedContent: &mappedFeaturedContent,
			MainFigures:     mappedMainFigures,
			ServiceMessage:  homepageContent.ServiceMessage,
		}

		return homepageData, nil
	}
}

func (hu *Updater) UpdateNavigationData(ctx context.Context, lang string) func() *topicModel.Navigation {
	return func() *topicModel.Navigation {
		headers := topicCli.Headers{}
		options := topicCli.Options{}

		switch lang {
		case "cy":
			options.Lang = topicCli.Welsh
		default:
			options.Lang = topicCli.English
		}

		navigationData, err := hu.clients.Topic.GetNavigationPublic(ctx, headers, options)
		if err != nil {
			logData := log.Data{
				"headers": headers,
				"options": options,
			}
			log.Error(ctx, "failed to get navigation data from client", err, logData)
		}

		return navigationData
	}
}
