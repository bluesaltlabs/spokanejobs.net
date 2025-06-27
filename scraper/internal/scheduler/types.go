package scheduler

// ScraperConfig represents a single scraper configuration
type ScraperConfig struct {
	Slug string `json:"slug"`
	Hour int    `json:"hour"`
}

// SchedulerConfig represents the overall scheduler configuration
type SchedulerConfig struct {
	Scrapers []ScraperConfig `json:"scrapers"`
}
