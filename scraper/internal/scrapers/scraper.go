package scrapers

import (
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/models"
)

type Scraper interface {
	GetName() 			string
  ScrapedJobs() 	[]models.ScrapedJob
  ScrapeJobDetails(job *models.ScrapedJob)
}

// todo: I don't know what I'm doing here. I haven't figured out how to extend from this properly yet.
