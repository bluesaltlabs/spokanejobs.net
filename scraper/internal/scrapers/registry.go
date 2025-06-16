package scrapers

import (
  "gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/openeye"
)

// Run all registered scrapers
func RunAll() {

	openeye.ScrapeJobs()
}


// todo: add a function to run a specific, named scraper.
