package endpoints

import (
	"github.com/deltrinos/rss-api/app"
	"github.com/deltrinos/rss-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func init() {
	app.Engine.Handler.POST("/addFeed", func(c *gin.Context) {
		var params models.AddFeedParams
		err := c.ShouldBindJSON(&params)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		_, err = url.ParseRequestURI(params.Url)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			app.Fetcher.AddUrl(params.Url)
		}

		c.JSON(http.StatusOK, gin.H{
			"url": params.Url,
		})
	})
}
