package companies

import (
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var huntwoodScrapeUrl string = "https://www.huntwood.com/employment-opportunities/"

// HuntwoodScraper implements a scraper for Huntwood Cabinets
type HuntwoodScraper struct {
	*types.BaseScraper
}

func NewHuntwoodScraper() *HuntwoodScraper {
	return &HuntwoodScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "huntwood",
			ScheduleHour: 9,
			Config: types.ScraperConfig{
				BaseURL: "https://www.huntwood.com/careers",
			},
		},
	}
}

func (h *HuntwoodScraper) GetName() string {
	return h.Name
}

func (h *HuntwoodScraper) ScrapedJobs() []types.ScrapedJob {
	if len(h.Jobs) == 0 {
		h.Jobs = h.scrapeJobs()
	}
	return h.Jobs
}

func (h *HuntwoodScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (h *HuntwoodScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := h.getCollector()

	collector.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, h.Name); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", h.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), h.Name)
		}
		// Removed: debug print to stdout
	})

	// Process Job Line - TODO: Add proper selectors
	collector.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	log.Printf("\n-----\n\nColly instance created: %+v\n\n", collector)
	collector.Visit(huntwoodScrapeUrl)
	return jobs
}

func (h *HuntwoodScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.huntwood.com", "huntwood.com"),
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

func (h *HuntwoodScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}

func (h *HuntwoodScraper) GetScheduleHour() int {
	return h.ScheduleHour
}
