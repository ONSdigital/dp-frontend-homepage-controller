package homepage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	model "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/log.go/log"
	"sync"
	"time"
)

type HomepageUpdater struct {
	clients *Clients
}

func (hu *HomepageUpdater) GetHomePageUpdateFor(ctx context.Context, userAccessToken, collectionID, lang string) func() (string, error) {
	return func() (string, error) {
		mappedMainFigures := make(map[string]*model.MainFigure)
		var wg sync.WaitGroup
		responses := make(chan *model.MainFigure, len(mainFigureMap))
		for id, figure := range mainFigureMap {
			wg.Add(1)
			go func(ctx context.Context, zcli ZebedeeClient, id string, figure MainFigure) {
				defer wg.Done()
				zebResponses := []zebedee.TimeseriesMainFigure{}
				for _, uri := range figure.uris {
					zebResponse, err := zcli.GetTimeseriesMainFigure(ctx, userAccessToken, collectionID, lang, uri)
					if err != nil {
						log.Event(ctx, "error getting timeseries data", log.ERROR, log.Error(err), log.Data{"timeseries-data": uri})
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
			log.Event(ctx, "the response of this request was", log.ERROR, log.Data{"response": response})
			mappedMainFigures[response.ID] = response
		}

		weekAgoTime := time.Now().AddDate(0, 0, -7)
		dateFromDay := weekAgoTime.Format("02")
		dateFromMonth := weekAgoTime.Format("01")
		dateFromYear := weekAgoTime.Format("2006")
		releaseCalResp, err := hu.clients.Babbage.GetReleaseCalendar(ctx, userAccessToken, dateFromDay, dateFromMonth, dateFromYear)
		if err != nil {
			log.Event(ctx, "error failed to get release calendar data from babbage ", log.ERROR, log.Error(err))
		}
		releaseCalModelData := mapper.ReleaseCalendar(releaseCalResp)

		// Get homepage data from Zebedee
		homepageContent, err := hu.clients.Zebedee.GetHomepageContent(ctx, userAccessToken, collectionID, lang, HomepagePath)
		if err != nil {
			log.Event(ctx, "error getting homepage data from client", log.ERROR, log.Error(err), log.Data{"content-path": HomepagePath})
		}

		var mappedFeaturedContent []model.Feature
		if len(homepageContent.FeaturedContent) > 0 {
			imageObjects := map[string]image.ImageDownload{}
			for _, fc := range homepageContent.FeaturedContent {
				if fc.ImageID != "" {
					image, err := hu.clients.ImageAPI.GetDownloadVariant(ctx, userAccessToken, "", "", fc.ImageID, ImageVariant)
					if err != nil {
						log.Event(ctx, "error getting image download variant", log.ERROR, log.Error(err), log.Data{"featured-content-entry": fc.Title})
					}
					imageObjects[fc.ImageID] = image
				}
			}
			mappedFeaturedContent = mapper.FeaturedContent(homepageContent, imageObjects)
		}

		m := mapper.Homepage(lang, mappedMainFigures, releaseCalModelData, &mappedFeaturedContent, homepageContent.ServiceMessage)

		b, err := json.Marshal(m)
		if err != nil {
			errMessage := "error marshalling body data to json"
			return "", fmt.Errorf("%s. error: %#v", errMessage, err)
		}

		templateHTML, err := hu.clients.Renderer.Do("homepage", b)
		if err != nil {
			errMessage := "error rendering page"
			return "", fmt.Errorf("%s. error: %#v", errMessage, err)
		}

		return string(templateHTML), nil
	}
}

