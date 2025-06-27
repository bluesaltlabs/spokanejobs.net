package scrapers

import (
	"encoding/json"
	"log"
	"os"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/git"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

// BaseScraperWrapper wraps the BaseScraper type to allow method definitions
type BaseScraperWrapper struct {
	*types.BaseScraper
}

// NewBaseScraper creates a new base scraper with the given configuration
func NewBaseScraper(name string, config types.ScraperConfig) types.Scraper {
	base := &types.BaseScraper{
		Jobs:   make([]types.ScrapedJob, 0),
		Name:   name,
		Config: config,
	}

	return &BaseScraperWrapper{BaseScraper: base}
}

// GetName returns the scraper name
func (b *BaseScraperWrapper) GetName() string {
	return b.Name
}

// ScrapedJobs returns the scraped jobs, scraping if necessary
func (b *BaseScraperWrapper) ScrapedJobs() []types.ScrapedJob {
	if len(b.Jobs) == 0 {
		b.Jobs = b.scrapeJobs()
	}
	return b.Jobs
}

// ScrapeJobDetails implements the interface method (can be overridden)
func (b *BaseScraperWrapper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden by custom scrapers
}

// SaveOutput saves the scraped jobs to a JSON file
func (b *BaseScraperWrapper) SaveOutput(outputDir string) error {
	err := utils.SaveJobsToJSON(b.Jobs, b.Name, outputDir)

	// Optionally sync to data repo if configured
	dataRepoPath := os.Getenv("DATA_REPO_PATH")
	dataRepoSubdir := os.Getenv("DATA_REPO_SUBDIR")
	if dataRepoPath != "" {
		gs := git.NewGitSync(dataRepoPath, b.Name, dataRepoSubdir)
		err2 := gs.SyncJobs(b.Jobs)
		if err2 != nil {
			log.Printf("[GitSync] Failed to sync jobs for %s: %v", b.Name, err2)
		} else {
			log.Printf("[GitSync] Successfully synced jobs for %s to data repo", b.Name)
		}
	}

	return err
}

// SetEngine sets the scraping engine for this scraper
func (b *BaseScraperWrapper) SetEngine(engine interface{}) {
	b.Engine = engine
}

// scrapeJobs performs the actual scraping using the configuration
func (b *BaseScraperWrapper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)

	// The base scraper doesn't implement scraping logic
	// Individual scrapers should override this method or implement their own scraping
	log.Printf("Base scraper scrapeJobs called for %s - this should be overridden by individual scrapers", b.Name)

	// Save jobs to JSON file
	if err := utils.SaveJobsToJSON(jobs, b.Name, "scraper_output"); err != nil {
		log.Printf("Error saving jobs to JSON for %s: %v", b.Name, err)
	} else {
		log.Printf("Saved %d jobs to JSON file for %s", len(jobs), b.Name)
	}

	// Also print to stdout for backward compatibility
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(jobs)

	log.Printf("\n-----\n\nScraping completed for %s\n\n", b.Name)

	return jobs
}
