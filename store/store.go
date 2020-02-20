package store

import (
	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type Store struct {
	// using channels instead of mutex
	feedsCh chan *gofeed.Feed
	itemsCh chan *gofeed.Item

	runOnce *sync.Once

	items      []*gofeed.Item
	itemsLink  map[string]*gofeed.Item
	lastUpdate time.Time
}

func NewStore() *Store {
	return &Store{
		feedsCh:   make(chan *gofeed.Feed),
		itemsCh:   make(chan *gofeed.Item),
		runOnce:   &sync.Once{},
		items:     []*gofeed.Item{},
		itemsLink: map[string]*gofeed.Item{},
	}
}

func (s *Store) Save(feed *gofeed.Feed) {
	if feed != nil {
		log.Debug().Msgf("save in store feed: %v", feed.Title)
		s.feedsCh <- feed
	}
}

func (s *Store) runFeedsConsumer() {
	for {
		select {
		case f := <-s.feedsCh:
			if f != nil && f.Items != nil {
				for _, item := range f.Items {
					if item.Categories == nil {
						item.Categories = []string{}
					}
					if item.Custom == nil {
						item.Custom = map[string]string{}
					}

					if len(item.Categories) == 0 {
						item.Categories = append(item.Categories, "Unlisted")
					}
					item.Custom[ItemCustomProviderKey] = f.Title

					s.itemsCh <- item
				}
			}
		}
	}
}

func (s *Store) runItemsConsumer() {
	for {
		select {
		case i := <-s.itemsCh:
			_, exists := s.itemsLink[i.Link]
			if !exists {
				s.itemsLink[i.Link] = i
				s.items = append(s.items, i)
				s.lastUpdate = time.Now()
			}
		}
	}
}

func (s *Store) Run() *Store {
	s.runOnce.Do(func() {
		// start Consumers
		go s.runFeedsConsumer()
		go s.runItemsConsumer()
	})
	return s
}
