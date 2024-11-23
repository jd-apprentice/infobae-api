package models

import (
	"encoding/xml"
)

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

type NewsIndex struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []URL    `xml:"url"`
}

type Location struct {
	Loc string `xml:"loc"`
}

type URL struct {
	Location
	LastMod    string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
}

type Sitemap struct {
	Location
}
