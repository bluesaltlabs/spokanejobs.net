// Avista
package avista

import (
	"os"
	"encoding/json"
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// Set the base URL to scrape
var scrapeUrlPrefix string 	= "https://recruiting2.ultipro.com/AVI1000AMAST/JobBoard/362abf68-95c3-4b17-a39d-76a6efe5ff18"
var scrapeUrl string 				= scrapeUrlPrefix + "?o=postedDateDesc&f4=ITpN1FEJvlWudZe0AayqWA"
var JobUrlPrefix string 		= scrapeUrlPrefix + "/OpportunityDetail?opportunityId="

func ScrapeJobs() []*types.ScrapedJob {

	// Create the Jobs collector
	jobs := make([]*types.ScrapedJob, 0)

	// Create a new collector
	c := getCollector()

	// Display all jobs after scraping completes as json to the standard output
  c.OnScraped(func(r *colly.Response) {
      enc := json.NewEncoder(os.Stdout)
      enc.SetIndent("", "  ")
      enc.Encode(jobs)
  })

  // Process Job Line
  c.OnHTML("div.opportunity", func(h *colly.HTMLElement) {
 	  j := &types.ScrapedJob{}
    selection := h.DOM

    // Retrieve attributes available
    url, _ := selection.Find("h3 a.opportunity-link").Attr("href")
    title  := selection.Find("h3 a.opportunity-link").Text()
    job_id := strings.TrimPrefix(url, JobUrlPrefix)

    // Set the attribute values
    j.JobId = job_id
    j.Url 	= url
    j.Title = title

    // Append it to the parent array
    jobs = append(jobs, j)
  })

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)

  log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)

	return jobs
}

func getCollector() colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("recruiting2.ultipro.com"),
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
