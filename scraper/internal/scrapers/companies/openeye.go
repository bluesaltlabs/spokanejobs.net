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
		if err := utils.SaveJobsToJSON(jobs, o.Name); err != nil {
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
		// Keep original description for backward compatibility
		j.Description = e.Text()

		// Extract rich text description with formatting
		j.RichDescription = extractRichDescription(e)

		// Extract posting date
		j.PostingDate = extractPostingDate(e)

		// Extract salary range
		j.SalaryRange = extractSalaryRange(e)
	})

	if err := detailEngine.Visit(j.Url); err != nil {
		log.Printf("Error visiting job details %s: %v", j.Url, err)
	}

	// Close the detail engine
	detailEngine.Close()
}

// extractRichDescription extracts formatted content from the job description
func extractRichDescription(element engine.Element) string {
	var content strings.Builder
	text := element.Text()

	// Split into lines and process each line
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Check if line looks like a header
		if isHeader(line) {
			content.WriteString("\n\n")
			content.WriteString(line)
			content.WriteString("\n")
		} else if isListItem(line) {
			content.WriteString("• ")
			content.WriteString(line)
			content.WriteString("\n")
		} else {
			content.WriteString(line)
			content.WriteString("\n")
		}
	}

	return strings.TrimSpace(content.String())
}

// extractPostingDate looks for posting date information
func extractPostingDate(element engine.Element) string {
	// Look for common date selectors
	dateSelectors := []string{
		".posting-date",
		".date-posted",
		".job-date",
		"[data-date]",
		".timestamp",
		".published-date",
	}

	for _, selector := range dateSelectors {
		dateElement := element.Find(selector)
		if dateElement != nil {
			dateText := utils.TrimSpaces(dateElement.Text())
			if dateText != "" {
				return dateText
			}
		}
	}

	return ""
}

// extractSalaryRange looks for salary information
func extractSalaryRange(element engine.Element) string {
	salarySelectors := []string{
		".salary",
		".salary-range",
		".compensation",
		".pay-range",
		"[data-salary]",
		".wage",
	}

	for _, selector := range salarySelectors {
		salaryElement := element.Find(selector)
		if salaryElement != nil {
			salaryText := utils.TrimSpaces(salaryElement.Text())
			if salaryText != "" {
				return salaryText
			}
		}
	}

	// Also look for salary patterns in text
	text := element.Text()
	salaryPatterns := []string{
		"$", "salary", "compensation", "pay", "wage", "hourly", "annual",
	}

	for _, pattern := range salaryPatterns {
		if strings.Contains(strings.ToLower(text), pattern) {
			// Extract the sentence containing salary info
			sentences := strings.Split(text, ".")
			for _, sentence := range sentences {
				if strings.Contains(strings.ToLower(sentence), pattern) {
					return utils.TrimSpaces(sentence) + "."
				}
			}
		}
	}

	return ""
}

// Helper functions
func isHeader(text string) bool {
	// Check if text looks like a header
	return len(text) < 100 && (strings.ToUpper(text) == text ||
		strings.Contains(strings.ToLower(text), "responsibilities") ||
		strings.Contains(strings.ToLower(text), "requirements") ||
		strings.Contains(strings.ToLower(text), "qualifications") ||
		strings.Contains(strings.ToLower(text), "benefits"))
}

func isListItem(text string) bool {
	// Check if text looks like a list item
	return strings.HasPrefix(strings.TrimSpace(text), "•") ||
		strings.HasPrefix(strings.TrimSpace(text), "-") ||
		strings.HasPrefix(strings.TrimSpace(text), "*")
}
