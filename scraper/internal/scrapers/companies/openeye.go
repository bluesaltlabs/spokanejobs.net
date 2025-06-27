package companies

import (
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/engine"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
)

var scrapeUrl string = "https://job-boards.greenhouse.io/openeye"
var jobUrlPrefix string = scrapeUrl + "/jobs/"

// OpenEyeScraper implements a scraper for OpenEye
type OpenEyeScraper struct {
	*types.BaseScraper
}

func NewOpenEyeScraper() *OpenEyeScraper {
	config := types.ScraperConfig{
		BaseURL: "https://openeye.net/careers",
	}

	base := &types.BaseScraper{
		Name:         "openeye",
		ScheduleHour: 15,
		Config:       config,
	}

	// Create the engine
	factory := engine.NewEngineFactory()
	scraperEngine, err := factory.CreateEngine(engine.EngineColly, config)
	if err != nil {
		log.Printf("Error creating engine for openeye: %v", err)
		// Fallback to colly engine
		scraperEngine, _ = factory.CreateEngine(engine.EngineColly, config)
	}

	base.Engine = scraperEngine

	return &OpenEyeScraper{
		BaseScraper: base,
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

func (o *OpenEyeScraper) GetScheduleHour() int {
	return o.ScheduleHour
}

func (o *OpenEyeScraper) scrapeJobs() []types.ScrapedJob {
	jobs := make([]types.ScrapedJob, 0)

	// Get the engine from the base scraper
	scraperEngine, ok := o.Engine.(engine.ScraperEngine)
	if !ok {
		log.Printf("Error: Engine is not properly configured for %s", o.Name)
		return jobs
	}

	scraperEngine.OnScraped(func() {
		if err := utils.SaveJobsToJSON(jobs, o.Name, "scraper_output"); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", o.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), o.Name)
		}
	})

	scraperEngine.OnHTML("div.job-posts--table tr.job-post a", func(e engine.Element) {
		j := &types.ScrapedJob{}
		url := e.Attr("href")

		job_id := strings.TrimPrefix(url, jobUrlPrefix)
		titleElement := e.Find("p.body.body--medium")
		title := ""
		if titleElement != nil {
			title = utils.TrimSpaces(titleElement.Text())
		}

		j.JobId = job_id
		j.Url = url
		j.Title = title

		getJobDetails(j, scraperEngine)
		jobs = append(jobs, *j)
	})

	if err := scraperEngine.Visit(scrapeUrl); err != nil {
		log.Printf("Error visiting %s: %v", scrapeUrl, err)
	}

	log.Printf("\n-----\n\nScraping completed for %s\n\n", o.Name)
	return jobs
}

func getJobDetails(j *types.ScrapedJob, scraperEngine engine.ScraperEngine) {
	// Create a new engine for job details
	factory := engine.NewEngineFactory()
	config := types.ScraperConfig{
		BaseURL: "https://openeye.net/careers",
	}

	detailEngine, err := factory.CreateEngine(engine.EngineColly, config)
	if err != nil {
		log.Printf("Error creating detail engine: %v", err)
		return
	}

	detailEngine.OnHTML("h1.section-header", func(e engine.Element) {
		if j.Title == "" || strings.EqualFold(j.Title, e.Text()) {
			j.Title = e.Text()
		}
	})

	detailEngine.OnHTML("div.job__location", func(e engine.Element) {
		location := e.Text()
		j.City, j.State, _ = strings.Cut(location, ", ")
		j.City = utils.TrimSpaces(j.City)
		j.State = utils.TrimSpaces(j.State)
	})

	detailEngine.OnHTML("div.job__description", func(e engine.Element) {
		j.Description = e.Text()
	})

	if err := detailEngine.Visit(j.Url); err != nil {
		log.Printf("Error visiting job details %s: %v", j.Url, err)
	}

	// Close the detail engine
	detailEngine.Close()
}

func (o *OpenEyeScraper) SaveOutput(outputDir string) error {
	return utils.SaveJobsToJSON(o.Jobs, o.Name, outputDir)
}
