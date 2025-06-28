package scrapers

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/engine"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

type BaseScraper struct {
	Name         string
	Slug         string
	ScheduleHour int
	Engine       *engine.Engine
}

func (b *BaseScraper) GetName() string {
	return b.Name
}

func (b *BaseScraper) GetSlug() string {
	return b.Slug
}

func (b *BaseScraper) GetScheduleHour() int {
	return b.ScheduleHour
}

func (b *BaseScraper) SaveOutput(jobs []types.ScrapedJob) error {
	companiesDir := filepath.Join("scraper_output/companies")
	if err := os.MkdirAll(companiesDir, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(companiesDir, b.Slug+".json")
	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
