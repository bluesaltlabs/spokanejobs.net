package companies

import (
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

type SpokaneComputerScraper struct {
	*types.BaseScraper
}

func NewSpokaneComputerScraper() *SpokaneComputerScraper {
	return &SpokaneComputerScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "spokane_computer",
			ScheduleHour: 21,
			Config: types.ScraperConfig{
				BaseURL: "https://www.spokanecomputer.com/careers",
			},
		},
	}
}

func (scraper *SpokaneComputerScraper) ScrapeJobs() []types.ScrapedJob {
	var jobs []types.ScrapedJob
	c := utils.NewCollector(scraper.Config)

	c.OnHTML(".career-opportunity", func(h *colly.HTMLElement) {
		title := strings.TrimSpace(h.DOM.Find(".position-title").Text())
		location := strings.TrimSpace(h.DOM.Find(".position-location").Text())
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

func (scraper *SpokaneComputerScraper) GetName() string {
	return scraper.Name
}

func (scraper *SpokaneComputerScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.ScrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *SpokaneComputerScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}
