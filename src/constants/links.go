package constants

type ObjectUrls map[string]string

const (
	BaseURL       = "https://www.infobae.com"
	BaseFeed      = "/arc/outboundfeeds"
	BaseSection   = BaseFeed + "/sitemap-index-static/"
	RandomSection = BaseFeed + "/news-sitemap2/"
	InfobaeIndex  = BaseURL + BaseSection
)
