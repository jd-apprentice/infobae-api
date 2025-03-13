package routes

import (
	"InfobaeAPI/src/controllers"

	"github.com/gin-gonic/gin"
)

func InfobaeEndpoints(group *gin.RouterGroup) {
	infobae := group.Group("/infobae")
	infobae.GET("/", controllers.LastPost)
	infobae.GET("/:topic", controllers.PostByTopic)
	infobae.GET("/details", controllers.LastDetailedPost)
	infobae.GET("/details/:topic", controllers.DetailsByTopic)
}
