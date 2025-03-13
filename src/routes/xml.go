package routes

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/controllers"

	"github.com/gin-gonic/gin"
)

func XMLEndpoints(group *gin.RouterGroup) {
	xml := group.Group("/xml")
	xml.GET("/", func(c *gin.Context) {
		c.JSON(constants.SucessStatusCode, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})
	xml.GET("/sitemap", controllers.GetSitemaps)
	xml.GET("/news", controllers.GetNews)
}
