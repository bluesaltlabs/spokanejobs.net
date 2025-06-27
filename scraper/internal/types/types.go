package types

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ScrapedJob struct {
	JobId       string `json:"job_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	City        string `json:"city"`
	State       string `json:"state"`
	Url         string `json:"url"`
	Company     string `json:"company"`
	// New fields for enhanced content
	RichDescription string `json:"rich_description,omitempty"`
	PostingDate     string `json:"posting_date,omitempty"`
	SalaryRange     string `json:"salary_range,omitempty"`
	// todo: needs more attributes, need to state which ones are optional, etc.
}

type Scraper interface {
	GetName() string
	GetScheduleHour() int
	ScrapedJobs() []ScrapedJob
	ScrapeJobDetails(job *ScrapedJob)
	SaveOutput() error
	// SetEngine is optional - scrapers can implement it if they want to use the new engine interface
	// Default implementation is provided in BaseScraper
}

// ScraperConfig is kept minimal for backward compatibility
type ScraperConfig struct {
	AllowedDomains []string `json:"allowedDomains"`
	BaseURL        string   `json:"baseURL"`
	JobURLPrefix   string   `json:"jobURLPrefix"`
}

type BaseScraper struct {
	Jobs         []ScrapedJob
	Name         string
	Config       ScraperConfig
	ScheduleHour int         // Hour of day when this scraper should run (0-23)
	Engine       interface{} // New field for the engine
}

// GetScheduleHour returns the hour when this scraper should run
func (b *BaseScraper) GetScheduleHour() int {
	return b.ScheduleHour
}

// SetEngine provides a default implementation for the optional SetEngine method
func (b *BaseScraper) SetEngine(engine interface{}) {
	b.Engine = engine
}

// SaveOutput saves the scraped jobs to a JSON file
func (b *BaseScraper) SaveOutput() error {
	// Create companies directory within the output directory
	companiesDir := filepath.Join("scraper_output/companies")
	if err := os.MkdirAll(companiesDir, 0755); err != nil {
		return fmt.Errorf("failed to create companies directory: %w", err)
	}

	// Create filename without timestamp for consistency
	filename := fmt.Sprintf("%s.json", b.Name)
	filepath := filepath.Join(companiesDir, filename)

	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filepath, err)
	}
	defer file.Close()

	// Create JSON encoder with indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Encode the jobs
	if err := encoder.Encode(b.Jobs); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
