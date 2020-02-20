package store

import (
	"github.com/deltrinos/rss-api/models"
	"github.com/mmcdole/gofeed"
	"sync"
	"time"
)

type CacheResult struct {
	Items      []*gofeed.Item
	lastUpdate time.Time
}

type Cache struct {
	results *sync.Map
}

// Provide caching in the API to allow for faster response times
// based on sync.Map
func (c *Cache) Get(params models.ListParams, storeLastUpdate time.Time) ([]*gofeed.Item, bool) {
	res, exists := c.results.Load(params)
	if exists {
		cache, cacheExists := res.(*CacheResult)
		if cacheExists {
			if cache.lastUpdate == storeLastUpdate {
				return cache.Items, true
			}
		}
	}
	return nil, false
}

func (c *Cache) Save(params models.ListParams, res []*gofeed.Item, storeLastUpdate time.Time) []*gofeed.Item {
	c.results.Store(params, &CacheResult{
		Items:      res,
		lastUpdate: storeLastUpdate,
	})
	return res
}
