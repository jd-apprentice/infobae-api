package constants

type ObjectUrls map[string]string

// TODO: FIND NEW LINKS
const (
	baseUrl     = "https://www.infobae.com"
	baseSection = "/feeds/sitemap/sites/"
)

var URL = ObjectUrls{
	"autos":     baseUrl + baseSection + "autos/",
	"economia":  baseUrl + baseSection + "economia/",
	"horoscopo": baseUrl + baseSection + "horoscopo/",
	"sociedad":  baseUrl + baseSection + "sociedad/",
	"salud":     baseUrl + baseSection + "salud/",
	"ciencia":   baseUrl + baseSection + "salud/ciencia/",
	"fotos":     baseUrl + baseSection + "america/fotos/",
	"mundo":     baseUrl + baseSection + "america/mundo/",
	"politica":  baseUrl + baseSection + "politica/",
	"init":      "https://www.infobae.com/sitemap.xml",
}
