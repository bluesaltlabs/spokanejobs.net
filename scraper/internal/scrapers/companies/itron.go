package companies

import (
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

type ItronScraper struct {
	*types.BaseScraper
}

func NewItronScraper() *ItronScraper {
	return &ItronScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "itron",
			ScheduleHour: 11,
			Config: types.ScraperConfig{
				BaseURL: "https://careers.itron.com/",
			},
		},
	}
}

func (scraper *ItronScraper) GetName() string {
	return scraper.Name
}

func (scraper *ItronScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.ScrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *ItronScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (scraper *ItronScraper) ScrapeJobs() []types.ScrapedJob {
	var jobs []types.ScrapedJob
	c := utils.NewCollector(scraper.Config)

	c.OnHTML(`div[data-automation-id="job-card"]`, func(h *colly.HTMLElement) {
		title := strings.TrimSpace(h.DOM.Find(`h3[data-automation-id="job-title"]`).Text())
		location := strings.TrimSpace(h.DOM.Find(`span[data-automation-id="job-location"]`).Text())
		link, _ := h.DOM.Find(`a[data-automation-id="job-title-link"]`).Attr("href")

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
