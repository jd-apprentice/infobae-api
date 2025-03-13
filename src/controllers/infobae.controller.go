package controllers

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/models"
	"InfobaeAPI/src/services"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LastPost
//
//	@Summary		LastPost returns the last post from the /xml/news endpoint of Infobae.
//	@Schemes		http
//	@Description	The response will be a JSON object with the keys "message", "url", "lastmod", and "changefreq".
//	@Description	The "message" key will contain the message "This is the last Infobae post from /xml/news".
//	@Description	The "url", "lastmod", and "changefreq" keys will contain the values of the last post.
//	@Tags			Infobae
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/infobae/ [get]
func LastPost(context *gin.Context) {
	// TODO: Try to REUSE logic from GetAllNews.
	// LastPost returns the last post from the /xml/news endpoint of Infobae.
	// The response will be a JSON object with the keys "message", "url", "lastmod", and "changefreq".
	// The "message" key will contain the message "This is the last Infobae post from /xml/news".
	// The "url", "lastmod", and "changefreq" keys will contain the values of the last post.
	url := fmt.Sprintf("%s%s", constants.BaseURL, constants.RandomSection)

	newsIndex := &models.NewsIndex{
		URL: []models.URL{},
		XMLName: xml.Name{
			Local: "urlset",
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
		},
	}

	err := services.FetchAndParseXML(url, newsIndex)

	news := services.FilterNews(newsIndex.URL)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching news index: %v", err)})

		return
	}

	lastNew := news[len(news)-1]

	context.JSON(http.StatusOK, gin.H{
		"message":    constants.LastPostMessage,
		"url":        lastNew["url"],
		"lastmod":    lastNew["lastmod"],
		"changefreq": lastNew["changefreq"],
	})
}

// PostByTopic
//
//	@Summary		PostByTopic returns a JSON response with the last posts for the given topic.
//	@Schemes		http
//	@Description	The response will be a JSON object with the keys "url", "lastmod", and "changefreq".
//	@Description	The "url", "lastmod", and "changefreq" keys will contain the values of the last post.
//	@Description	The size parameter is optional and defaults to 5.
//	@Description	The topic should be one of the categories returned by the /sitemap endpoint.
//	@Description	The response will be a JSON object with a single key "news" which is an array
//	@Tags			Infobae
//	@Accept			json
//	@Produce		json
//	@Param			topic	path		string	true	"Topic"
//	@Success		200		{object}	map[string]interface{}
//	@Router			/infobae/{topic} [get]
func PostByTopic(context *gin.Context) {
	category := context.Param("topic")
	size := context.Query("size")

	if size == "" {
		size = "5"
	}

	sizeInt, err := strconv.Atoi(size)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})

		return
	}

	url := constants.InfobaeIndex

	sitemapIndex := &models.SitemapIndex{}
	newsIndex := &models.NewsIndex{}

	err = services.FetchAndParseXML(url, sitemapIndex)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching sitemap index: %v", err)})

		return
	}

	sitemaps := services.FilterSitemaps(sitemapIndex.Sitemaps)

	for _, sitemap := range sitemaps {
		sitemapURL := fmt.Sprintf("%s", sitemap["url"])
		if sitemap["category"] == category {
			err := services.FetchAndParseXML(sitemapURL, newsIndex)
			news := services.FilterNews(newsIndex.URL)

			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching news index: %v", err)})

				return
			}

			news = news[:sizeInt]

			context.JSON(http.StatusOK, gin.H{
				"news": news,
			})

			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Topic %s not found", category)})
}

func LastDetailedPost(_ *gin.Context) {
	// TODO: Create a function that scrapes the last post from the /xml/news endpoint of Infobae.
	// The response should contain the title, description, and content (images) of the post.
}
func DetailsByTopic(_ *gin.Context) {
	// TODO: Same as LastDetailedPost but for a specific topic.
}
