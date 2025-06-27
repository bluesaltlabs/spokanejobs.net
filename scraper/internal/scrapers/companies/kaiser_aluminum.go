package companies

import (
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var kaiserAluminumScrapeUrl string = "https://myjobs.adp.com/kaisercareers/cx/job-listing?FIELD1=Washington&FIELD2=Spokane"

// KaiserAluminumScraper implements a scraper for Kaiser Aluminum
type KaiserAluminumScraper struct {
	*types.BaseScraper
}

func NewKaiserAluminumScraper() *KaiserAluminumScraper {
	return &KaiserAluminumScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "kaiser_aluminum",
			ScheduleHour: 12,
			Config: types.ScraperConfig{
				BaseURL: "https://www.kaiseraluminum.com/careers",
			},
		},
	}
}

func (k *KaiserAluminumScraper) GetName() string {
	return k.Name
}

func (k *KaiserAluminumScraper) ScrapedJobs() []types.ScrapedJob {
	if len(k.Jobs) == 0 {
		k.Jobs = k.scrapeJobs()
	}
	return k.Jobs
}

func (k *KaiserAluminumScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (k *KaiserAluminumScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := k.getCollector()

	collector.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, k.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", k.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), k.Name)
		}
		// Removed: debug print to stdout
	})

	// Process Job Line - TODO: Add proper selectors
	collector.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	log.Printf("\n-----\n\nColly instance created: %+v\n\n", collector)
	collector.Visit(kaiserAluminumScrapeUrl)
	return jobs
}

func (k *KaiserAluminumScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("myjobs.adp.com"),
		colly.CacheDir("./scraper_cache"),
	)

	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return collector
}

func (k *KaiserAluminumScraper) GetScheduleHour() int {
	return k.ScheduleHour
}

func (k *KaiserAluminumScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(k.Jobs, k.Name, outputDir)
}
