package app

import (
	"github.com/deltrinos/rss-api/engine"
	"github.com/deltrinos/rss-api/fetch"
	"github.com/deltrinos/rss-api/interfaces"
	"github.com/deltrinos/rss-api/store"
)

var (
	Engine  *engine.Engine
	Store   *store.Store
	Fetcher *fetch.Fetcher

	Queries interfaces.IStorageQueries
)

// this interface is used by endpoints to make queries from store
func SetQueries(q interfaces.IStorageQueries) {
	Queries = q
}
