package scrapers

import (
	"fmt"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/openeye"
)

// Run all registered scrapers
func RunAll() {

	openeye.ScrapeJobs()
}

func RunScraper(name string) {
	switch name {
		case "openeye":
			openeye.ScrapeJobs()
		default:
			fmt.Errorf("Scraper %s not found\n", name)
	}
}
