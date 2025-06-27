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

var numericaScrapeUrl string = "https://recruiting.paylocity.com/recruiting/jobs/All/3e0b08f9-66f1-4617-9a26-159571b53dc2/Numerica-Credit-Union"

// NumericaScraper implements a scraper for Numerica
type NumericaScraper struct {
	*types.BaseScraper
}

func NewNumericaScraper() *NumericaScraper {
	return &NumericaScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "numerica",
			ScheduleHour: 14,
			Config: types.ScraperConfig{
				BaseURL: "https://numericacu.com/careers",
			},
		},
	}
}

func (n *NumericaScraper) GetName() string {
	return n.Name
}

func (n *NumericaScraper) ScrapedJobs() []types.ScrapedJob {
	if len(n.Jobs) == 0 {
		n.Jobs = n.scrapeJobs()
	}
	return n.Jobs
}

func (n *NumericaScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (n *NumericaScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(n.Jobs, n.Name, outputDir)
}

func (n *NumericaScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := n.getCollector()

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, n.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", n.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), n.Name)
		}

		// Also print to stdout for backward compatibility
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	// Process Job Line - TODO: Add proper selectors
	c.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	log.Printf("\n-----\n\nColly instance created: %+v\n\n", c)
	c.Visit(numericaScrapeUrl)
	return jobs
}

func (n *NumericaScraper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("recruiting.paylocity.com"),
		colly.CacheDir("./scraper_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return c
}

func (n *NumericaScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
