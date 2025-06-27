package engine

import (
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/utils"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// RodElement wraps rod.Element to implement the Element interface
type RodElement struct {
	element *rod.Element
}

func (r *RodElement) Attr(name string) string {
	attr, err := r.element.Attribute(name)
	if err != nil || attr == nil {
		return ""
	}
	return *attr
}

func (r *RodElement) Text() string {
	text, err := r.element.Text()
	if err != nil {
		return ""
	}
	return text
}

func (r *RodElement) Find(selector string) Element {
	element, err := r.element.Element(selector)
	if err != nil {
		return nil
	}
	return &RodElement{element: element}
}

// RodEngineWrapper wraps RodEngine to implement ScraperEngine interface
type RodEngineWrapper struct {
	Engine *RodEngine
}

func (r *RodEngineWrapper) OnScraped(callback func()) {
	// Rod doesn't have an OnScraped equivalent, so we'll call it immediately
	callback()
}

func (r *RodEngineWrapper) OnHTML(selector string, callback func(Element)) {
	// For Rod, we'll need to implement this differently
	// This is a simplified implementation
	elements, err := r.Engine.Page.Elements(selector)
	if err != nil {
		return
	}
	for _, element := range elements {
		callback(&RodElement{element: element})
	}
}

func (r *RodEngineWrapper) Visit(url string) error {
	return r.Engine.Page.Navigate(url)
}

func (r *RodEngineWrapper) Close() {
	r.Engine.Close()
}

// RodEngine extends BaseEngine and provides direct access to rod.Browser and rod.Page
type RodEngine struct {
	*BaseEngine
	Browser *rod.Browser
	Page    *rod.Page
}

// NewRodEngine creates a new Rod-based scraping engine
func NewRodEngine(config types.ScraperConfig) (*RodEngine, error) {
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

	return &RodEngine{
		BaseEngine: NewBaseEngine(config),
		Browser:    browser,
		Page:       page,
	}, nil
}

// GetBrowser returns the underlying rod.Browser for direct access
func (r *RodEngine) GetBrowser() *rod.Browser {
	return r.Browser
}

// GetPage returns the underlying rod.Page for direct access
func (r *RodEngine) GetPage() *rod.Page {
	return r.Page
}

// Close closes the browser and page
func (r *RodEngine) Close() error {
	if r.Page != nil {
		r.Page.Close()
	}
	if r.Browser != nil {
		r.Browser.Close()
	}
	return nil
}
