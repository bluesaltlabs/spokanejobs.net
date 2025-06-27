package companies

import (
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

type SpokaneLibraryScraper struct {
	*types.BaseScraper
}

func NewSpokaneLibraryScraper() *SpokaneLibraryScraper {
	return &SpokaneLibraryScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "spokane_library",
			ScheduleHour: 22,
			Config: types.ScraperConfig{
				BaseURL: "https://www.spokanelibrary.org/careers",
			},
		},
	}
}

func (scraper *SpokaneLibraryScraper) ScrapeJobs() []types.ScrapedJob {
	var jobs []types.ScrapedJob
	c := utils.NewCollector(scraper.Config)

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

func (scraper *SpokaneLibraryScraper) GetName() string {
	return scraper.Name
}

func (scraper *SpokaneLibraryScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.ScrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *SpokaneLibraryScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}
