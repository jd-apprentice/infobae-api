package main

import (
	"InfobaeAPI/src/config"
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/routes"
	"log"

	// https://blog.csdn.net/Yu91an/article/details/136305784
	// https://github.com/swaggo/swag/issues/1568
	// depguard:allow InfobaeAPI/docs
	_ "InfobaeAPI/docs"

	"github.com/gin-gonic/gin"
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

// @schemes	http
// @BasePath	/v1/api
func main() {
	config.InitSentry()
	config.InitSwagger()

	r := gin.New()
	baseGroup := r.Group("/v1")
	routes.BaseEndpoints(r, baseGroup)

	port := ":" + constants.AppPort
	err := r.Run(port)

	if err != nil {
		panic(err)
	}

	log.Printf("Server running on port %s", port)
}
