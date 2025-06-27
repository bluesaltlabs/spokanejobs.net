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

var enhancedSoftwareProductsScrapeUrl string = "https://www.espsolution.net/careers"
var enhancedSoftwareProductsJobUrlPrefix string = enhancedSoftwareProductsScrapeUrl + ""

// EnhancedSoftwareProductsScraper implements a scraper for Enhanced Software Products
type EnhancedSoftwareProductsScraper struct {
	*types.BaseScraper
}

func NewEnhancedSoftwareProductsScraper() *EnhancedSoftwareProductsScraper {
	return &EnhancedSoftwareProductsScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "enhanced_software_products",
			ScheduleHour: 5,
			Config: types.ScraperConfig{
				BaseURL: "https://www.enhancedsoftware.com/careers",
			},
		},
	}
}

func (e *EnhancedSoftwareProductsScraper) GetName() string {
	return e.Name
}

func (e *EnhancedSoftwareProductsScraper) ScrapedJobs() []types.ScrapedJob {
	if len(e.Jobs) == 0 {
		e.Jobs = e.scrapeJobs()
	}
	return e.Jobs
}

func (e *EnhancedSoftwareProductsScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (e *EnhancedSoftwareProductsScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(e.Jobs, e.Name, outputDir)
}

func (e *EnhancedSoftwareProductsScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := e.getCollector()

	collector.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, e.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", e.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), e.Name)
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

	collector.Visit(enhancedSoftwareProductsScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", collector)
	return jobs
}

func (e *EnhancedSoftwareProductsScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.espsolution.net", "espsolution.net"),
		colly.CacheDir("./scraper_cache"),
	)

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", utils.DefaultUserAgent)
		log.Printf("visiting: %s\n", r.URL.String())
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return collector
}

func (e *EnhancedSoftwareProductsScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
