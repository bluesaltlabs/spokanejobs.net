package scheduler

import (
	"log"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

// RunScheduledScrapers runs scrapers based on the current hour and their individual schedule settings
func RunScheduledScrapers() error {
	// Get current hour
	now := time.Now()
	currentHour := now.Hour()

	log.Printf("Scheduler mode enabled")
	log.Printf("Current hour: %d", currentHour)

	// Get all available scrapers
	allScrapers := scrapers.GetAllScrapers()
	log.Printf("Loaded %d scrapers", len(allScrapers))

	// Find scrapers scheduled for current hour
	var scrapersToRun []string
	for _, scraper := range allScrapers {
		if scraper.GetScheduleHour() == currentHour {
			scrapersToRun = append(scrapersToRun, scraper.GetName())
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
