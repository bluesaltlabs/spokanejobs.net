package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scheduler"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

func main() {
	fmt.Printf("=== Scraper Init ===\n")
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	// Define command line flags
	runAll := flag.Bool("all", false, "Run all scrapers")
	flag.Parse()

	// Check for additional arguments (specific scraper)
	args := flag.Args()

	if *runAll {
		// run all scrapers
		fmt.Printf("Running all scrapers\n")
		scrapers.RunAll()
	} else if len(args) == 1 {
		// Run specific scraper
		scraperName := args[0]
		fmt.Printf("Running specific scraper: %s\n", scraperName)
		scrapers.RunScraper(scraperName)
	} else if len(args) > 1 {
		log.Printf("Too many arguments provided: %v", args)
		os.Exit(1)
	} else {
		// Default behavior: run scheduler mode
		if err := scheduler.RunScheduledScrapers(); err != nil {
			log.Printf("Error running scheduled scrapers: %v", err)
			os.Exit(1)
		}
	}
}
