package endpoints

import (
	"github.com/deltrinos/rss-api/app"
	"github.com/deltrinos/rss-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	app.Engine.Handler.GET("/list", func(c *gin.Context) {
		var params models.ListParams
		err := c.ShouldBindQuery(&params)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		res := app.Queries.ListBy(params)
		if res != nil {
			c.JSON(http.StatusOK, res)
		} else {
			c.AbortWithStatus(http.StatusNoContent)
		}
	})
}
