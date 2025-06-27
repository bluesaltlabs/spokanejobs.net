package companies

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var egnyteScrapeUrl string = "https://jobs.jobvite.com/egnyte/jobs"
var egnyteJobUrlPrefix string = egnyteScrapeUrl + ""

// EgnyteScraper implements a scraper for Egnyte
type EgnyteScraper struct {
	*types.BaseScraper
}

func NewEgnyteScraper() *EgnyteScraper {
	return &EgnyteScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "egnyte",
			ScheduleHour: 4,
			Config: types.ScraperConfig{
				BaseURL: "https://www.egnyte.com/careers",
			},
		},
	}
}

func (e *EgnyteScraper) GetName() string {
	return e.Name
}

func (e *EgnyteScraper) ScrapedJobs() []types.ScrapedJob {
	if len(e.Jobs) == 0 {
		e.Jobs = e.scrapeJobs()
	}
	return e.Jobs
}

func (e *EgnyteScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (e *EgnyteScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := e.getCollector()

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, e.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", e.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), e.Name)
		}
	})

	// Process Job Line - TODO: Add proper selectors
	c.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	c.Visit(egnyteScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
	return jobs
}

func (e *EgnyteScraper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("jobs.jobvite.com"),
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

func (e *EgnyteScraper) GetScheduleHour() int {
	return e.ScheduleHour
}

func (e *EgnyteScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(e.Jobs, e.Name, outputDir)
}
