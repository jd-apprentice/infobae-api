package main

import (
	"InfobaeAPI/src/config"
	"InfobaeAPI/src/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.GetConfig()

	r := gin.Default()
	for _, proxy := range config.Proxies {
		r.SetTrustedProxies([]string{proxy})
	}

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

	port := fmt.Sprintf(":%s", config.App.Port)
	r.Run(port)
}
