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

// Run all registered scrapers
func RunAll() {
	avista.ScrapeJobs()
	city_of_spokane.ScrapeJobs()
	codespeed.ScrapeJobs()
	cvsd.ScrapeJobs()
	egnyte.ScrapeJobs()
	enhanced_software_products.ScrapeJobs()
	f5.ScrapeJobs()
	gestalt_diagnostics.ScrapeJobs()
	gravity_jack.ScrapeJobs()
	huntwood.ScrapeJobs()
	intellitect.ScrapeJobs()
	itron.ScrapeJobs()
	kaiser_aluminum.ScrapeJobs()
	limelyte.ScrapeJobs()
	numerica.ScrapeJobs()
	openeye.ScrapeJobs()
	paytrace.ScrapeJobs()
	providence.ScrapeJobs()
	risklens.ScrapeJobs()
	scld.ScrapeJobs()
	sel.ScrapeJobs()
	spokane_computer.ScrapeJobs()
	spokane_library.ScrapeJobs()
	sps.ScrapeJobs()
	stcu.ScrapeJobs()
	synergisticit.ScrapeJobs()
	treasury4.ScrapeJobs()
	two_barrels.ScrapeJobs()
	urm.ScrapeJobs()
	wagstaff.ScrapeJobs()
	winco.ScrapeJobs()
}

func RunScraper(name string) {
	switch name {
		case "avista":
			avista.ScrapeJobs()
		case "city_of_spokane":
			city_of_spokane.ScrapeJobs()
		case "codespeed":
			codespeed.ScrapeJobs()
		case "cvsd":
			cvsd.ScrapeJobs()
		case "egnyte":
			egnyte.ScrapeJobs()
		case "enhanced_software_products":
			enhanced_software_products.ScrapeJobs()
		case "f5":
			f5.ScrapeJobs()
		case "gestalt_diagnostics":
			gestalt_diagnostics.ScrapeJobs()
		case "gravity_jack":
			gravity_jack.ScrapeJobs()
		case "huntwood":
			huntwood.ScrapeJobs()
		case "intellitect":
			intellitect.ScrapeJobs()
		case "itron":
			itron.ScrapeJobs()
		case "kaiser_aluminum":
			kaiser_aluminum.ScrapeJobs()
		case "limelyte":
			limelyte.ScrapeJobs()
		case "numerica":
			numerica.ScrapeJobs()
		case "openeye":
			openeye.ScrapeJobs()
		case "paytrace":
			paytrace.ScrapeJobs()
		case "providence":
			providence.ScrapeJobs()
		case "risklens":
			risklens.ScrapeJobs()
		case "scld":
			scld.ScrapeJobs()
		case "sel":
			sel.ScrapeJobs()
		case "spokane_computer":
			spokane_computer.ScrapeJobs()
		case "spokane_library":
			spokane_library.ScrapeJobs()
		case "sps":
			sps.ScrapeJobs()
		case "stcu":
			stcu.ScrapeJobs()
		case "synergisticit":
			synergisticit.ScrapeJobs()
		case "treasury4":
			treasury4.ScrapeJobs()
		case "two_barrels":
			two_barrels.ScrapeJobs()
		case "urm":
			urm.ScrapeJobs()
		case "wagstaff":
			wagstaff.ScrapeJobs()
		case "winco":
			winco.ScrapeJobs()
		default:
			//log.Errorf("Scraper %s not found\n", name)
			log.Printf("Scraper %s not found\n", name)
	}
}
