package companies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var cityOfSpokaneScrapeUrlPrefix string = "https://www.governmentjobs.com/careers/spokanecity"
var cityOfSpokaneScrapeUrl string = cityOfSpokaneScrapeUrlPrefix + "?sort=PostingDate%7CDescendingz"
var cityOfSpokaneJobUrlPrefix string = cityOfSpokaneScrapeUrlPrefix + "/jobs/"

// CityOfSpokaneScraper implements a scraper for City of Spokane
type CityOfSpokaneScraper struct {
	*types.BaseScraper
}

func NewCityOfSpokaneScraper() *CityOfSpokaneScraper {
	return &CityOfSpokaneScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "city_of_spokane",
			ScheduleHour: 1,
			Config: types.ScraperConfig{
				BaseURL: "https://www.spokanecity.org/careers",
			},
		},
	}
}

func (c *CityOfSpokaneScraper) GetName() string {
	return c.Name
}

func (c *CityOfSpokaneScraper) ScrapedJobs() []types.ScrapedJob {
	if len(c.Jobs) == 0 {
		c.Jobs = c.scrapeJobs()
	}
	return c.Jobs
}

func (c *CityOfSpokaneScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (c *CityOfSpokaneScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(c.Jobs, c.Name, outputDir)
}

func (c *CityOfSpokaneScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := c.getCollector()

	collector.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, c.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", c.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), c.Name)
		}

		// Also print to stdout for backward compatibility
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	collector.OnHTML("li.list-item a", func(h *colly.HTMLElement) {
		log.Printf("\n-----\n\ncity_of_spokane-pass: %+v\n-----\n", h) // debug

		j := &types.ScrapedJob{}
		url := h.Attr("href")
		selection := h.DOM

		job_id := strings.TrimPrefix(url, cityOfSpokaneJobUrlPrefix)
		title := selection.Text()

		j.JobId = job_id
		j.Url = url
		j.Title = title

		jobs = append(jobs, *j)
	})

	collector.Visit(cityOfSpokaneScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", collector)
	return jobs
}

func (c *CityOfSpokaneScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.governmentjobs.com", "governmentjobs.com"),
		colly.CacheDir("./scraper_cache"),
	)

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return collector
}

func (c *CityOfSpokaneScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
