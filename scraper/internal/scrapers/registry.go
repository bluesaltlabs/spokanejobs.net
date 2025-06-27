package scrapers

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/avista"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/city_of_spokane"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/codespeed"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/cvsd"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/egnyte"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/enhanced_software_products"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/f5"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/gestalt_diagnostics"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/gravity_jack"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/huntwood"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/intellitect"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/itron"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/kaiser_aluminum"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/limelyte"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/numerica"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/openeye"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/paytrace"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/providence"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/risklens"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/scld"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/sel"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/spokane_computer"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/spokane_library"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/sps"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/stcu"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/synergisticit"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/treasury4"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/two_barrels"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/urm"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/wagstaff"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/winco"
)

// scraperFuncs maps company slugs to their ScrapeJobs functions
var scraperFuncs = map[string]func(){
	"avista":                   func() { avista.ScrapeJobs() },
	"city_of_spokane":          func() { city_of_spokane.ScrapeJobs() },
	"codespeed":                func() { codespeed.ScrapeJobs() },
	"cvsd":                     func() { cvsd.ScrapeJobs() },
	"egnyte":                   func() { egnyte.ScrapeJobs() },
	"enhanced_software_products": func() { enhanced_software_products.ScrapeJobs() },
	"f5":                       func() { f5.ScrapeJobs() },
	"gestalt_diagnostics":      func() { gestalt_diagnostics.ScrapeJobs() },
	"gravity_jack":             func() { gravity_jack.ScrapeJobs() },
	"huntwood":                 func() { huntwood.ScrapeJobs() },
	"intellitect":              func() { intellitect.ScrapeJobs() },
	"itron":                    func() { itron.ScrapeJobs() },
	"kaiser_aluminum":          func() { kaiser_aluminum.ScrapeJobs() },
	"limelyte":                 func() { limelyte.ScrapeJobs() },
	"numerica":                 func() { numerica.ScrapeJobs() },
	"openeye":                  func() { openeye.ScrapeJobs() },
	"paytrace":                 func() { paytrace.ScrapeJobs() },
	"providence":               func() { providence.ScrapeJobs() },
	"risklens":                 func() { risklens.ScrapeJobs() },
	"scld":                     func() { scld.ScrapeJobs() },
	"sel":                      func() { sel.ScrapeJobs() },
	"spokane_computer":         func() { spokane_computer.ScrapeJobs() },
	"spokane_library":          func() { spokane_library.ScrapeJobs() },
	"sps":                      func() { sps.ScrapeJobs() },
	"stcu":                     func() { stcu.ScrapeJobs() },
	"synergisticit":            func() { synergisticit.ScrapeJobs() },
	"treasury4":                func() { treasury4.ScrapeJobs() },
	"two_barrels":              func() { two_barrels.ScrapeJobs() },
	"urm":                      func() { urm.ScrapeJobs() },
	"wagstaff":                 func() { wagstaff.ScrapeJobs() },
	"winco":                    func() { winco.ScrapeJobs() },
}

// RunAll runs all registered scrapers using CompanySlugs
func RunAll() {
	for _, slug := range CompanySlugs {
		if fn, ok := scraperFuncs[slug]; ok {
			fn()
		} else {
			log.Printf("Scraper for %s not found\n", slug)
		}
	}
}

// RunScraper runs a single scraper by slug
func RunScraper(name string) {
	if fn, ok := scraperFuncs[name]; ok {
		fn()
	} else {
		log.Printf("Scraper %s not found\n", name)
	}
}
