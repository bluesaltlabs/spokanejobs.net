package companies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/engine"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

var avistaScrapeUrlPrefix string = "https://recruiting2.ultipro.com/AVI1000AMAST/JobBoard/362abf68-95c3-4b17-a39d-76a6efe5ff18"
var avistaScrapeUrl string = avistaScrapeUrlPrefix + "?o=postedDateDesc&f4=ITpN1FEJvlWudZe0AayqWA"
var avistaJobUrlPrefix string = avistaScrapeUrlPrefix + "/OpportunityDetail?opportunityId="

// AvistaScraper implements a scraper for Avista using the simplified engine approach with rod
type AvistaScraper struct {
	*types.BaseScraper
	Engine *engine.RodEngine
}

func NewAvistaScraper() (*AvistaScraper, error) {
	config := types.ScraperConfig{
		BaseURL:        "https://www.myavista.com/about-us/working-at-avista/careers",
		AllowedDomains: []string{"recruiting2.ultipro.com"},
	}

	rodEngine, err := engine.NewRodEngine(config)
	if err != nil {
		return nil, err
	}

	return &AvistaScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "avista",
			ScheduleHour: 0,
			Config:       config,
		},
		Engine: rodEngine,
	}, nil
}

func (a *AvistaScraper) GetName() string {
	return a.Name
}

func (a *AvistaScraper) ScrapedJobs() []types.ScrapedJob {
	if len(a.Jobs) == 0 {
		a.Jobs = a.scrapeJobs()
	}
	return a.Jobs
}

func (a *AvistaScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (a *AvistaScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(a.Jobs, a.Name, outputDir)
}

func (a *AvistaScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)

	// Get the rod page directly from the engine
	page := a.Engine.GetPage()

	log.Printf("visiting: %s\n", avistaScrapeUrl)

	// Navigate to the page and wait for it to load
	err := page.Navigate(avistaScrapeUrl)
	if err != nil {
		log.Printf("Error navigating to %s: %v", avistaScrapeUrl, err)
		return jobs
	}

	// Wait for the page to load completely
	page.MustWaitLoad()

	// Wait for the job listings to appear (div.opportunity elements)
	page.MustWait("div.opportunity")

	// Find all job opportunity elements
	elements := page.MustElements("div.opportunity")

	log.Printf("Found %d job opportunities", len(elements))

	// Process each job element
	for _, element := range elements {
		j := &types.ScrapedJob{}

		// Find the link element within this opportunity
		linkElement, err := element.Element("h3 a.opportunity-link")
		if err != nil {
			log.Printf("Error finding link element: %v", err)
			continue
		}

		// Get the href attribute
		url, err := linkElement.Attribute("href")
		if err != nil || url == nil {
			log.Printf("Error getting href attribute: %v", err)
			continue
		}

		// Get the text content (job title)
		title, err := linkElement.Text()
		if err != nil {
			log.Printf("Error getting title text: %v", err)
			continue
		}

		// Extract job ID from URL
		job_id := strings.TrimPrefix(*url, avistaJobUrlPrefix)

		j.JobId = job_id
		j.Url = *url
		j.Title = strings.TrimSpace(title)

		jobs = append(jobs, *j)
		log.Printf("Found job: %s - %s", j.JobId, j.Title)
	}

	// Save to file
	if err := utils.SaveJobsToJSON(jobs, a.Name, "scraper_output"); err != nil {
		log.Printf("Error saving jobs to JSON for %s: %v", a.Name, err)
	} else {
		log.Printf("Saved %d jobs to JSON file for %s", len(jobs), a.Name)
	}

	// Also print to stdout for backward compatibility
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(jobs)

	log.Printf("\n-----\n\nRod instance done: %+v\n\n", page)
	return jobs
}
