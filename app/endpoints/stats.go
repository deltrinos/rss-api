package endpoints

import (
	"github.com/deltrinos/rss-api/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	app.Engine.Handler.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"size":       app.Queries.Size(),
			"categories": app.Queries.Categories(),
			"providers":  app.Queries.Providers(),
			"lastUpdate": app.Queries.LastUpdate(),
		})
	})
}
