package scheduler

import (
	"encoding/json"
	"log"
	"os"
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
	return &SchedulerConfig{
		Scrapers: []ScraperConfig{
			{Slug: "avista", Hour: 0},
			{Slug: "city_of_spokane", Hour: 0},
			{Slug: "codespeed", Hour: 0},
			{Slug: "cvsd", Hour: 0},
			{Slug: "egnyte", Hour: 0},
			{Slug: "enhanced_software_products", Hour: 0},
			{Slug: "f5", Hour: 0},
			{Slug: "gestalt_diagnostics", Hour: 0},
			{Slug: "gravity_jack", Hour: 0},
			{Slug: "huntwood", Hour: 0},
			{Slug: "intellitect", Hour: 0},
			{Slug: "itron", Hour: 0},
			{Slug: "kaiser_aluminum", Hour: 0},
			{Slug: "limelyte", Hour: 0},
			{Slug: "numerica", Hour: 0},
			{Slug: "openeye", Hour: 0},
			{Slug: "paytrace", Hour: 0},
			{Slug: "providence", Hour: 0},
			{Slug: "risklens", Hour: 0},
			{Slug: "scld", Hour: 0},
			{Slug: "sel", Hour: 0},
			{Slug: "spokane_computer", Hour: 0},
			{Slug: "spokane_library", Hour: 0},
			{Slug: "sps", Hour: 0},
			{Slug: "stcu", Hour: 0},
			{Slug: "synergisticit", Hour: 0},
			{Slug: "treasury4", Hour: 0},
			{Slug: "two_barrels", Hour: 0},
			{Slug: "urm", Hour: 0},
			{Slug: "wagstaff", Hour: 0},
			{Slug: "winco", Hour: 0},
		},
	}
}
