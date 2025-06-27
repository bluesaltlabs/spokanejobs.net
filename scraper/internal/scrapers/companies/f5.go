package companies

import (
	"encoding/json"
	"log"
	"os"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var f5ScrapeUrl string = "https://www.f5.com/company/careers/locations/spokane"
var f5JobUrlPrefix string = f5ScrapeUrl + ""

// F5Scraper implements a scraper for F5
type F5Scraper struct {
	*types.BaseScraper
}

func NewF5Scraper() *F5Scraper {
	return &F5Scraper{
		BaseScraper: &types.BaseScraper{
			Name:         "f5",
			ScheduleHour: 6,
			Config: types.ScraperConfig{
				BaseURL: "https://f5.com/careers",
			},
		},
	}
}

func (f *F5Scraper) GetName() string {
	return f.Name
}

func (f *F5Scraper) ScrapedJobs() []types.ScrapedJob {
	if len(f.Jobs) == 0 {
		f.Jobs = f.scrapeJobs()
	}
	return f.Jobs
}

func (f *F5Scraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (f *F5Scraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(f.Jobs, f.Name, outputDir)
}

func (f *F5Scraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := f.getCollector()

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, f.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", f.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), f.Name)
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

	c.Visit(f5ScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
	return jobs
}

func (f *F5Scraper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.espsolution.net", "espsolution.net"),
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
