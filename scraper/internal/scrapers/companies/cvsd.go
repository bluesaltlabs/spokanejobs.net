package companies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
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

func (c *CVSDScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := c.getCollector()

	collector.OnScraped(func(r *colly.Response) {
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

func (c *CVSDScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
