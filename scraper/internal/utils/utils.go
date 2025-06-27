package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/git"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// TrimSpaces returns the string with leading and trailing spaces removed.
func TrimSpaces(s string) string {
	return strings.TrimSpace(s)
}

// SaveJobsToJSON saves scraped jobs to a JSON file
func SaveJobsToJSON(jobs []types.ScrapedJob, scraperName string) error {
	// Create companies directory within the output directory
	companiesDir := filepath.Join("scraper_output/companies")
	if err := os.MkdirAll(companiesDir, 0755); err != nil {
		return fmt.Errorf("failed to create companies directory: %w", err)
	}

	// Create filename without timestamp for consistency
	filename := fmt.Sprintf("%s.json", scraperName)
	filepath := filepath.Join(companiesDir, filename)

	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filepath, err)
	}
	defer file.Close()

	// Create JSON encoder with indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Encode the jobs
	if err := encoder.Encode(jobs); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

// ConsolidateJobsToJSON reads all company JSON files and combines them into a single jobs.json file
func ConsolidateJobsToJSON() error {
	companiesDir := filepath.Join("scraper_output", "companies")

	// Check if companies directory exists
	if _, err := os.Stat(companiesDir); os.IsNotExist(err) {
		return fmt.Errorf("companies directory does not exist: %s", companiesDir)
	}

	// Read all JSON files in the companies directory
	files, err := os.ReadDir(companiesDir)
	if err != nil {
		return fmt.Errorf("failed to read companies directory: %w", err)
	}

	allJobs := make([]types.ScrapedJob, 0)

	// Process each company file
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(companiesDir, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Warning: failed to read file %s: %v", filePath, err)
			continue
		}

		var companyJobs []types.ScrapedJob
		if err := json.Unmarshal(fileData, &companyJobs); err != nil {
			log.Printf("Warning: failed to parse JSON from %s: %v", filePath, err)
			continue
		}

		// Add company name to each job if not already present
		companyName := strings.TrimSuffix(file.Name(), ".json")
		for i := range companyJobs {
			if companyJobs[i].Company == "" {
				companyJobs[i].Company = companyName
			}
		}

		allJobs = append(allJobs, companyJobs...)
		log.Printf("Added %d jobs from %s", len(companyJobs), file.Name())
	}

	// Create the consolidated jobs.json file
	jobsFilePath := filepath.Join("scraper_output", "jobs.json")
	file, err := os.Create(jobsFilePath)
	if err != nil {
		return fmt.Errorf("failed to create jobs.json file: %w", err)
	}
	defer file.Close()

	// Create JSON encoder with indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Encode all jobs
	if err := encoder.Encode(allJobs); err != nil {
		return fmt.Errorf("failed to encode consolidated jobs JSON: %w", err)
	}

	log.Printf("Successfully consolidated %d total jobs into %s", len(allJobs), jobsFilePath)

	// Sync consolidated jobs.json to data repo if configured
	// Use hard-coded path for Docker container consistency
	dataRepoPath := "/repo_data"
	dataRepoSubdir := os.Getenv("DATA_REPO_SUBDIR")
	if dataRepoSubdir == "" {
		dataRepoSubdir = "api" // Default fallback
	}

	// Check if the data repo path exists (only sync if it does)
	if _, err := os.Stat(dataRepoPath); err == nil {
		// Initialize git sync for the consolidated file
		gs := git.NewGitSync(dataRepoPath, "consolidated", dataRepoSubdir)
		if err := gs.SyncJobs(allJobs); err != nil {
			log.Printf("[GitSync] Failed to sync consolidated jobs.json: %v", err)
		} else {
			log.Printf("[GitSync] Successfully synced consolidated jobs.json to data repo")
		}
	} else {
		log.Printf("[GitSync] Data repo path %s not found, skipping git sync", dataRepoPath)
	}

	return nil
}
