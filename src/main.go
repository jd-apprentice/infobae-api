package main

import (
	"InfobaeAPI/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	xml := r.Group("/xml")
	infobae := api.Group("/infobae")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Infobae API",
		})
	})

	xml.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})

	infobae.GET("/", controllers.RandomInfobaePost)
	infobae.GET("/:topic", controllers.InfobaePostByTopic)
	xml.GET("/sitemap", controllers.SitemapIndex)

	r.Run(":8080")
}
