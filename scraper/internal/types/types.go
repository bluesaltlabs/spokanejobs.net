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
}

// New configuration-driven types
type ScraperConfig struct {
	AllowedDomains []string     `json:"allowedDomains"`
	BaseURL        string       `json:"baseURL"`
	JobURLPrefix   string       `json:"jobURLPrefix"`
	Selectors      JobSelectors `json:"selectors"`
}

type JobSelectors struct {
	JobContainer string `json:"jobContainer"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Description  string `json:"description"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type BaseScraper struct {
	Jobs         []ScrapedJob
	Name         string
	Config       ScraperConfig
	ScheduleHour int // Hour of day when this scraper should run (0-23)
}

// GetScheduleHour returns the hour when this scraper should run
func (b *BaseScraper) GetScheduleHour() int {
	return b.ScheduleHour
}

// ConfigDrivenScraper implements the Scraper interface using configuration
type ConfigDrivenScraper struct {
	BaseScraper
}
