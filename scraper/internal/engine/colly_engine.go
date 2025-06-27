package engine

import (
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/gocolly/colly"
)

// CollyElement wraps colly.HTMLElement to implement the Element interface
type CollyElement struct {
	*colly.HTMLElement
}

func (c *CollyElement) Attr(name string) string {
	return c.HTMLElement.Attr(name)
}

func (c *CollyElement) Text() string {
	return c.HTMLElement.Text
}

func (c *CollyElement) Find(selector string) Element {
	element := c.HTMLElement.DOM.Find(selector)
	if element.Length() == 0 {
		return nil
	}
	return &CollyElement{HTMLElement: &colly.HTMLElement{DOM: element}}
}

// CollyEngine extends BaseEngine and provides direct access to colly.Collector
type CollyEngine struct {
	*BaseEngine
	Collector *colly.Collector
}

// NewCollyEngine creates a new Colly-based scraping engine
func NewCollyEngine(config types.ScraperConfig) *CollyEngine {
	// Use the utility function to create a standardized collector
	c := utils.NewCollector(config)

	return &CollyEngine{
		BaseEngine: NewBaseEngine(config),
		Collector:  c,
	}
}

// GetCollector returns the underlying colly.Collector for direct access
func (c *CollyEngine) GetCollector() *colly.Collector {
	return c.Collector
}

// OnScraped implements ScraperEngine interface
func (c *CollyEngine) OnScraped(callback func()) {
	c.Collector.OnScraped(func(r *colly.Response) {
		callback()
	})
}

// OnHTML implements ScraperEngine interface
func (c *CollyEngine) OnHTML(selector string, callback func(Element)) {
	c.Collector.OnHTML(selector, func(h *colly.HTMLElement) {
		callback(&CollyElement{HTMLElement: h})
	})
}

// Visit implements ScraperEngine interface
func (c *CollyEngine) Visit(url string) error {
	return c.Collector.Visit(url)
}

// Close implements ScraperEngine interface
func (c *CollyEngine) Close() {
	// Colly doesn't have a close method, but we can clean up if needed
}
