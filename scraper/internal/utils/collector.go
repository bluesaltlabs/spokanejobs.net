package utils

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"github.com/gocolly/colly"
)

// NewCollector creates a standardized collector with common configuration
func NewCollector(config types.ScraperConfig) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(config.AllowedDomains...),
		colly.CacheDir("./scraper_cache"),
	)

	// Standard headers and error handling
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return c
}
