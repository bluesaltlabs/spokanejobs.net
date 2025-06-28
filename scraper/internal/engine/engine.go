package engine

import (
	"fmt"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gocolly/colly"
)

// EngineType represents the type of scraping engine
type EngineType string

const (
	EngineColly EngineType = "colly"
	EngineRod   EngineType = "rod"
)

// Element represents a DOM element for scraping
type Element interface {
	Attr(name string) string
	Text() string
	Find(selector string) Element
}

// Engine provides a unified interface for both Colly and Rod engines
type Engine struct {
	Type    EngineType
	Config  types.ScraperConfig
	Colly   *colly.Collector
	Rod     *rod.Browser
	RodPage *rod.Page
}

// NewEngine creates a new engine of the specified type
func NewEngine(engineType EngineType, config types.ScraperConfig) (*Engine, error) {
	engine := &Engine{
		Type:   engineType,
		Config: config,
	}

	switch engineType {
	case EngineColly:
		engine.Colly = utils.NewCollector(config)
		return engine, nil
	case EngineRod:
		// Configure launcher with defaults
		launcherURL := launcher.New().
			Headless(true).
			MustLaunch()

		// Create browser
		browser := rod.New().ControlURL(launcherURL)
		if err := browser.Connect(); err != nil {
			return nil, err
		}

		// Create page
		page := browser.MustPage()

		// Set default user agent
		page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
			UserAgent: utils.DefaultUserAgent,
		})

		engine.Rod = browser
		engine.RodPage = page
		return engine, nil
	default:
		return nil, fmt.Errorf("unknown engine type: %s", engineType)
	}
}

// OnHTML registers a callback for HTML elements matching the selector
func (e *Engine) OnHTML(selector string, callback func(Element)) {
	switch e.Type {
	case EngineColly:
		e.Colly.OnHTML(selector, func(h *colly.HTMLElement) {
			callback(&CollyElement{HTMLElement: h})
		})
	case EngineRod:
		// For Rod, we'll process elements when Visit is called
		// This is a simplified approach - in practice you might want more sophisticated handling
	}
}

// Visit navigates to the specified URL
func (e *Engine) Visit(url string) error {
	switch e.Type {
	case EngineColly:
		return e.Colly.Visit(url)
	case EngineRod:
		return e.RodPage.Navigate(url)
	default:
		return fmt.Errorf("unknown engine type")
	}
}

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
