package store

import (
	"github.com/deltrinos/rss-api/app"
	"github.com/deltrinos/rss-api/store"
)

func init() {
	// store will contents all feeds
	app.Store = store.NewStore().Run()

	// make store serve queries
	app.SetQueries(app.Store)
}
