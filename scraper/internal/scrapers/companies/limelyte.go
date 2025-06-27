package companies

import (
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"github.com/gocolly/colly"
)

var limelyteScrapeUrl string = "https://apply.workable.com/limelyte-technology-group/"

// LimelyteScraper implements a scraper for Limelyte
type LimelyteScraper struct {
	*types.BaseScraper
}

func NewLimelyteScraper() *LimelyteScraper {
	return &LimelyteScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "limelyte",
			ScheduleHour: 21,
			Config: types.ScraperConfig{
				BaseURL: "https://limelyte.com/careers",
			},
		},
	}
}

func (l *LimelyteScraper) GetName() string {
	return l.Name
}

func (l *LimelyteScraper) ScrapedJobs() []types.ScrapedJob {
	if len(l.Jobs) == 0 {
		l.Jobs = l.scrapeJobs()
	}
	return l.Jobs
}

func (l *LimelyteScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (l *LimelyteScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	collector := l.getCollector()

	// Process Job Line - TODO: Add proper selectors
	collector.OnHTML("", func(h *colly.HTMLElement) {
		// todo: implement job extraction
	})

	log.Printf("\n-----\n\nColly instance created: %+v\n\n", collector)
	collector.Visit(limelyteScrapeUrl)
	return jobs
}

func (l *LimelyteScraper) getCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("apply.workable.com"),
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

func (l *LimelyteScraper) trimSpaces(s string) string {
	return strings.TrimSpace(s)
}
