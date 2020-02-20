package interfaces

import (
	"github.com/deltrinos/rss-api/models"
	"github.com/mmcdole/gofeed"
	"time"
)

type IStorageQueries interface {
	ListBy(params models.ListParams) []*gofeed.Item
	Size() int
	Categories() []string
	Providers() []string
	LastUpdate() time.Time
}
