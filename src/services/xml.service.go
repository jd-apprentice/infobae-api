package services

import (
	"InfobaeAPI/src/models"
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchAndParseSitemapIndex(url string) (*models.SitemapIndex, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sitemapIndex models.SitemapIndex
	err = xml.Unmarshal(body, &sitemapIndex)
	if err != nil {
		return nil, err
	}

	return &sitemapIndex, nil
}

func FilterSitemaps(sitemaps []models.Sitemap) []gin.H {
	filtered := []gin.H{}
	for _, sitemap := range sitemaps {
		categoryURL := strings.Split(sitemap.Loc, "/")
		categoryName := categoryURL[len(categoryURL)-2]

		if strings.Contains(sitemap.Loc, "/sitemap2/") || categoryName == "news-sitemap2" {
			continue
		}

		filtered = append(filtered, gin.H{
			"url":      sitemap.Loc,
			"category": categoryName,
		})
	}
	return filtered
}
