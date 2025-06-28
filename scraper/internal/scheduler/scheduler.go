package scheduler

import (
	"log"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

type Scheduler struct {
	Scrapers []types.Scraper
}

func New(scrapers []types.Scraper) *Scheduler {
	return &Scheduler{Scrapers: scrapers}
}

func (s *Scheduler) RunDueScrapers(now time.Time) {
	hour := now.Hour()
	for _, scraper := range s.Scrapers {
		if scraper.GetScheduleHour() == hour {
			log.Printf("Running scheduled scraper: %s", scraper.GetName())
			jobs := scraper.ScrapedJobs()
			_ = scraper.SaveOutput(jobs)
		}
	}
}

func (s *Scheduler) RunAllScrapers() {
	for _, scraper := range s.Scrapers {
		log.Printf("Running scraper: %s", scraper.GetName())
		jobs := scraper.ScrapedJobs()
		_ = scraper.SaveOutput(jobs)
	}
	log.Printf("Consolidating all jobs and syncing to git repo...")
	err := utils.ConsolidateJobsToJSON()
	if err != nil {
		log.Printf("Error during consolidation and git sync: %v", err)
	} else {
		log.Printf("Consolidation and git sync complete.")
	}
}
