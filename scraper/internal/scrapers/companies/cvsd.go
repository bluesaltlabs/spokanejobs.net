package companies

import (
	"encoding/json"
	"log"
	"os"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var cvsdScrapeUrl string = "https://jobs.redroverk12.com/org/cvsd"
var cvsdJobUrlPrefix string = cvsdScrapeUrl + ""

// CVSDScraper implements a scraper for Central Valley School District
type CVSDScraper struct {
	*types.BaseScraper
}

func NewCVSDScraper() *CVSDScraper {
	return &CVSDScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "cvsd",
			ScheduleHour: 3,
			Config: types.ScraperConfig{
				BaseURL: "https://www.cvsd.org/careers",
			},
		},
	}
}

func (c *CVSDScraper) GetName() string {
	return c.Name
}

func (c *CVSDScraper) ScrapedJobs() []types.ScrapedJob {
	if len(c.Jobs) == 0 {
		c.Jobs = c.scrapeJobs()
	}
	return c.Jobs
}

func (c *CVSDScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (c *CVSDScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(c.Jobs, c.Name, outputDir)
}

func (c *CVSDScraper) scrapeJobs() []types.ScrapedJob {
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

	// Process Job Line - TODO: Add proper selectors
	collector.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	collector.Visit(cvsdScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", collector)
	return jobs
}

func (c *CVSDScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("jobs.redroverk12.com"),
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
