package scrapers

import (
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

type Scraper interface {
	GetName() 			string
  ScrapedJobs() 	[]types.ScrapedJob
  ScrapeJobDetails(job *types.ScrapedJob)
}

// todo: I don't know what I'm doing here. I haven't figured out how to extend from this properly yet.
