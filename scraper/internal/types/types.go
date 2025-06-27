package types

type ScrapedJob struct {
	JobId       string `json:"job_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	City        string `json:"city"`
	State       string `json:"state"`
	Url         string `json:"url"`
	// todo: needs more attributes, need to state which ones are optional, etc.
}

type Scraper interface {
	GetName() string
	GetScheduleHour() int
	ScrapedJobs() []ScrapedJob
	ScrapeJobDetails(job *ScrapedJob)
	SaveOutput(outputDir string) error
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
