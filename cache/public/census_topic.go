package public

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"github.com/ONSdigital/dp-topic-api/models"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
	"github.com/ONSdigital/log.go/v2/log"
)

// UpdateCensusTopic is a function to update the census topic cache in web (public) mode.
// This function talks to the dp-topic-api via its public endpoints to retrieve the census topic and its subtopic ids
// The data returned by the dp-topic-api is of type *models.PublicSubtopics which is then transformed to *cache.Topic in this function for the controller
// If an error has occurred, this is captured in log.Error and then an empty census topic is returned
func UpdateCensusTopic(ctx context.Context, cfg *config.Config, topicClient topicCli.Clienter) func() *cache.Topic {
	return func() *cache.Topic {
		// get root topics from dp-topic-api
		censusTopic, err := topicClient.GetTopicPublic(ctx, topicCli.Headers{}, cfg.CensusTopicID)
		censusSubtopics, err := topicClient.GetSubtopicsPublic(ctx, topicCli.Headers{}, cfg.CensusTopicID)
		if err != nil {
			logData := log.Data{
				"req_headers": topicCli.Headers{},
			}
			log.Error(ctx, "failed to get root topics from topic-api", err, logData)
			return cache.GetEmptyCensusTopic()
		}

		if censusSubtopics.PublicItems == nil {
			err := errors.New("root topic public items is nil")
			log.Error(ctx, "failed to deference root topics items pointer", err)
			return cache.GetEmptyCensusTopic()
		}

		var censusTopicCache *cache.Topic

		// census topic data for caching which includes subtopic ids
		subtopicsIDChan := make(chan models.Topic)
		censusTopicCache = getRootTopicCachePublic(ctx, subtopicsIDChan, topicClient, *censusTopic)

		if censusTopicCache == nil {
			err := errors.New("census root topic not found")
			log.Error(ctx, "failed to get census topic to cache", err)
			return cache.GetEmptyCensusTopic()
		}

		return censusTopicCache
	}
}

func getRootTopicCachePublic(ctx context.Context, subtopicsChan chan models.Topic, topicClient topicCli.Clienter, rootTopic models.Topic) *cache.Topic {
	rootTopicCache := &cache.Topic{
		ID:              rootTopic.ID,
		LocaliseKeyName: rootTopic.Title,
	}

	subtopicsIDMap := cache.NewSubTopicsMap()
	subtopicsIDMap.AppendSubtopicItems(rootTopic)

	var wg sync.WaitGroup
	wg.Add(2)

	// get subtopics ids
	go func() {
		defer wg.Done()
		getSubtopicsIDsPublic(ctx, subtopicsChan, topicClient, rootTopic.ID)
		close(subtopicsChan)
	}()

	// extract subtopic id from channel to update rootTopicCache
	go func() {
		defer wg.Done()
		for subtopic := range subtopicsChan {
			subtopicsIDMap.AppendSubtopicItems(subtopic)
		}
	}()

	wg.Wait()

	rootTopicCache.List = subtopicsIDMap
	rootTopicCache.Query = subtopicsIDMap.GetSubtopicsIDsQuery()

	return rootTopicCache
}

func getSubtopicsIDsPublic(ctx context.Context, subtopicsIDChan chan models.Topic, topicClient topicCli.Clienter, topLevelTopicID string) {
	// get subtopics from dp-topic-api
	subTopics, err := topicClient.GetSubtopicsPublic(ctx, topicCli.Headers{}, topLevelTopicID)
	if err != nil {
		if err.Status() != http.StatusNotFound {
			logData := log.Data{
				"req_headers":        topicCli.Headers{},
				"top_level_topic_id": topLevelTopicID,
			}
			log.Error(ctx, "failed to get subtopics from topic-api", err, logData)
		}

		// stop as there are no subtopics items or failed to get subtopics
		return
	}

	// deference sub topics items to allow ranging through them
	var subTopicItems []models.Topic
	if subTopics.PublicItems != nil {
		subTopicItems = *subTopics.PublicItems
	} else {
		err := errors.New("sub topics public items is nil")
		log.Error(ctx, "failed to deference sub topics items pointer", err)
		return
	}

	var wg sync.WaitGroup

	// get subtopics ids of the subtopics items if they exist
	for i := range subTopicItems {
		wg.Add(1)

		// send subtopic id to channel
		subtopicsIDChan <- subTopicItems[i]

		go func(index int) {
			defer wg.Done()
			getSubtopicsIDsPublic(ctx, subtopicsIDChan, topicClient, subTopicItems[index].ID)
		}(i)
	}
	wg.Wait()
}