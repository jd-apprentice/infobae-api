package models

import (
	"encoding/xml"

	"gorm.io/gorm"
)

type New struct {
	gorm.Model
	URL      string
	Category string
}

type NewDetail struct {
	gorm.Model
	URL        string
	LastMod    string
	ChangeFreq string
	Title      string
}

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	Loc string `xml:"loc"`
}
