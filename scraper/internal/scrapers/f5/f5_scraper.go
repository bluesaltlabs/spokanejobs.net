// F5
package f5

import (
  "os"
  "encoding/json"
  "log"
  "strings"
  "github.com/gocolly/colly"
  "gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

// Set the base URL to scrape
var scrapeUrl string		= "https://www.f5.com/company/careers/locations/spokane"
var JobUrlPrefix string = scrapeUrl + ""

func ScrapeJobs() []*models.ScrapedJob {

  // Create the Jobs collector
  jobs := make([]*models.ScrapedJob, 0)

  // Create a new collector
  c := getCollector()

    // Display all jobs after scraping completes as json to the standard output
  c.OnScraped(func(r *colly.Response) {
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    enc.Encode(jobs)
  })

  // Process Job Line
  c.OnHTML("", func(h *colly.HTMLElement) {
    // todo
  })

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)

  log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)

    return jobs
}

func getCollector() colly.Collector {
  c := colly.NewCollector(
    colly.AllowedDomains("www.espsolution.net", "espsolution.net"),
    colly.CacheDir("./scraper_cache"),
  )

          // Set headers and log the link visited on each request
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Accept-Language", "en-US")
    r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
    log.Printf("visiting: %s\n", r.URL.String())
  })

  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    log.Printf("Error while scraping: %s\n", e.Error())
  })

  return *c
}

func trimSpaces(s string) string {
  return strings.TrimSpace(s)
}
