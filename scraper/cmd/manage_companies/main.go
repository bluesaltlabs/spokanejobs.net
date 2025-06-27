package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scheduler"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

func main() {
	listFlag := flag.Bool("list", false, "List all available company slugs")
	validateFlag := flag.Bool("validate", false, "Validate the current scheduler configuration")
	generateFlag := flag.Bool("generate", false, "Generate a default scheduler configuration")
	flag.Parse()

	if *listFlag {
		listCompanies()
	} else if *validateFlag {
		validateConfig()
	} else if *generateFlag {
		generateConfig()
	} else {
		fmt.Println("Usage:")
		fmt.Println("  ./manage_companies -list     # List all available companies")
		fmt.Println("  ./manage_companies -validate # Validate current config")
		fmt.Println("  ./manage_companies -generate # Generate default config")
	}
}

func listCompanies() {
	fmt.Println("Available company slugs:")
	companies := scrapers.GetCompanySlugs()
	for i, company := range companies {
		fmt.Printf("%2d. %s\n", i+1, company)
	}
	fmt.Printf("\nTotal: %d companies\n", len(companies))
}

func validateConfig() {
	config, err := scheduler.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	fmt.Println("Validating scheduler configuration...")
	validCount := 0
	invalidCount := 0

	for _, scraper := range config.Scrapers {
		if scrapers.IsValidCompanySlug(scraper.Slug) {
			validCount++
			fmt.Printf("✓ %s (hour %d)\n", scraper.Slug, scraper.Hour)
		} else {
			invalidCount++
			fmt.Printf("✗ %s (hour %d) - INVALID SLUG\n", scraper.Slug, scraper.Hour)
		}
	}

	fmt.Printf("\nValidation complete:\n")
	fmt.Printf("  Valid scrapers: %d\n", validCount)
	fmt.Printf("  Invalid scrapers: %d\n", invalidCount)
	fmt.Printf("  Total configured: %d\n", len(config.Scrapers))
}

func generateConfig() {
	companies := scrapers.GetCompanySlugs()
	config := scheduler.SchedulerConfig{
		Scrapers: make([]scheduler.ScraperConfig, len(companies)),
	}

	// Distribute companies across 24 hours
	for i, company := range companies {
		hour := i % 24
		config.Scrapers[i] = scheduler.ScraperConfig{
			Slug: company,
			Hour: hour,
		}
	}

	// Write to file
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Printf("Error marshaling config: %v", err)
		return
	}

	err = os.WriteFile("configs/scheduler_config.json", data, 0644)
	if err != nil {
		log.Printf("Error writing config file: %v", err)
		return
	}

	fmt.Printf("Generated default configuration with %d companies distributed across 24 hours\n", len(companies))
	fmt.Println("Configuration saved to: configs/scheduler_config.json")
}
