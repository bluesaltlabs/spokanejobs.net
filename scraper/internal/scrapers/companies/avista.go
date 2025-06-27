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

var avistaScrapeUrlPrefix string = "https://recruiting2.ultipro.com/AVI1000AMAST/JobBoard/362abf68-95c3-4b17-a39d-76a6efe5ff18"
var avistaScrapeUrl string = avistaScrapeUrlPrefix + "?o=postedDateDesc&f4=ITpN1FEJvlWudZe0AayqWA"
var avistaJobUrlPrefix string = avistaScrapeUrlPrefix + "/OpportunityDetail?opportunityId="

// AvistaScraper implements a scraper for Avista
type AvistaScraper struct {
	*types.BaseScraper
}

func NewAvistaScraper() *AvistaScraper {
	return &AvistaScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "avista",
			ScheduleHour: 0,
			Config: types.ScraperConfig{
				BaseURL: "https://www.myavista.com/about-us/working-at-avista/careers",
			},
		},
	}
}

func (a *AvistaScraper) GetName() string {
	return a.Name
}

func (a *AvistaScraper) ScrapedJobs() []types.ScrapedJob {
	if len(a.Jobs) == 0 {
		a.Jobs = a.scrapeJobs()
	}
	return a.Jobs
}

func (a *AvistaScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (a *AvistaScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(a.Jobs, a.Name, outputDir)
}

func (a *AvistaScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := a.getCollector()

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, a.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", a.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), a.Name)
		}

		// Also print to stdout for backward compatibility
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	c.OnHTML("div.opportunity", func(h *colly.HTMLElement) {
		j := &types.ScrapedJob{}
		selection := h.DOM

		url, _ := selection.Find("h3 a.opportunity-link").Attr("href")
		title := selection.Find("h3 a.opportunity-link").Text()
		job_id := strings.TrimPrefix(url, avistaJobUrlPrefix)

		j.JobId = job_id
		j.Url = url
		j.Title = title

		jobs = append(jobs, *j)
	})

	c.Visit(avistaScrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
	return jobs
}

func (a *AvistaScraper) getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("recruiting2.ultipro.com"),
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
