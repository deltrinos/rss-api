package store

import (
	"github.com/ahmetb/go-linq"
	"github.com/deltrinos/rss-api/models"
	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"
	"time"
)

const ItemCustomProviderKey = "provider"

func (s *Store) ListBy(params models.ListParams) []*gofeed.Item {
	// serve from cache
	qRes, qExists := queriesCache.Get(params, s.lastUpdate)
	if qExists {
		log.Debug().Msgf("response from cache")
		return qRes
	}

	// or fetch from store and store to cache
	cnt := int(params.ItemsByPage)
	page := 0
	if params.Page > 0 {
		page = int(params.Page - 1)
	}

	// use qo-linq it's quite easy for handle sorting, filtering and paginating : for c# lovers :)
	var res []*gofeed.Item

	// Allow the selection of different sources of news by category and provider
	// Present news articles in the order in which they are published
	query := linq.From(s.items).
		OrderByDescending(func(i interface{}) interface{} {
			return i.(*gofeed.Item).PublishedParsed.Unix()
		}).Query

	// filtering by category
	if params.Category != "" {
		query = query.WhereT(func(i *gofeed.Item) bool {
			if i.Categories != nil {
				for _, c := range i.Categories {
					if c == params.Category {
						return true
					}
				}
			}
			return false
		})
	}

	// filtering by provider
	if params.Provider != "" {
		query = query.WhereT(func(i *gofeed.Item) bool {
			if i.Custom != nil {
				v, vExists := i.Custom[ItemCustomProviderKey]
				if vExists {
					if v == params.Provider {
						return true
					}
				}
			}
			return false
		})
	}

	// paginating
	query.
		Skip(cnt * page).
		Take(cnt).
		ToSlice(&res)

	// save to cache and returns
	return queriesCache.Save(params, res, s.lastUpdate)
}

func (s *Store) Size() int {
	return len(s.items)
}

func (s *Store) Categories() []string {
	var cats []string

	linq.From(s.items).SelectMany(
		func(item interface{}) linq.Query {
			return linq.From(item.(*gofeed.Item).Categories)
		}).
		Distinct().
		ToSlice(&cats)
	return cats
}

func (s *Store) Providers() []string {
	var providers []string

	linq.From(s.items).
		Select(func(i interface{}) interface{} {
			provider := i.(*gofeed.Item).Custom[ItemCustomProviderKey]
			return provider
		}).
		Where(func(i interface{}) bool {
			return i != nil
		}).
		Distinct().
		ToSlice(&providers)
	return providers
}

func (s *Store) LastUpdate() time.Time {
	return s.lastUpdate
}
