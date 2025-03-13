package tests_test

import (
	"InfobaeAPI/src/constants"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoint(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Infobae API",
		})
	})

	req, _ := http.NewRequest(http.MethodGet, "/api/", nil)
	www := httptest.NewRecorder()

	router.ServeHTTP(www, req)
	assert.Equal(t, 200, www.Code)
	message := constants.WelcomeMessage
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, www.Body.String())
}

func TestInfobaeEndpoint(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	infobae := router.Group("/api/infobae")
	infobae.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the last Infobae post from /xml/news",
		})
	})

	req, _ := http.NewRequest(http.MethodGet, "/api/infobae/", nil)
	www := httptest.NewRecorder()

	router.ServeHTTP(www, req)
	assert.Equal(t, 200, www.Code)
	message := constants.LastPostMessage
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, www.Body.String())
}

func TestXMLEndpoint(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	xml := router.Group("/xml")
	xml.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Available index are at /xml/sitemap",
		})
	})

	req, _ := http.NewRequest(http.MethodGet, "/xml/", nil)
	www := httptest.NewRecorder()

	router.ServeHTTP(www, req)
	assert.Equal(t, 200, www.Code)
	message := "Available index are at /xml/sitemap"
	expectedBody := `{"message":"` + message + `"}`
	assert.JSONEq(t, expectedBody, www.Body.String())
}
