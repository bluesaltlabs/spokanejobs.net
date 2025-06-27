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

type ScldScraper struct {
	*types.BaseScraper
}

func NewScldScraper() *ScldScraper {
	return &ScldScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "scld",
			ScheduleHour: 19,
			Config: types.ScraperConfig{
				BaseURL: "https://www.scld.org/careers/",
			},
		},
	}
}

func (scraper *ScldScraper) ScrapeJobs() []types.ScrapedJob {
	var jobs []types.ScrapedJob
	c := utils.NewCollector(scraper.Config)

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, scraper.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", scraper.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), scraper.Name)
		}

		// Also print to stdout for backward compatibility
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	c.OnHTML(".job-posting", func(h *colly.HTMLElement) {
		title := strings.TrimSpace(h.DOM.Find(".job-title").Text())
		location := strings.TrimSpace(h.DOM.Find(".job-location").Text())
		link, _ := h.DOM.Find("a").Attr("href")

		if title != "" && link != "" {
			// Make link absolute if it's relative
			if !strings.HasPrefix(link, "http") {
				link = scraper.Config.BaseURL + strings.TrimPrefix(link, "/")
			}

			job := types.ScrapedJob{
				Title: title,
				Url:   link,
				City:  location,
			}
			jobs = append(jobs, job)
		}
	})

	c.Visit(scraper.Config.BaseURL)
	return jobs
}

func (scraper *ScldScraper) GetName() string {
	return scraper.Name
}

func (scraper *ScldScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.ScrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *ScldScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (scraper *ScldScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(scraper.Jobs, scraper.Name, outputDir)
}
