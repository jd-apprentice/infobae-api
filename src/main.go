package main

import (
	"InfobaeAPI/src/config"
	"InfobaeAPI/src/constants"
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
			"message": constants.WelcomeMessage,
		})
	})

	xml.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})

	infobae.GET("/", controllers.LastPost)
	infobae.GET("/:topic", controllers.PostByTopic)
	xml.GET("/sitemap", controllers.GetSitemaps)
	xml.GET("/news", controllers.GetNews)

	port := fmt.Sprintf(":%s", config.App.Port)
	r.Run(port)
}
