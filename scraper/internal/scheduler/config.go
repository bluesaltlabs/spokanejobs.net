package scheduler

import (
	"encoding/json"
	"log"
	"os"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

// LoadConfig loads the scheduler configuration from the configs/scheduler_config.json file
func LoadConfig() (*SchedulerConfig, error) {
	configPath := "configs/scheduler_config.json"

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("Warning: Config file %s not found, using default configuration", configPath)
		return getDefaultConfig(), nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return getDefaultConfig(), nil
	}

	// Parse JSON
	var config SchedulerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Printf("Error parsing config file: %v", err)
		return getDefaultConfig(), nil
	}

	return &config, nil
}

// getDefaultConfig returns a default configuration with all scrapers at hour 0
func getDefaultConfig() *SchedulerConfig {
	companySlugs := scrapers.GetCompanySlugs()
	scrapers := make([]ScraperConfig, len(companySlugs))

	for i, slug := range companySlugs {
		scrapers[i] = ScraperConfig{
			Slug: slug,
			Hour: 0,
		}
	}

	return &SchedulerConfig{
		Scrapers: scrapers,
	}
}
