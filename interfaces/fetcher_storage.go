package interfaces

import "github.com/mmcdole/gofeed"

type IFetcherStorage interface {
	Save(feed *gofeed.Feed)
}
