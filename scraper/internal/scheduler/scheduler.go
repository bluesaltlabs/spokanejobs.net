package scheduler

import (
	"log"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

// RunScheduledScrapers runs scrapers based on the current hour and configuration
func RunScheduledScrapers() error {
	// Load configuration
	config, err := LoadConfig()
	if err != nil {
		log.Printf("Error loading scheduler config: %v", err)
		return err
	}

	// Get current hour
	now := time.Now()
	currentHour := now.Hour()

	log.Printf("Scheduler mode enabled")
	log.Printf("Loaded %d scraper configurations", len(config.Scrapers))
	log.Printf("Current hour: %d", currentHour)

	// Find scrapers scheduled for current hour
	var scrapersToRun []string
	for _, scraper := range config.Scrapers {
		if scraper.Hour == currentHour {
			// Validate that the scraper exists
			if scrapers.IsValidCompanySlug(scraper.Slug) {
				scrapersToRun = append(scrapersToRun, scraper.Slug)
			} else {
				log.Printf("Warning: Invalid scraper slug '%s' in configuration, skipping", scraper.Slug)
			}
		}
	}

	if len(scrapersToRun) == 0 {
		log.Printf("No scrapers scheduled for hour %d", currentHour)
		return nil
	}

	log.Printf("Running scrapers scheduled for hour %d: %v", currentHour, scrapersToRun)

	// Run each scheduled scraper
	for _, scraperSlug := range scrapersToRun {
		log.Printf("Starting scraper: %s", scraperSlug)
		scrapers.RunScraper(scraperSlug)
		log.Printf("Scraper completed: %s", scraperSlug)
	}

	return nil
}
