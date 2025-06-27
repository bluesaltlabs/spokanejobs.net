package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scheduler"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("=== Scraper Init ===\n")
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	// Load environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Try .env.local for local development
		if err := godotenv.Load(".env.local"); err != nil {
			log.Printf("No .env or .env.local file found, using system environment variables")
		} else {
			log.Printf("Loaded environment variables from .env.local")
		}
	} else {
		log.Printf("Loaded environment variables from .env")
	}

	// Log environment variables for debugging
	if dataRepoPath := os.Getenv("DATA_REPO_PATH"); dataRepoPath != "" {
		log.Printf("DATA_REPO_PATH: %s", dataRepoPath)
	} else {
		log.Printf("DATA_REPO_PATH not set")
	}

	if dataRepoSubdir := os.Getenv("DATA_REPO_SUBDIR"); dataRepoSubdir != "" {
		log.Printf("DATA_REPO_SUBDIR: %s", dataRepoSubdir)
	} else {
		log.Printf("DATA_REPO_SUBDIR not set")
	}

	// Check for SCRAPER_ARGS environment variable
	if scraperArgs := os.Getenv("SCRAPER_ARGS"); scraperArgs != "" {
		log.Printf("SCRAPER_ARGS environment variable found: %s", scraperArgs)
		// Parse the environment variable arguments
		args := strings.Fields(scraperArgs)
		processArgs(args)
		return
	}

	// Define command line flags
	runAll := flag.Bool("all", false, "Run all scrapers")
	flag.Parse()

	// Check for additional arguments (specific scraper)
	args := flag.Args()
	processArgsWithAllFlag(args, *runAll)
}

// processArgs handles the argument processing logic
func processArgs(args []string) {
	if len(args) == 0 {
		// Default behavior: run scheduler mode
		if err := scheduler.RunScheduledScrapers(); err != nil {
			log.Printf("Error running scheduled scrapers: %v", err)
			os.Exit(1)
		}
		return
	}

	// Check if first argument is -all
	if args[0] == "-all" {
		fmt.Printf("Running all scrapers\n")
		scrapers.RunAll()
		return
	}

	// Run specific scraper(s)
	if len(args) == 1 {
		scraperName := args[0]
		fmt.Printf("Running specific scraper: %s\n", scraperName)
		scrapers.RunScraper(scraperName)
	} else {
		// Run multiple specific scrapers
		fmt.Printf("Running multiple scrapers: %v\n", args)
		for _, scraperName := range args {
			fmt.Printf("Running scraper: %s\n", scraperName)
			scrapers.RunScraper(scraperName)
		}
	}
}

// processArgsWithAllFlag handles the argument processing with the -all flag from command line
func processArgsWithAllFlag(args []string, runAll bool) {
	if runAll {
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
