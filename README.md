# rss-api

This project aggregate RSS Feeds and provide an API for fetch news articles.

## Fetching RSS Feeds

It will run fetcher every 1 minute by default (duration can be override by environment variable FETCHER_DURATION).

It starts a http server on :3000 by default (environment variable ADDR for modify).

PRODUCTION environment variable is available for switch off debug logs.

try it quickly:
```
go run main.go
``` 

## API

### GET /stats

show all availables categories and providers. It provides number of items and last update in the Store.


### GET /list

list news articles in the order by published time.
Query parameters can be added :

| Parameter     | Mandatory     |                          |
| ------------- |:-------------:| :----------------------- |
| itemsByPage   | Yes           |   number of news         |
| page          | No            |   page number            |
| category      | No            |    filtering by category |
| provider      | No            |    filtering by provider |

The mandatory parameter is for demonstrate ability to checking parameters.

Example :
```
http://localhost:3000/list?itemsByPage=5

http://localhost:3000/list?itemsByPage=10&page=2&category=domesticNews&provider=Reuters:%20UK
```

all list requests are cached until Store lastUpdate is modified.

### POST /addFeed

it allows add a RSS feed

payload
```
{"url": "http://feeds.bbci.co.uk/news/video_and_audio/business/rss.xml"}
```

use curl to add a new feed

```
curl --request POST --data '{"url":"http://feeds.bbci.co.uk/news/video_and_audio/business/rss.xml"}'  http://localhost:3000/addFeed
```
