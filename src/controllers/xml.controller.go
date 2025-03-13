package controllers

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/models"
	"InfobaeAPI/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSitemaps fetches and parses the sitemap index at
// https://www.infobae.com/arc/outboundfeeds/sitemap-index-static/ and
// filters out the sitemap for /sitemap2/ and news-sitemap2. The rest of
// the sitemaps are returned as a JSON array in the response body.
func GetSitemaps(context *gin.Context) {
	url := constants.InfobaeIndex

	sitemapIndex := &models.SitemapIndex{}

	err := services.FetchAndParseXML(url, sitemapIndex)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching sitemap index: %v", err)})

		return
	}

	sitemaps := services.FilterSitemaps(sitemapIndex.Sitemaps)

	context.JSON(http.StatusOK, gin.H{
		"sitemaps": sitemaps,
	})
}

// GetNews fetches and parses the news index at a random section of the
// wwwsite and filters out the news. The news are returned as a JSON array
// in the response body.
func GetNews(context *gin.Context) {
	url := fmt.Sprintf("%s%s", constants.BaseURL, constants.RandomSection)

	newsIndex := &models.NewsIndex{}

	err := services.FetchAndParseXML(url, newsIndex)

	news := services.FilterNews(newsIndex.URL)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching news index: %v", err)})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"news": news,
	})
}
