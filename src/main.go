package main

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/controllers"
	"fmt"

	// https://blog.csdn.net/Yu91an/article/details/136305784
	// https://github.com/swaggo/swag/issues/1568
	_ "InfobaeAPI/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger embed files

//	@title			InfobaeAPI
//	@version		1.0
//	@description	InfobaeAPI is a REST API that provides access to the latest news from Infobae.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome stranger",
			"problem":  "There is nothing here",
			"solution": "Check docs at /api/docs/index.html",
		})
	})

	err := r.SetTrustedProxies([]string{constants.TrustedProxies})

	if err != nil {
		panic(err)
	}

	// BASE
	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": constants.WelcomeMessage,
		})
	})

	//SWAG
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// INFOBAE
	infobae := api.Group("/infobae")
	infobae.GET("/", controllers.LastPost)
	infobae.GET("/:topic", controllers.PostByTopic)

	// XML
	xml := r.Group("/xml")
	xml.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})
	xml.GET("/sitemap", controllers.GetSitemaps)
	xml.GET("/news", controllers.GetNews)

	port := fmt.Sprintf(":%s", constants.AppPort)

	err = r.Run(port)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on port", port)
}
