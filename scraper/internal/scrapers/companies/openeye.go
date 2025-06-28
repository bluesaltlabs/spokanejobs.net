package companies

import (
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/engine"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

const (
	baseURL      = "https://job-boards.greenhouse.io/openeye"
	scrapeUrl    = baseURL
	jobUrlPrefix = baseURL + "/jobs/"
)

// Register OpenEye scraper at init
func init() {
	scrapers.Register("OpenEye", "openeye", NewOpenEyeScraper)
}

type OpenEyeScraper struct {
	*scrapers.BaseScraper
}

func NewOpenEyeScraper() types.Scraper {
	config := types.ScraperConfig{
		BaseURL: baseURL,
	}

	engine, err := engine.NewEngine(engine.EngineColly, config)
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return nil
	}

	return &OpenEyeScraper{
		BaseScraper: &scrapers.BaseScraper{
			Name:         "OpenEye",
			Slug:         "openeye",
			ScheduleHour: 15,
			Engine:       engine,
		},
	}
}

func (o *OpenEyeScraper) ScrapedJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)

	if o.Engine == nil {
		log.Printf("Error: Engine is not properly configured for %s", o.Name)
		return jobs
	}

	o.Engine.OnHTML("div.job-posts--table tr.job-post a", func(e engine.Element) {
		url := e.Attr("href")
		jobId := strings.TrimPrefix(url, jobUrlPrefix)

		titleElement := e.Find("p.body.body--medium")
		title := ""
		if titleElement != nil {
			title = utils.TrimSpaces(titleElement.Text())
		}

		job := &types.ScrapedJob{
			JobId:   jobId,
			Url:     url,
			Title:   title,
			Company: o.Slug,
		}

		// Get job details using the same engine
		o.getJobDetails(job)
		jobs = append(jobs, *job)
	})

	if err := o.Engine.Visit(scrapeUrl); err != nil {
		log.Printf("Error visiting %s: %v", scrapeUrl, err)
	}

	log.Printf("Scraping completed for %s", o.Name)
	return jobs
}

func (o *OpenEyeScraper) getJobDetails(job *types.ScrapedJob) {
	// Use the same engine to visit the job details page
	// This is more efficient than creating a new engine for each job

	// Set up callbacks for job details
	o.Engine.OnHTML("h1.section-header", func(e engine.Element) {
		if job.Title == "" || strings.EqualFold(job.Title, e.Text()) {
			job.Title = e.Text()
		}
	})

	o.Engine.OnHTML("div.job__location", func(e engine.Element) {
		location := e.Text()
		job.City, job.State, _ = strings.Cut(location, ", ")
		job.City = utils.TrimSpaces(job.City)
		job.State = utils.TrimSpaces(job.State)
	})

	o.Engine.OnHTML("div.job__description", func(e engine.Element) {
		job.Description = e.Text()
	})

	if err := o.Engine.Visit(job.Url); err != nil {
		log.Printf("Error visiting job details %s: %v", job.Url, err)
	}
}
