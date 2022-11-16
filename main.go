package main

import (
	"net/http"

	"url-shortener/items"

	"github.com/gin-gonic/gin"
)

var urls = map[string]string{}

func main() {
	r := gin.Default()
	r.GET("/add", func(c *gin.Context) {
		url := c.Query("url")

		if len(url) == 0 || !items.IsUrl(url) {
			c.JSON(http.StatusBadRequest, items.ErrorResponse("Invalid URL"))
			return
		}
		urlIndex := items.GetURLIndex()
		if _, ok := urls[urlIndex]; !ok {
			urls[urlIndex] = url
		}
		c.JSON(http.StatusOK, utils.SuccessResonse(map[string]interface{}{
			"index": urlIndex,
		}))
	})

	r.GET("/:index", func(c *gin.Context) {
		urlIndex := c.Param("index")
		if _, ok := urls[urlIndex]; !ok {
			c.JSON(http.StatusNotFound, items.ErrorResponse("URL not found"))
			return
		}
		c.JSON(http.StatusOK, items.SuccessResonse(map[string]interface{}{
			"url": urls[urlIndex],
		}))
	})

	r.Run()
}
