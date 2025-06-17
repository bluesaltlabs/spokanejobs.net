// City Of Spokane
package city_of_spokane

import (
	"os"
	"encoding/json"
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

// Set the base URL to scrape
var scrapeUrlPrefix string 	= "https://www.governmentjobs.com/careers/spokanecity"
var scrapeUrl string 				= scrapeUrlPrefix + "?sort=PostingDate%7CDescendingz"
var JobUrlPrefix string 		= scrapeUrlPrefix + "/jobs/"

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
  c.OnHTML("li.list-item a", func(h *colly.HTMLElement) {
  log.Printf("\n-----\n\ncity_of_spokane-pass: %+v\n-----\n", h)  // debug

 		// Create the scraped_job struct instance
    j := &models.ScrapedJob{}
    url := h.Attr("href")
    selection := h.DOM

    // Retrieve attributes available
    job_id := strings.TrimPrefix(url, JobUrlPrefix)
    title := selection.Text()

    // Set the attribute values
    j.JobId = job_id
    j.Url = url
    j.Title = title

    // Visit the specified job to retrieve further details
    //getJobDetails(j)

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
		colly.AllowedDomains("www.governmentjobs.com", "governmentjobs.com"),
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
