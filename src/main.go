package main

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	err := r.SetTrustedProxies([]string{constants.TrustedProxies})

	if err != nil {
		panic(err)
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

	port := fmt.Sprintf(":%s", constants.AppPort)

	err = r.Run(port)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on port", port)
}
