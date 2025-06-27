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

var gravityJackScrapeUrl string = "https://careers.smartrecruiters.com/GravityJack"

// GravityJackScraper implements a scraper for Gravity Jack
type GravityJackScraper struct {
	*types.BaseScraper
}

func NewGravityJackScraper() *GravityJackScraper {
	return &GravityJackScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "gravity_jack",
			ScheduleHour: 8,
			Config: types.ScraperConfig{
				BaseURL: "https://gravityjack.com/careers",
			},
		},
	}
}

func (g *GravityJackScraper) GetName() string {
	return g.Name
}

func (g *GravityJackScraper) ScrapedJobs() []types.ScrapedJob {
	if len(g.Jobs) == 0 {
		g.Jobs = g.scrapeJobs()
	}
	return g.Jobs
}

func (g *GravityJackScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (g *GravityJackScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(g.Jobs, g.Name, outputDir)
}

func (g *GravityJackScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := g.getCollector()

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, g.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", g.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), g.Name)
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
	c.Visit(gravityJackScrapeUrl)
	return jobs
}

func (g *GravityJackScraper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("careers.smartrecruiters.com"),
		colly.CacheDir("./scraper_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", utils.DefaultUserAgent)
		log.Printf("visiting: %s\n", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return c
}

func (g *GravityJackScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
