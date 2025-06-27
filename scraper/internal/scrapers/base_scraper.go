package scrapers

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

// BaseScraperWrapper wraps the BaseScraper type to allow method definitions
type BaseScraperWrapper struct {
	*types.BaseScraper
}

// NewBaseScraper creates a new base scraper with the given configuration
func NewBaseScraper(name string, config types.ScraperConfig) types.Scraper {
	base := &types.BaseScraper{
		Jobs:   make([]types.ScrapedJob, 0),
		Name:   name,
		Config: config,
	}
	return &BaseScraperWrapper{BaseScraper: base}
}

// GetName returns the scraper name
func (b *BaseScraperWrapper) GetName() string {
	return b.Name
}

// ScrapedJobs returns the scraped jobs, scraping if necessary
func (b *BaseScraperWrapper) ScrapedJobs() []types.ScrapedJob {
	if len(b.Jobs) == 0 {
		b.Jobs = b.scrapeJobs()
	}
	return b.Jobs
}

// ScrapeJobDetails implements the interface method (can be overridden)
func (b *BaseScraperWrapper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden by custom scrapers
}

// scrapeJobs performs the actual scraping using the configuration
func (b *BaseScraperWrapper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := utils.NewCollector(b.Config)

	// Display all jobs after scraping completes as json to the standard output
	c.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	// Process job elements using configured selectors
	c.OnHTML(b.Config.Selectors.JobContainer, func(h *colly.HTMLElement) {
		job := b.extractJobFromElement(h)
		if job != nil {
			jobs = append(jobs, *job)
		}
	})

	// Visit the configured URL
	c.Visit(b.Config.BaseURL)

	log.Printf("\n-----\n\nColly instance done for %s: %+v\n\n", b.Name, c)

	return jobs
}

// extractJobFromElement extracts job data from an HTML element using configured selectors
func (b *BaseScraperWrapper) extractJobFromElement(h *colly.HTMLElement) *types.ScrapedJob {
	job := &types.ScrapedJob{}
	selection := h.DOM

	// Extract title
	if b.Config.Selectors.Title != "" {
		job.Title = strings.TrimSpace(selection.Find(b.Config.Selectors.Title).Text())
	}

	// Extract URL
	if b.Config.Selectors.URL != "" {
		url, exists := selection.Find(b.Config.Selectors.URL).Attr("href")
		if exists {
			job.Url = url
			// Extract job ID from URL if prefix is configured
			if b.Config.JobURLPrefix != "" && strings.HasPrefix(url, b.Config.JobURLPrefix) {
				job.JobId = strings.TrimPrefix(url, b.Config.JobURLPrefix)
			}
		}
	}

	// Extract description
	if b.Config.Selectors.Description != "" {
		job.Description = strings.TrimSpace(selection.Find(b.Config.Selectors.Description).Text())
	}

	// Extract city
	if b.Config.Selectors.City != "" {
		job.City = strings.TrimSpace(selection.Find(b.Config.Selectors.City).Text())
	}

	// Extract state
	if b.Config.Selectors.State != "" {
		job.State = strings.TrimSpace(selection.Find(b.Config.Selectors.State).Text())
	}

	// Only return job if we have at least a title or URL
	if job.Title != "" || job.Url != "" {
		return job
	}

	return nil
}
