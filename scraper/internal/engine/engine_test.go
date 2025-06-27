package engine

import (
	"testing"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
)

func TestEngineFactory(t *testing.T) {
	factory := NewEngineFactory()

	config := types.ScraperConfig{
		AllowedDomains: []string{"example.com"},
		BaseURL:        "https://example.com",
	}

	// Test creating colly engine
	collyEngine, err := factory.CreateEngine(EngineColly, config)
	if err != nil {
		t.Fatalf("Failed to create colly engine: %v", err)
	}

	if collyEngine == nil {
		t.Fatal("Colly engine is nil")
	}

	// Test that it implements the interface
	var _ ScraperEngine = collyEngine

	// Test creating rod engine
	rodEngine, err := factory.CreateEngine(EngineRod, config)
	if err != nil {
		t.Fatalf("Failed to create rod engine: %v", err)
	}

	if rodEngine == nil {
		t.Fatal("Rod engine is nil")
	}

	// Test that it implements the interface
	var _ ScraperEngine = rodEngine

	// Test default engine (should be colly)
	defaultEngine, err := factory.CreateEngine("", config)
	if err != nil {
		t.Fatalf("Failed to create default engine: %v", err)
	}

	if defaultEngine == nil {
		t.Fatal("Default engine is nil")
	}

	// Test invalid engine type
	_, err = factory.CreateEngine("invalid", config)
	if err == nil {
		t.Fatal("Expected error for invalid engine type")
	}
}

func TestCollyEngine(t *testing.T) {
	config := types.ScraperConfig{
		AllowedDomains: []string{"example.com"},
		BaseURL:        "https://example.com",
	}

	engine := NewCollyEngine(config)
	if engine == nil {
		t.Fatal("Colly engine is nil")
	}

	// Test that it implements the interface
	var _ ScraperEngine = engine

	// Test basic functionality
	engine.SetUserAgent("New User Agent")
	engine.SetHeaders(map[string]string{"Test-Header": "test-value"})

	// Test OnHTML (should not panic)
	engine.OnHTML("test", func(e Element) {
		// This should not be called in this test
	})

	// Test OnScraped (should not panic)
	engine.OnScraped(func() {
		// This should not be called in this test
	})

	// Test Close
	err := engine.Close()
	if err != nil {
		t.Fatalf("Failed to close engine: %v", err)
	}
}
