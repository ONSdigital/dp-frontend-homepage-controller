package dpcache

import (
	"context"
	"fmt"
	"github.com/ONSdigital/log.go/log"
	"sync"
	"time"
)

type DpCacher interface {
	Close()
	Get(key string) (interface{}, bool)
	Set(key string, data interface{})
	AddUpdateFunc(key string, updateFunc func() (string, error))
	StartUpdates(ctx context.Context)
}

type DpCache struct {
	cache          sync.Map
	updateInterval time.Duration
	close          chan struct{}
	updateFuncs    map[string]func() (string, error)
}

func (dc *DpCache) Get(key string) (interface{}, bool) {
	return dc.cache.Load(key)
}

func (dc *DpCache) Set(key string, data interface{}) {
	dc.cache.Store(key, data)
}

func (dc *DpCache) Close() {
	dc.close <- struct{}{}
	dc.cache = sync.Map{}
}

func NewDpCache(updateInterval time.Duration) DpCacher {
	return &DpCache{
		cache:          sync.Map{},
		updateInterval: updateInterval,
		close:          make(chan struct{}),
	}
}

func (dc *DpCache) AddUpdateFunc(key string, updateFunc func() (string, error)) {
	dc.updateFuncs[key] = updateFunc
}

func (dc *DpCache) StartUpdates(ctx context.Context) {
	ticker := time.NewTicker(dc.updateInterval)
	if len(dc.updateFuncs) == 0 {
		return
	}

	for {
		select {
		case <-ticker.C:
			for key, updateFunc := range dc.updateFuncs {
				updatedContent, err := updateFunc()
				if err != nil {
					log.Event(ctx, fmt.Sprintf("HOMEPAGE_CACHE_UPDATE_FAILED. failed to update homepage cache for %s", key), log.Error(err), log.ERROR)
					continue
				}
				dc.Set(key, updatedContent)
			}

		case <-dc.close:
			return
		case <-ctx.Done():
			return
		}

	}
}
