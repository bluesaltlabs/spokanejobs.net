package openeye

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"github.com/gocolly/colly"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

// type OpeneyeScraper struct{}

// func (o OpenEyeScraper) GetName() string {
// 	return "OpenEye"
//}

func ScrapeJobs() []models.ScrapedJob {

	// Set the base URL to scrape
	//scrapeUrl := "https://job-boards.greenhouse.io/openeye"
	scrapeUrl := "https://linkding.bluesaltlabs.com/assets/284"
	//jobSscrapeUrl := "https://job-boards.greenhouse.io/openeye/jobs/7997742002"
	//jobSscrapeUrl := "https://linkding.bluesaltlabs.com/assets/285"

	// Create the Job IDs collector
  //jobIDs := []string{}

	// Create a new collector
	//c := colly.NewCollector(colly.AllowedDomains("job-boards.greenhouse.io"))
	c := colly.NewCollector(
		colly.AllowedDomains("linkding.bluesaltlabs.com", "bluesaltlabs.com", "job-boards.greenhouse.io"),
		colly.CacheDir("./scraper_cache"),
	)


  // A simple check to prove the library was imported correctly
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Accept-Language", "en-US")
    r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
    fmt.Printf("Colly is visiting: %s\n", r.URL)
  })


  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    fmt.Printf("Error while scraping: %s\n", e.Error())
  })


  // Process Job Line
  c.OnHTML("div.job-posts--table tr.job-post a", func(h *colly.HTMLElement) {

  	// todo: consider making a map of jobs here.
    // Create the scraped_job struct instance
    job := models.ScrapedJob{}
    selection := h.DOM

    // job detail url prefix (todo: extract JobId from the url)
    // https://job-boards.greenhouse.io/openeye/jobs/<JobId>

    // Retrieve attributes available
    url, _ := selection.Attr("href")
    title := trimSpaces(selection.Find("p.body.body--medium").Text())
    location := trimSpaces(selection.Find("p.body.body__secondary.body--metadata").Text())
    job.Url = url
    job.Title = title
    job.Description = location

    // create the json encoder
	  enc := json.NewEncoder(os.Stdout)
	  enc.SetIndent("", " ")
		enc.Encode(job)

		// print the job attributes.
    //fmt.Printf("%s (%s) | %s\n", job.Title, location, url)
  })



  fmt.Printf("\n-----\n\nColly instance created: %+v\n\n", c)

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)


  fmt.Printf("\n-----\n\nColly instance done: %+v\n\n", c)

  return []models.ScrapedJob{}
}

// todo: Go through and write out specific steps this scraper needs to take.




func trimSpaces(s string) string {
  return strings.TrimSpace(s)
}
