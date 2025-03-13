package config

import (
	"InfobaeAPI/docs"
	"InfobaeAPI/src/constants"
	"os"
)

func InitSwagger() {
	docs.SwaggerInfo.Host = "localhost:" + constants.NodePort

	if os.Getenv("GIN_MODE") == "release" {
		docs.SwaggerInfo.Host = constants.ProdURL
	}
}
