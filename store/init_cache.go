package store

import "sync"

var queriesCache *Cache

func init() {
	queriesCache = &Cache{
		results: &sync.Map{},
	}
}
