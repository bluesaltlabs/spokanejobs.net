// Egnyte
package egnyte

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

func ScrapeJobs() []models.ScrapedJob {
	jobs := []models.ScrapedJob{} // todo: initialize this when number of jobs is known
	scrapeUrl := "https://jobs.jobvite.com/egnyte/jobs"


	// Create the new collector
	c := colly.NewCollector(
		colly.AllowedDomains("linkding.bluesaltlabs.com", "bluesaltlabs.com", "jobs.jobvite.com"),
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
