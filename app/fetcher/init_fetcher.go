package fetcher

import (
	"context"
	"github.com/deltrinos/rss-api/app"
	"github.com/deltrinos/rss-api/app/conf"
	"github.com/deltrinos/rss-api/fetch"
	"time"
)

func init() {
	app.Fetcher = fetch.NewFetcher().
		WithStorage(app.Store).

		// add some default rss links with update duration
		AddUrl("http://feeds.bbci.co.uk/news/uk/rss.xml").
		AddUrlWithDuration("http://feeds.bbci.co.uk/news/technology/rss.xml", 1*time.Minute).
		AddUrlWithDuration("http://feeds.reuters.com/reuters/UKdomesticNews?format=xml", 3*time.Minute).
		AddUrlWithDuration("http://feeds.reuters.com/reuters/technologyNews?format=xml", 2*time.Minute).

		// add some french and chinese news rss
		//AddUrl("http://www.lefigaro.fr/rss/figaro_actualites.xml").
		//AddUrl("https://www.lemonde.fr/rss/une.xml").
		//AddUrl("https://www.scmp.com/rss/91/feed").

		// run fetcher cron
		RunEvery(conf.Env.FetcherDuration, context.Background())
}
