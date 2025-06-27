package scrapers

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/companies"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

// scraperInstances holds all active scraper instances
var scraperInstances map[string]types.Scraper

// InitializeScrapers loads all scrapers from configuration
func InitializeScrapers() error {
	// Initialize the scraper instances map
	scraperInstances = make(map[string]types.Scraper)

	// TODO: The Avista scraper is incomplete and causes a runtime panic. Fix before re-enabling.
	// if avista, err := companies.NewAvistaScraper(); err == nil {
	// 	scraperInstances["avista"] = avista
	// } else {
	// 	log.Printf("Failed to initialize avista scraper: %v", err)
	// }

	// TODO: City of Spokane scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["city_of_spokane"] = companies.NewCityOfSpokaneScraper()

	// TODO: Codespeed scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["codespeed"] = companies.NewCodespeedScraper()

	// TODO: CVSD scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["cvsd"] = companies.NewCVSDScraper()

	// TODO: Egnyte scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["egnyte"] = companies.NewEgnyteScraper()

	// TODO: Enhanced Software Products scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["enhanced_software_products"] = companies.NewEnhancedSoftwareProductsScraper()

	// TODO: F5 scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["f5"] = companies.NewF5Scraper()

	// TODO: Gestalt Diagnostics scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["gestalt_diagnostics"] = companies.NewGestaltDiagnosticsScraper()

	// TODO: Gravity Jack scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["gravity_jack"] = companies.NewGravityJackScraper()

	// TODO: Huntwood scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["huntwood"] = companies.NewHuntwoodScraper()

	// TODO: Intellitect scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["intellitect"] = companies.NewIntellitectScraper()

	// TODO: Itron scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["itron"] = companies.NewItronScraper()

	// TODO: Kaiser Aluminum scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["kaiser_aluminum"] = companies.NewKaiserAluminumScraper()

	// TODO: Limelyte scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["limelyte"] = companies.NewLimelyteScraper()

	// TODO: Numerica scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["numerica"] = companies.NewNumericaScraper()

	// Working scraper - successfully finds jobs
	scraperInstances["openeye"] = companies.NewOpenEyeScraper()

	// TODO: Paytrace scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["paytrace"] = companies.NewPaytraceScraper()

	// TODO: Providence scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["providence"] = companies.NewProvidenceScraper()

	// TODO: RiskLens scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["risklens"] = companies.NewRisklensScraper()

	// TODO: SCLD scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["scld"] = companies.NewScldScraper()

	// TODO: SEL scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["sel"] = companies.NewSelScraper()

	// TODO: Spokane Computer scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["spokane_computer"] = companies.NewSpokaneComputerScraper()

	// TODO: Spokane Library scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["spokane_library"] = companies.NewSpokaneLibraryScraper()

	// TODO: SPS scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["sps"] = companies.NewSpsScraper()

	// TODO: STCU scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["stcu"] = companies.NewStcuScraper()

	// TODO: SynergisticIT scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["synergisticit"] = companies.NewSynergisticitScraper()

	// TODO: Treasury4 scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["treasury4"] = companies.NewTreasury4Scraper()

	// TODO: Two Barrels scraper times out after 10 seconds. Increase timeout or investigate site performance.
	// scraperInstances["two_barrels"] = companies.NewTwoBarrelsScraper()

	// TODO: URM scraper returns zero jobs. Update selectors to match current site structure.
	// scraperInstances["urm"] = companies.NewUrmScraper()

	// TODO: Wagstaff scraper returns 404 error. Update careers page URL to match current site structure.
	// scraperInstances["wagstaff"] = companies.NewWagstaffScraper()

	// TODO: Winco scraper returns 403 Forbidden error. Site may be blocking automated access. Consider different user agent or approach.
	// scraperInstances["winco"] = companies.NewWincoScraper()

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

		// Save output to JSON file
		if err := scraper.SaveOutput(); err != nil {
			log.Printf("Error saving output for %s: %v", scraper.GetName(), err)
		}

		log.Printf("Completed scraper: %s", scraper.GetName())
	}

	// Consolidate all company jobs into a single jobs.json file
	if err := utils.ConsolidateJobsToJSON(); err != nil {
		log.Printf("Error consolidating jobs: %v", err)
	} else {
		log.Printf("Successfully consolidated all jobs into jobs.json")
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

		// Save output to JSON file
		if err := scraper.SaveOutput(); err != nil {
			log.Printf("Error saving output for %s: %v", name, err)
		}

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
