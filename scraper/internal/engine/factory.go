package engine

import (
	"fmt"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// EngineFactory creates scraping engines based on configuration
type EngineFactory struct{}

// CreateEngine creates a new scraping engine based on the engine type and configuration
func (f *EngineFactory) CreateEngine(engineType EngineType, config types.ScraperConfig) (ScraperEngine, error) {
	switch engineType {
	case EngineColly:
		return NewCollyEngine(config), nil
	case EngineRod:
		return NewRodEngine(config)
	case "":
		// Default to colly if no engine type specified
		return NewCollyEngine(config), nil
	default:
		return nil, fmt.Errorf("unknown engine type: %s", engineType)
	}
}

// NewEngineFactory creates a new engine factory
func NewEngineFactory() *EngineFactory {
	return &EngineFactory{}
}
