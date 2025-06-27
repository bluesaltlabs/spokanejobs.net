package companies

import (
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

type IntellitectScraper struct {
	*types.BaseScraper
}

func NewIntellitectScraper() *IntellitectScraper {
	return &IntellitectScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "intellitect",
			ScheduleHour: 10,
			Config: types.ScraperConfig{
				BaseURL: "https://intellitect.com/careers/",
			},
		},
	}
}

func (scraper *IntellitectScraper) ScrapeJobs() []types.ScrapedJob {
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

func (scraper *IntellitectScraper) GetName() string {
	return scraper.Name
}

func (scraper *IntellitectScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.ScrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *IntellitectScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}
