package services

import (
	"InfobaeAPI/src/constants"
	"InfobaeAPI/src/models"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/20478577/why-does-json-unmarshal-work-with-reference-but-not-pointer
func FetchAndParseXML(website string, model interface{}) error {

	parsedURL, err := url.Parse(website)
	if err != nil {
		return err
	}

	if parsedURL.Scheme != constants.ValidSchema || parsedURL.Host != constants.ValidBaseURL {
		return errors.New(constants.UnauthorizedURL)
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		var ErrFetchingXML = errors.New("error fetching XML")

		return fmt.Errorf(constants.ErrorFormat, ErrFetchingXML, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		var ErrReadingResponseBody = errors.New("error reading response body")

		return fmt.Errorf(constants.ErrorFormat, ErrReadingResponseBody, err)
	}

	switch model := model.(type) {
	case *models.SitemapIndex:
		err = xml.Unmarshal(body, model)
	case *models.NewsIndex:
		err = xml.Unmarshal(body, model)
	default:
		var ErrInvalidModelType = errors.New("invalid model type")

		return fmt.Errorf(constants.ErrorFormatModels, ErrInvalidModelType, model)
	}

	if err != nil {
		var ErrParsingXML = errors.New("error parsing XML")

		return fmt.Errorf(constants.ErrorFormat, ErrParsingXML, err)
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
