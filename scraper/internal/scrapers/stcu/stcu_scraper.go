// Spokane Teacher's Credit Union
package stcu

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

func ScrapeJobs() []*types.ScrapedJob {
	// Create the Jobs collector
	jobs := make([]*types.ScrapedJob, 0)

	scrapeUrl := "https://careers.smartrecruiters.com/STCU1"


	// Create the new collector
	c := colly.NewCollector(
		colly.AllowedDomains("careers.smartrecruiters.com"),
		colly.CacheDir("./scraper_cache"),
	)


	// Set the language and User Agent headers,
	// and log each page visit.
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Accept-Language", "en-US")
    r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
    log.Printf("visiting: %s\n", r.URL.String())
  })

  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    log.Printf("Error while scraping: %s\n", e.Error())
  })

  // Process Job Line
  c.OnHTML("", func(h *colly.HTMLElement) {
  	// todo
  })


  // log the beginning of the script.
  log.Printf("\n-----\n\nColly instance created: %+v\n\n", c)

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)


	return jobs
}

func trimSpaces(s string) string {
  return strings.TrimSpace(s)
}
