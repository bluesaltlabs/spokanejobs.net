// OpenEye
package openeye

import (
	"os"
	"encoding/json"
	"log"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

// Set the base URL to scrape
var ScrapeUrl string = "https://job-boards.greenhouse.io/openeye"
var JobUrlPrefix string = ScrapeUrl + "/jobs/"

func ScrapeJobs() []*models.ScrapedJob {

	// Create the Jobs collector
	// todo: this would be part of the parent class
	jobs := make([]*models.ScrapedJob, 0)

  // todo: move this to a generic class.
	// Create a new collector
	c := getCollector()

	// Display all jobs after scraping completes as json to the standard output
  c.OnScraped(func(r *colly.Response) {
      enc := json.NewEncoder(os.Stdout)
      enc.SetIndent("", "  ")
      enc.Encode(jobs)
  })

  log.Printf("\n-----\n\nColly instance created: %+v\n\n", c)

  // Process Job Line
  c.OnHTML("div.job-posts--table tr.job-post a", func(h *colly.HTMLElement) {
    // Create the scraped_job struct instance
    j := &models.ScrapedJob{}
    url := h.Attr("href")
    selection := h.DOM

    // Retrieve attributes available
    job_id := strings.TrimPrefix(url, JobUrlPrefix)
    title := trimSpaces(selection.Find("p.body.body--medium").Text())

    // Set the attribute values
    j.JobId = job_id
    j.Url = url
    j.Title = title

    // Visit the specified job to retrieve further details
    getJobDetails(j)

    // Append it to the parent array
    jobs = append(jobs, j)
  })

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(ScrapeUrl)

  log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)

  return jobs
}

func getJobDetails(j *models.ScrapedJob) {
	// Create a new collector
	c := getCollector()

	c.OnHTML("", func(h *colly.HTMLElement) {

		j.Description = "" // todo
	})

	// Visit the job detail URL
	c.Visit(j.Url)
}

func getCollector() colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("job-boards.greenhouse.io"),
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

// todo: move this to a helper function
func trimSpaces(s string) string {
  return strings.TrimSpace(s)
}
