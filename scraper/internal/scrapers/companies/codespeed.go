package companies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"github.com/gocolly/colly"
)

var codespeedScrapeUrl string = "https://www.codespeed.com/careers"
var codespeedJobUrlPrefix string = codespeedScrapeUrl + ""

// CodespeedScraper implements a scraper for Codespeed
type CodespeedScraper struct {
	*types.BaseScraper
}

func NewCodespeedScraper() *CodespeedScraper {
	return &CodespeedScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "codespeed",
			ScheduleHour: 2,
			Config: types.ScraperConfig{
				BaseURL: "https://codespeed.com/careers",
			},
		},
	}
}

func (c *CodespeedScraper) GetName() string {
	return c.Name
}

func (c *CodespeedScraper) ScrapedJobs() []types.ScrapedJob {
	if len(c.Jobs) == 0 {
		c.Jobs = c.scrapeJobs()
	}
	return c.Jobs
}

func (c *CodespeedScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (c *CodespeedScraper) scrapeJobs() []types.ScrapedJob {
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

	collector.Visit(codespeedScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", collector)
	return jobs
}

func (c *CodespeedScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.codespeed.com", "codespeed.com"),
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

func (c *CodespeedScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
