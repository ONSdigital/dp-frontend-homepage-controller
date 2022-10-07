package cache

import (
	"context"
	"errors"
	"fmt"
	dpcache "github.com/ONSdigital/dp-cache"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
	"github.com/ONSdigital/log.go/v2/log"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CensusTopicID = "topicData"
)

type TopicCache struct {
	*dpcache.Cache
}

func UpdateTopic(ctx context.Context, topicClient topicCli.Clienter) func() *topicModel.Topic {
	return func() *topicModel.Topic {
		// add logic to get topic from dp-topic-api
		go func() {
			resp, err := http.Get("http://localhost:25300/topics/1")
			defer resp.Body.Close()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println("data: ", string(body))
		}()
		// then return topic
		return nil
	}
}

// NewTopicCache create a topic cache object to be used in the service which will update at every updateInterval
// If updateInterval is nil, this means that the cache will only be updated once at the start of the service
func NewTopicCache(ctx context.Context, updateInterval *time.Duration) (*TopicCache, error) {
	config := dpcache.Config{
		UpdateInterval: updateInterval,
	}

	cache, err := dpcache.NewCache(ctx, config)
	if err != nil {
		logData := log.Data{
			"update_interval": updateInterval,
		}
		log.Error(ctx, "failed to create cache from dpcache", err, logData)
		return nil, err
	}

	return &TopicCache{cache}, nil
}

func (dc *TopicCache) GetData(ctx context.Context, key string) (*topicModel.Topic, error) {
	topicCacheInterface, ok := dc.Get(key)
	if !ok {
		err := fmt.Errorf("cached topic data with key %s not found", key)
		log.Error(ctx, "failed to get cached topic data", err)
		return &topicModel.Topic{}, err
	}

	topicCacheData, ok := topicCacheInterface.(*topicModel.Topic)
	if !ok {
		err := errors.New("topicCacheInterface is not type *Topic")
		log.Error(ctx, "failed type assertion on topicCacheInterface", err)
		return &topicModel.Topic{}, err
	}

	if topicCacheData == nil {
		err := errors.New("topicCacheData is nil")
		log.Error(ctx, "cached topic data is nil", err)
		return &topicModel.Topic{}, err
	}

	return topicCacheData, nil
}

// AddUpdateFunc adds an update function to the topic cache for a topic with the `key` passed to the function
// This update function will then be triggered once or at every fixed interval as per the prior setup of the TopicCache
func (dc *TopicCache) AddUpdateFunc(key string, updateFunc func() *topicModel.Topic) {
	dc.UpdateFuncs[key] = func() (interface{}, error) {
		// error handling is done within the updateFunc
		return updateFunc(), nil
	}
}
