package fetch

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"time"
)

type Feed struct {
	url        string
	duration   time.Duration
	lastUpdate time.Time
	startTimes uint
}

func NewFeedWithDuration(url string, duration time.Duration) *Feed {
	return &Feed{
		url:      url,
		duration: duration,
	}
}

func (feed *Feed) ParseURL() (*gofeed.Feed, error) {
	if time.Now().After(feed.lastUpdate.Add(feed.duration)) {
		feed.startTimes++

		parser := gofeed.NewParser()
		res, err := parser.ParseURL(feed.url)
		if err != nil {
			return nil, err
		} else {
			if res != nil {
				feed.lastUpdate = time.Now()
				return res, nil
			}
		}
	}
	return nil, fmt.Errorf("it's too early for fetch again: %s, lastUpdate: %v", feed.url, feed.lastUpdate)
}
