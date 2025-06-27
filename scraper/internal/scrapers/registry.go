package scrapers

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/companies"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// scraperInstances holds all active scraper instances
var scraperInstances map[string]types.Scraper

// InitializeScrapers loads all scrapers from configuration
func InitializeScrapers() error {
	// Initialize the scraper instances map
	scraperInstances = make(map[string]types.Scraper)

	// Map company slugs to scrapers.
	scraperInstances["openeye"] = companies.NewOpenEyeScraper()
	scraperInstances["avista"] = companies.NewAvistaScraper()
	scraperInstances["f5"] = companies.NewF5Scraper()
	scraperInstances["numerica"] = companies.NewNumericaScraper()
	scraperInstances["gravity_jack"] = companies.NewGravityJackScraper()
	scraperInstances["egnyte"] = companies.NewEgnyteScraper()
	scraperInstances["city_of_spokane"] = companies.NewCityOfSpokaneScraper()
	scraperInstances["codespeed"] = companies.NewCodespeedScraper()
	scraperInstances["cvsd"] = companies.NewCVSDScraper()
	scraperInstances["enhanced_software_products"] = companies.NewEnhancedSoftwareProductsScraper()
	scraperInstances["huntwood"] = companies.NewHuntwoodScraper()
	scraperInstances["kaiser_aluminum"] = companies.NewKaiserAluminumScraper()
	scraperInstances["limelyte"] = companies.NewLimelyteScraper()
	scraperInstances["itron"] = companies.NewItronScraper()
	scraperInstances["two_barrels"] = companies.NewTwoBarrelsScraper()
	scraperInstances["treasury4"] = companies.NewTreasury4Scraper()
	scraperInstances["scld"] = companies.NewScldScraper()
	scraperInstances["sel"] = companies.NewSelScraper()
	scraperInstances["urm"] = companies.NewUrmScraper()
	scraperInstances["paytrace"] = companies.NewPaytraceScraper()
	scraperInstances["winco"] = companies.NewWincoScraper()
	scraperInstances["intellitect"] = companies.NewIntellitectScraper()
	scraperInstances["risklens"] = companies.NewRisklensScraper()
	scraperInstances["spokane_computer"] = companies.NewSpokaneComputerScraper()
	scraperInstances["providence"] = companies.NewProvidenceScraper()
	scraperInstances["wagstaff"] = companies.NewWagstaffScraper()
	scraperInstances["gestalt_diagnostics"] = companies.NewGestaltDiagnosticsScraper()
	scraperInstances["synergisticit"] = companies.NewSynergisticitScraper()
	scraperInstances["spokane_library"] = companies.NewSpokaneLibraryScraper()
	scraperInstances["stcu"] = companies.NewStcuScraper()
	scraperInstances["sps"] = companies.NewSpsScraper()

	log.Printf("Initialized %d scrapers from configuration (with company scrapers)", len(scraperInstances))
	return nil
}

// RunAll runs all registered scrapers using CompanySlugs
func RunAll() {
	if scraperInstances == nil {
		if err := InitializeScrapers(); err != nil {
			log.Printf("Failed to initialize scrapers: %v", err)
			return
		}
	}

	// Run each scraper instance
	for _, scraper := range scraperInstances {
		log.Printf("Running scraper: %s", scraper.GetName())
		_ = scraper.ScrapedJobs()
		log.Printf("Completed scraper: %s", scraper.GetName())
	}
}

// RunScraper runs a single scraper by slug
func RunScraper(name string) {
	if scraperInstances == nil {
		if err := InitializeScrapers(); err != nil {
			log.Printf("Failed to initialize scrapers: %v", err)
			return
		}
	}

	if scraper, ok := scraperInstances[name]; ok {
		log.Printf("Running scraper: %s", name)
		_ = scraper.ScrapedJobs()
		log.Printf("Completed scraper: %s", name)
	} else {
		log.Printf("Scraper %s not found in configuration", name)
	}
}

// GetScraper returns a scraper instance by name
func GetScraper(name string) (types.Scraper, bool) {
	if scraperInstances == nil {
		if err := InitializeScrapers(); err != nil {
			log.Printf("Failed to initialize scrapers: %v", err)
			return nil, false
		}
	}

	scraper, exists := scraperInstances[name]
	return scraper, exists
}

// GetAllScrapers returns all scraper instances as a slice
func GetAllScrapers() []types.Scraper {
	if scraperInstances == nil {
		if err := InitializeScrapers(); err != nil {
			log.Printf("Failed to initialize scrapers: %v", err)
			return nil
		}
	}

	scrapers := make([]types.Scraper, 0, len(scraperInstances))
	for _, scraper := range scraperInstances {
		scrapers = append(scrapers, scraper)
	}
	return scrapers
}
