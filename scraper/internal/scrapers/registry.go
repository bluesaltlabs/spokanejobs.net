package scrapers

import (
	"sync"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

var (
	registryMu sync.Mutex
	registry   = map[string]func() types.Scraper{}
)

// Register allows company scrapers to self-register
func Register(name, slug string, constructor func() types.Scraper) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[name] = constructor
}

// LoadAllScrapers returns all registered scrapers
func LoadAllScrapers() []types.Scraper {
	registryMu.Lock()
	defer registryMu.Unlock()
	scrapers := make([]types.Scraper, 0, len(registry))
	for _, constructor := range registry {
		scrapers = append(scrapers, constructor())
	}
	return scrapers
}
