package main

import (
	"InfobaeAPI/src/config"
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/controllers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.GetConfig()

	r := gin.Default()
	for _, proxy := range config.Proxies {
		err := r.SetTrustedProxies([]string{proxy})

		if err != nil {
			log.Printf("Error setting trusted proxies: %v", err)
		}
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

	if port == ":" {
		port = ":3000"
	}

	err := r.Run(port)

	if err != nil {
		log.Fatal(err)
	}
}
