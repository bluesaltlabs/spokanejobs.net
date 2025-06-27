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

type SpsScraper struct {
	*types.BaseScraper
}

func NewSpsScraper() *SpsScraper {
	return &SpsScraper{
		BaseScraper: &types.BaseScraper{
			Name:         "sps",
			ScheduleHour: 23,
			Config: types.ScraperConfig{
				BaseURL: "https://www.spokaneschools.org/careers/",
			},
		},
	}
}

func (scraper *SpsScraper) GetName() string {
	return scraper.Name
}

func (scraper *SpsScraper) ScrapedJobs() []types.ScrapedJob {
	if len(scraper.Jobs) == 0 {
		scraper.Jobs = scraper.scrapeJobs()
	}
	return scraper.Jobs
}

func (scraper *SpsScraper) ScrapeJobDetails(job *types.ScrapedJob) {
	// Default implementation - can be overridden if needed
}

func (scraper *SpsScraper) scrapeJobs() []types.ScrapedJob {
	var jobs []types.ScrapedJob
	c := utils.NewCollector(scraper.Config)

	c.OnScraped(func(r *colly.Response) {
		// Save to file instead of just printing to stdout
		if err := utils.SaveJobsToJSON(jobs, scraper.Name); err != nil {
			log.Printf("Error saving jobs to JSON for %s: %v", scraper.Name, err)
		} else {
			log.Printf("Saved %d jobs to JSON file for %s", len(jobs), scraper.Name)
		}

		// Also print to stdout for backward compatibility
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(jobs)
	})

	c.OnHTML(".job-listing", func(h *colly.HTMLElement) {
		title := strings.TrimSpace(h.DOM.Find(".job-title").Text())
		location := strings.TrimSpace(h.DOM.Find(".job-location").Text())
		link, _ := h.DOM.Find("a").Attr("href")

		if title != "" && link != "" {
			// Make link absolute if it's relative
			if !strings.HasPrefix(link, "http") {
				link = scraper.Config.BaseURL + strings.TrimPrefix(link, "/")
			}

			job := types.ScrapedJob{
				Title: title,
				Url:   link,
				City:  location,
			}

			// Extract enhanced content from job details page
			scraper.extractJobDetails(&job)

			jobs = append(jobs, job)
		}
	})

	c.Visit(scraper.Config.BaseURL)
	return jobs
}

// extractJobDetails extracts rich content from individual job pages
func (scraper *SpsScraper) extractJobDetails(job *types.ScrapedJob) {
	detailCollector := utils.NewCollector(scraper.Config)

	detailCollector.OnHTML(".job-description", func(h *colly.HTMLElement) {
		// Keep original description for backward compatibility
		job.Description = h.Text

		// Extract rich text description with formatting
		job.RichDescription = scraper.extractRichDescription(h)

		// Extract posting date
		job.PostingDate = scraper.extractPostingDate(h)

		// Extract salary range
		job.SalaryRange = scraper.extractSalaryRange(h)
	})

	detailCollector.Visit(job.Url)
}

// extractRichDescription extracts formatted content from the job description
func (scraper *SpsScraper) extractRichDescription(h *colly.HTMLElement) string {
	var content strings.Builder
	text := h.Text

	// Split into lines and process each line
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Check if line looks like a header
		if scraper.isHeader(line) {
			content.WriteString("\n\n")
			content.WriteString(line)
			content.WriteString("\n")
		} else if scraper.isListItem(line) {
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
func (scraper *SpsScraper) extractPostingDate(h *colly.HTMLElement) string {
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
		dateText := strings.TrimSpace(h.DOM.Find(selector).Text())
		if dateText != "" {
			return dateText
		}
	}

	return ""
}

// extractSalaryRange looks for salary information
func (scraper *SpsScraper) extractSalaryRange(h *colly.HTMLElement) string {
	salarySelectors := []string{
		".salary",
		".salary-range",
		".compensation",
		".pay-range",
		"[data-salary]",
		".wage",
	}

	for _, selector := range salarySelectors {
		salaryText := strings.TrimSpace(h.DOM.Find(selector).Text())
		if salaryText != "" {
			return salaryText
		}
	}

	// Also look for salary patterns in text
	text := h.Text
	salaryPatterns := []string{
		"$", "salary", "compensation", "pay", "wage", "hourly", "annual",
	}

	for _, pattern := range salaryPatterns {
		if strings.Contains(strings.ToLower(text), pattern) {
			// Extract the sentence containing salary info
			sentences := strings.Split(text, ".")
			for _, sentence := range sentences {
				if strings.Contains(strings.ToLower(sentence), pattern) {
					return strings.TrimSpace(sentence) + "."
				}
			}
		}
	}

	return ""
}

// Helper functions
func (scraper *SpsScraper) isHeader(text string) bool {
	// Check if text looks like a header
	return len(text) < 100 && (strings.ToUpper(text) == text ||
		strings.Contains(strings.ToLower(text), "responsibilities") ||
		strings.Contains(strings.ToLower(text), "requirements") ||
		strings.Contains(strings.ToLower(text), "qualifications") ||
		strings.Contains(strings.ToLower(text), "benefits"))
}

func (scraper *SpsScraper) isListItem(text string) bool {
	// Check if text looks like a list item
	return strings.HasPrefix(strings.TrimSpace(text), "•") ||
		strings.HasPrefix(strings.TrimSpace(text), "-") ||
		strings.HasPrefix(strings.TrimSpace(text), "*")
}
