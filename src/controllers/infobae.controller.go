package controllers

import "github.com/gin-gonic/gin"

// TODO: Transform to JSON and return a random post
func RandomInfobaePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Random Infobae Post",
	})
}

// TODO: Transform to JSON and return posts by topic
func InfobaePostByTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Infobae Post By Topic",
	})
}
