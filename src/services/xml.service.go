package services

import (
	"InfobaeAPI/src/models"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/20478577/why-does-json-unmarshal-work-with-reference-but-not-pointer
func FetchAndParseXML(url string, model interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch model := model.(type) {
	case *models.SitemapIndex:
		err = xml.Unmarshal(body, model)
	case *models.NewsIndex:
		err = xml.Unmarshal(body, model)
	default:
		return fmt.Errorf("invalid model type: %T", model)
	}

	if err != nil {
		return err
	}

	return nil
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

func FilterNews(news []models.URL) []gin.H {
	filtered := []gin.H{}
	for _, new := range news {
		filtered = append(filtered, gin.H{
			"url":        new.Loc,
			"lastmod":    new.LastMod,
			"changefreq": new.ChangeFreq,
		})
	}
	return filtered
}
