package controllers

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/models"
	"InfobaeAPI/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SitemapIndex handles the /xml/sitemap-index route by fetching and parsing the infobae sitemap index
// filtering out sitemap2 and news-sitemap2 entries
// and returning the filtered list of sitemaps in the response body.
func SitemapIndex(c *gin.Context) {
	url := constants.InfobaeIndex

	sitemapIndex := &models.SitemapIndex{}

	err := services.FetchAndParseXML(url, sitemapIndex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching sitemap index: %v", err)})
		return
	}

	sitemaps := services.FilterSitemaps(sitemapIndex.Sitemaps)

	c.JSON(http.StatusOK, gin.H{
		"sitemaps": sitemaps,
	})
}
