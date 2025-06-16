package scrapers

import (
	"fmt"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/avista"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/gestalt_diagnostics"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/numerica"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/treasury4"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/city_of_spokane"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/gravity_jack"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/openeye"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/sel"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/two_barrels"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/codespeed"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/huntwood"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/paytrace"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/spokane_computer"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/urm"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/cvsd"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/intellitect"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/providence"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/spokane_library"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/wagstaff"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/egnyte"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/itron"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/sps"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/winco"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/enhanced_software_products"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/kaiser_aluminum"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/risklens"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/stcu"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/f5"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/limelyte"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/scld"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers/synergisticit"
)

// Run all registered scrapers
func RunAll() {
	avista.ScrapeJobs()
	gestalt_diagnostics.ScrapeJobs()
	numerica.ScrapeJobs()
	treasury4.ScrapeJobs()
	city_of_spokane.ScrapeJobs()
	gravity_jack.ScrapeJobs()
	openeye.ScrapeJobs()
	sel.ScrapeJobs()
	two_barrels.ScrapeJobs()
	codespeed.ScrapeJobs()
	huntwood.ScrapeJobs()
	paytrace.ScrapeJobs()
	spokane_computer.ScrapeJobs()
	urm.ScrapeJobs()
	cvsd.ScrapeJobs()
	intellitect.ScrapeJobs()
	providence.ScrapeJobs()
	spokane_library.ScrapeJobs()
	wagstaff.ScrapeJobs()
	egnyte.ScrapeJobs()
	itron.ScrapeJobs()
	sps.ScrapeJobs()
	winco.ScrapeJobs()
	enhanced_software_products.ScrapeJobs()
	kaiser_aluminum.ScrapeJobs()
	risklens.ScrapeJobs()
	stcu.ScrapeJobs()
	f5.ScrapeJobs()
	limelyte.ScrapeJobs()
	scld.ScrapeJobs()
	synergisticit.ScrapeJobs()
}

func RunScraper(name string) {
	switch name {
		case "avista":
		  avista.ScrapeJobs()
		case "gestalt_diagnostics":
		  gestalt_diagnostics.ScrapeJobs()
		case "numerica":
		  numerica.ScrapeJobs()
		case "treasury4":
		  treasury4.ScrapeJobs()
		case "city_of_spokane":
		  city_of_spokane.ScrapeJobs()
		case "gravity_jack":
		  gravity_jack.ScrapeJobs()
		case "openeye":
		  openeye.ScrapeJobs()
		case "sel":
		  sel.ScrapeJobs()
		case "two_barrels":
		  two_barrels.ScrapeJobs()
		case "codespeed":
		  codespeed.ScrapeJobs()
		case "huntwood":
		  huntwood.ScrapeJobs()
		case "paytrace":
		  paytrace.ScrapeJobs()
		case "spokane_computer":
		  spokane_computer.ScrapeJobs()
		case "urm":
		  urm.ScrapeJobs()
		case "cvsd":
		  cvsd.ScrapeJobs()
		case "intellitect":
		  intellitect.ScrapeJobs()
		case "providence":
		  providence.ScrapeJobs()
		case "spokane_library":
		  spokane_library.ScrapeJobs()
		case "wagstaff":
		  wagstaff.ScrapeJobs()
		case "egnyte":
		  egnyte.ScrapeJobs()
		case "itron":
		  itron.ScrapeJobs()
		case "sps":
		  sps.ScrapeJobs()
		case "winco":
		  winco.ScrapeJobs()
		case "enhanced_software_products":
		  enhanced_software_products.ScrapeJobs()
		case "kaiser_aluminum":
		  kaiser_aluminum.ScrapeJobs()
		case "risklens":
		  risklens.ScrapeJobs()
		case "stcu":
		  stcu.ScrapeJobs()
		case "f5":
		  f5.ScrapeJobs()
		case "limelyte":
		  limelyte.ScrapeJobs()
		case "scld":
		  scld.ScrapeJobs()
		case "synergisticit":
		  synergisticit.ScrapeJobs()
		default:
			fmt.Errorf("Scraper %s not found\n", name)
	}
}
