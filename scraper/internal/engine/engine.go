package engine

import (
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// EngineType represents the type of scraping engine
type EngineType string

const (
	EngineColly EngineType = "colly"
	EngineRod   EngineType = "rod"
)

// Element represents a DOM element for scraping
type Element interface {
	Attr(name string) string
	Text() string
	Find(selector string) Element
}

// ScraperEngine provides a common interface for different scraping engines
type ScraperEngine interface {
	OnScraped(callback func())
	OnHTML(selector string, callback func(Element))
	Visit(url string) error
	Close()
}

// BaseEngine provides a minimal common interface for different scraping engines
type BaseEngine struct {
	Config types.ScraperConfig
}

// NewBaseEngine creates a new base engine with the given configuration
func NewBaseEngine(config types.ScraperConfig) *BaseEngine {
	return &BaseEngine{
		Config: config,
	}
}

// GetConfig returns the engine configuration
func (b *BaseEngine) GetConfig() types.ScraperConfig {
	return b.Config
}
