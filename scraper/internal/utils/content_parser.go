package utils

import (
	"strings"
)

// ExtractRichText extracts formatted content from an HTML element
func ExtractRichText(element interface{}) string {
	// For now, we'll use a simplified approach that works with the existing engine interface
	// This will need to be adapted based on the actual element interface

	// Process text content with basic formatting
	text := getElementText(element)
	if text == "" {
		return ""
	}

	// Split into lines and process each line
	lines := strings.Split(text, "\n")
	var content strings.Builder

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

// ExtractPostingDate looks for common date patterns in the job content
func ExtractPostingDate(element interface{}) string {
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
		dateElement := findElement(element, selector)
		if dateElement != nil {
			dateText := strings.TrimSpace(getElementText(dateElement))
			if dateText != "" {
				return dateText
			}
		}
	}

	return ""
}

// ExtractSalaryRange looks for salary information in the job content
func ExtractSalaryRange(element interface{}) string {
	salarySelectors := []string{
		".salary",
		".salary-range",
		".compensation",
		".pay-range",
		"[data-salary]",
		".wage",
	}

	for _, selector := range salarySelectors {
		salaryElement := findElement(element, selector)
		if salaryElement != nil {
			salaryText := strings.TrimSpace(getElementText(salaryElement))
			if salaryText != "" {
				return salaryText
			}
		}
	}

	// Also look for salary patterns in text
	text := getElementText(element)
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

// Helper functions for element interaction
func getElementText(element interface{}) string {
	// This would need to be implemented based on the actual element interface
	// For now, return empty string
	return ""
}

func findElement(element interface{}, selector string) interface{} {
	// This would need to be implemented based on the actual element interface
	// For now, return nil
	return nil
}

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
