// Spokane Public Schools
package sps

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

func ScrapeJobs() []*models.ScrapedJob {
	jobs := []models.ScrapedJob{} // todo: initialize this when number of jobs is known
	scrapeUrl := "https://www.schooljobs.com/careers/spokaneschools"


	// Create the new collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.schooljobs.com", "schooljobs.com"),
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
