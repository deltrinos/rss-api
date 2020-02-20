package engine

import (
	"github.com/deltrinos/rss-api/app"
	"github.com/deltrinos/rss-api/app/conf"
	"github.com/deltrinos/rss-api/engine"
)

func init() {
	app.Engine = engine.NewEngine(conf.Env.Addr)
}
