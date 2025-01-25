package tests

import (
	"InfobaeAPI/src/constants"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Infobae API",
		})
	})

	req, _ := http.NewRequest("GET", "/api/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	message := constants.WelcomeMessage
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestInfobaeEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	infobae := router.Group("/api/infobae")
	infobae.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the last Infobae post from /xml/news",
		})
	})

	req, _ := http.NewRequest("GET", "/api/infobae/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	message := constants.LastPostMessage
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestXMLEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	xml := router.Group("/xml")
	xml.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})

	req, _ := http.NewRequest("GET", "/xml/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	message := "Available index are at /xml/sitemap"
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
