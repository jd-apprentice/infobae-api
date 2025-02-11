package controllers

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/models"
	"InfobaeAPI/src/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Try to REUSE logic from GetAllNews.
// LastPost returns the last post from the /xml/news endpoint of Infobae.
// The response will be a JSON object with the keys "message", "url", "lastmod", and "changefreq".
// The "message" key will contain the message "This is the last Infobae post from /xml/news".
// The "url", "lastmod", and "changefreq" keys will contain the values of the last post.
func LastPost(c *gin.Context) {
	url := fmt.Sprintf("%s%s", constants.BaseUrl, constants.RandomSection)

	newsIndex := &models.NewsIndex{}

	err := services.FetchAndParseXML(url, newsIndex)

	news := services.FilterNews(newsIndex.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching news index: %v", err)})
		return
	}

	lastNew := news[len(news)-1]

	c.JSON(http.StatusOK, gin.H{
		"message":    constants.LastPostMessage,
		"url":        lastNew["url"],
		"lastmod":    lastNew["lastmod"],
		"changefreq": lastNew["changefreq"],
	})
}

// TODO: Investigate if this is efficient.
// PostByTopic returns a JSON response with the last posts for the given topic.
// The size parameter is optional and defaults to 5.
// The topic should be one of the categories returned by the /sitemap endpoint.
// The response will be a JSON object with a single key "news" which is an array
// of objects with the keys "url", "lastmod", and "changefreq".
func PostByTopic(c *gin.Context) {
	category := c.Param("topic")
	size := c.Query("size")

	if size == "" {
		size = "5"
	}

	sizeInt, err := strconv.Atoi(size)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
		return
	}

	url := constants.InfobaeIndex

	sitemapIndex := &models.SitemapIndex{}
	newsIndex := &models.NewsIndex{}

	err = services.FetchAndParseXML(url, sitemapIndex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching sitemap index: %v", err)})
		return
	}

	sitemaps := services.FilterSitemaps(sitemapIndex.Sitemaps)

	for _, sitemap := range sitemaps {

		sitemapURL := fmt.Sprintf("%s", sitemap["url"])

		if sitemap["category"] == category {
			err := services.FetchAndParseXML(sitemapURL, newsIndex)
			news := services.FilterNews(newsIndex.URL)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching news index: %v", err)})
				return
			}

			news = news[:sizeInt]

			c.JSON(http.StatusOK, gin.H{
				"news": news,
			})

			return
		}
	}
}
