package cache

import (
	"context"
	"time"

	dpcache "github.com/ONSdigital/dp-cache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	"github.com/ONSdigital/log.go/v2/log"
)

// HomepageCache is a wrapper to dpcache.Cache which has additional fields and methods specifically for caching homepage data
type HomepageCache struct {
	*dpcache.Cache
}

// NewHomepageCache create a homepage cache object to be used in the service which will update at every updateInterval
// If updateInterval is nil, this means that the cache will only be updated once at the start of the service
func NewHomepageCache(ctx context.Context, updateInterval *time.Duration) (*HomepageCache, error) {
	config := dpcache.Config{
		UpdateInterval: updateInterval,
	}

	cache, err := dpcache.NewCache(ctx, config)
	if err != nil || cache == nil {
		logData := log.Data{
			"config": config,
		}
		log.Error(ctx, "failed to create cache from dpcache", err, logData)
		return nil, err
	}

	return &HomepageCache{cache}, nil
}

// AddUpdateFunc adds an update function to the homepage cache
func (hc *HomepageCache) AddUpdateFunc(key string, updateFunc func() (*model.HomepageData, error)) {
	hc.UpdateFuncs[key] = func() (interface{}, error) {
		return updateFunc()
	}
}
