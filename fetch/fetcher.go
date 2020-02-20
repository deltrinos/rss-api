package fetch

import (
	"context"
	"github.com/deltrinos/rss-api/interfaces"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type Fetcher struct {
	list    []*Feed
	runOnce *sync.Once
	store   interfaces.IFetcherStorage
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		list:    []*Feed{},
		runOnce: &sync.Once{},
	}
}

func (f *Fetcher) AddUrl(url string) *Fetcher {
	return f.AddUrlWithDuration(url, 1*time.Hour)
}

func (f *Fetcher) AddUrlWithDuration(url string, duration time.Duration) *Fetcher {
	log.Debug().Msgf("fetch %s every %s", url, duration)
	f.list = append(f.list, NewFeedWithDuration(url, duration))
	return f
}

func (f *Fetcher) WithStorage(storage interfaces.IFetcherStorage) *Fetcher {
	f.store = storage
	return f
}

func (f *Fetcher) FetchList() {
	for _, feed := range f.list {
		go func(feed *Feed) {
			res, err := feed.ParseURL()
			if err != nil {
				log.Debug().Msgf("failed to parse url: %v", err)
			} else {
				if f.store != nil {
					f.store.Save(res)
				}
			}
		}(feed)
	}
}

func (f *Fetcher) RunEvery(periodicDuration time.Duration, ctx context.Context) *Fetcher {
	f.runOnce.Do(func() {
		log.Info().Msgf("running periodic fetcher every %v", periodicDuration)
		// run it first time
		go f.FetchList()

		// start cron into a goroutine
		go func() {
			for {
				select {
				case <-time.After(periodicDuration):
					log.Debug().Msgf("fetch periodic cron started...")
					go f.FetchList()
				case <-ctx.Done():
					log.Info().Msgf("fetcher ctx done.")
					break
				}
			}
			log.Info().Msgf("fetcher cron stopped")
		}()
	})
	return f
}
