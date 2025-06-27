package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

// TrimSpaces returns the string with leading and trailing spaces removed.
func TrimSpaces(s string) string {
	return strings.TrimSpace(s)
}

// SaveJobsToJSON saves scraped jobs to a JSON file
func SaveJobsToJSON(jobs []types.ScrapedJob, scraperName, outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create filename with timestamp
	timestamp := time.Now().Format("yyyy-mm-dd HH:mm:ss")
	filename := fmt.Sprintf("%s_%s.json", scraperName, timestamp)
	filepath := filepath.Join(outputDir, filename)

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
