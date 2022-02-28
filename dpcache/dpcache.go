package dpcache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	"github.com/ONSdigital/log.go/v2/log"
)

type DpCacher interface {
	Close()
	Get(key string) (interface{}, bool)
	Set(key string, data *model.HomepageData)
	AddUpdateFunc(key string, updateFunc func() (*model.HomepageData, error))
	StartUpdates(ctx context.Context, channel chan error)
}

type DpCache struct {
	cache          sync.Map
	updateInterval time.Duration
	close          chan struct{}
	updateFuncs    map[string]func() (*model.HomepageData, error)
}

func (dc *DpCache) Get(key string) (interface{}, bool) {
	return dc.cache.Load(key)
}

func (dc *DpCache) Set(key string, data *model.HomepageData) {
	dc.cache.Store(key, data)
}

func (dc *DpCache) Close() {
	dc.close <- struct{}{}
	for key, _ := range dc.updateFuncs {
		dc.cache.Store(key, "")
	}
	dc.updateFuncs = make(map[string]func() (*model.HomepageData, error))
}

func NewDpCache(updateInterval time.Duration) DpCacher {
	return &DpCache{
		cache:          sync.Map{},
		updateInterval: updateInterval,
		close:          make(chan struct{}),
		updateFuncs:    make(map[string]func() (*model.HomepageData, error)),
	}
}

func (dc *DpCache) AddUpdateFunc(key string, updateFunc func() (*model.HomepageData, error)) {
	dc.updateFuncs[key] = updateFunc
}

func (dc *DpCache) UpdateContent(ctx context.Context) error {
	for key, updateFunc := range dc.updateFuncs {
		updatedContent, err := updateFunc()
		if err != nil {
			return fmt.Errorf("HOMEPAGE_CACHE_UPDATE_FAILED. failed to update homepage cache for %s. error: %v", key, err)
		}
		dc.Set(key, updatedContent)
	}
	return nil
}

func (dc *DpCache) StartUpdates(ctx context.Context, errorChannel chan error) {
	ticker := time.NewTicker(dc.updateInterval)
	if len(dc.updateFuncs) == 0 {
		return
	}

	err := dc.UpdateContent(ctx)
	if err != nil {
		errorChannel <- err
		dc.Close()
		return
	}

	for {
		select {
		case <-ticker.C:
			err := dc.UpdateContent(ctx)
			if err != nil {
				log.Error(ctx, err.Error(), err)
			}

		case <-dc.close:
			return
		case <-ctx.Done():
			return
		}
	}
}
