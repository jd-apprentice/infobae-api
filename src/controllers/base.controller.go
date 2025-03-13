package controllers

import (
	"InfobaeAPI/src/constants"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(constants.SucessStatusCode, gin.H{
		"message": constants.WelcomeMessage,
	})
}

func Help(c *gin.Context) {
	c.JSON(constants.SucessStatusCode, gin.H{
		"message":  "Welcome stranger",
		"problem":  "There is nothing here",
		"solution": "Check docs at /api/docs/index.html",
	})
}
