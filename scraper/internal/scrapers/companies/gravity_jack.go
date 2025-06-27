package companies

import (
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
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

func (g *GravityJackScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := g.getCollector()

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
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
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
