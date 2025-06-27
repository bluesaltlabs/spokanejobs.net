package companies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

var scrapeUrl string = "https://job-boards.greenhouse.io/openeye"
var jobUrlPrefix string = scrapeUrl + "/jobs/"

// OpenEyeScraper implements a scraper for OpenEye
// This is a drop-in replacement for the config-driven version, but with custom logic
// for job and detail extraction.
type OpenEyeScraper struct {
	*types.BaseScraper
}

func NewOpenEyeScraper() *OpenEyeScraper {
	return &OpenEyeScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "openeye",
			ScheduleHour: 15,
			Config: types.ScraperConfig{
				BaseURL: "https://openeye.net/careers",
			},
		},
	}
}

func (o *OpenEyeScraper) GetName() string {
	return o.Name
}

func (o *OpenEyeScraper) ScrapedJobs() []types.ScrapedJob {
	if len(o.Jobs) == 0 {
		o.Jobs = o.scrapeJobs()
	}
	return o.Jobs
}

func (o *OpenEyeScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (o *OpenEyeScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)
	c := getCollector()

	c.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	c.OnHTML("div.job-posts--table tr.job-post a", func(h *colly.HTMLElement) {
		j := &types.ScrapedJob{}
		url := h.Attr("href")
		selection := h.DOM

		job_id := strings.TrimPrefix(url, jobUrlPrefix)
		title := utils.TrimSpaces(selection.Find("p.body.body--medium").Text())

		j.JobId = job_id
		j.Url = url
		j.Title = title

		getJobDetails(j)
		jobs = append(jobs, *j)
	})

	c.Visit(scrapeUrl)
	log.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
	return jobs
}

func getJobDetails(j *types.ScrapedJob) {
	c := getCollector()

	c.OnHTML("h1.section-header", func(h *colly.HTMLElement) {
		if j.Title == "" || strings.EqualFold(j.Title, h.Text) {
			j.Title = h.Text
		}
	})

	c.OnHTML("div.job__location", func(h *colly.HTMLElement) {
		location := h.DOM.Text()
		j.City, j.State, _ = strings.Cut(location, ", ")
		j.City = utils.TrimSpaces(j.City)
		j.State = utils.TrimSpaces(j.State)
	})

	c.OnHTML("div.job__description", func(h *colly.HTMLElement) {
		j.Description = h.DOM.Text()
	})

	c.Visit(j.Url)
}

func getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("job-boards.greenhouse.io"),
		colly.CacheDir("./scraper_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return c
}
