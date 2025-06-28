package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scheduler"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
	"github.com/joho/godotenv"

	// Import companies package to ensure scrapers are registered
	_ "gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/companies"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Command line flags
	runAll := flag.Bool("all", false, "Run all scrapers regardless of schedule")
	debug := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	// Check for SCRAPER_ARGS environment variable
	scraperArgs := os.Getenv("SCRAPER_ARGS")
	if scraperArgs != "" {
		log.Printf("Found SCRAPER_ARGS: %s", scraperArgs)
		// Parse SCRAPER_ARGS as if they were command line arguments
		args := strings.Fields(scraperArgs)
		flag.CommandLine.Parse(args)
	}

	// Auto-register all scrapers via init() in their packages
	allScrapers := scrapers.LoadAllScrapers()

	if *debug {
		log.Printf("Loaded %d scrapers", len(allScrapers))
		for _, scraper := range allScrapers {
			log.Printf("  - %s (scheduled for hour: %d)",
				scraper.GetName(), scraper.GetScheduleHour())
		}
	}

	sched := scheduler.New(allScrapers)

	if *runAll {
		log.Println("Running all scrapers (--all flag specified)")
		sched.RunAllScrapers()
	} else {
		now := time.Now()
		log.Printf("Current time: %s (hour: %d)", now.Format("2006-01-02 15:04:05"), now.Hour())
		sched.RunDueScrapers(now)
	}
}
