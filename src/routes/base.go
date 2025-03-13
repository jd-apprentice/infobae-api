package routes

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/controllers"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BaseEndpoints(router *gin.Engine, group *gin.RouterGroup) {
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: true,
		Timeout:         constants.DefaultTimeout,
	}))
	router.GET("/", controllers.Help)

	err := router.SetTrustedProxies([]string{constants.TrustedProxies})

	if err != nil {
		panic(err)
	}

	// BASE
	base := group.Group("/api")
	base.GET("/", controllers.Welcome)
	base.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// ENDPOINTS
	InfobaeEndpoints(base)
	XMLEndpoints(base)
}
