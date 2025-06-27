package engine

import (
	"context"
	"log"
	"strings"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// RodEngine wraps rod.Browser to implement ScraperEngine interface
type RodEngine struct {
	browser *rod.Browser
	page    *rod.Page
	config  types.ScraperConfig
	// Store handlers for later execution
	htmlHandlers    map[string][]func(Element)
	scrapedHandlers []func()
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
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1",
	})

	return &RodEngine{
		browser:         browser,
		page:            page,
		config:          config,
		htmlHandlers:    make(map[string][]func(Element)),
		scrapedHandlers: make([]func(), 0),
	}, nil
}

// Visit implements ScraperEngine.Visit
func (r *RodEngine) Visit(url string) error {
	log.Printf("visiting: %s\n", url)
	err := r.page.Navigate(url)
	if err != nil {
		return err
	}

	// Wait for page to load
	r.page.MustWaitLoad()

	// Execute stored HTML handlers
	for selector, handlers := range r.htmlHandlers {
		elements, err := r.page.Elements(selector)
		if err != nil {
			log.Printf("Error finding elements with selector %s: %v", selector, err)
			continue
		}

		for _, element := range elements {
			rodElement := NewRodElement(element)
			for _, handler := range handlers {
				handler(rodElement)
			}
		}
	}

	// Execute scraped handlers
	for _, handler := range r.scrapedHandlers {
		handler()
	}

	return nil
}

// OnHTML implements ScraperEngine.OnHTML
func (r *RodEngine) OnHTML(selector string, handler func(Element)) error {
	if r.htmlHandlers[selector] == nil {
		r.htmlHandlers[selector] = make([]func(Element), 0)
	}
	r.htmlHandlers[selector] = append(r.htmlHandlers[selector], handler)
	return nil
}

// OnScraped implements ScraperEngine.OnScraped
func (r *RodEngine) OnScraped(handler func()) error {
	r.scrapedHandlers = append(r.scrapedHandlers, handler)
	return nil
}

// OnRequest implements ScraperEngine.OnRequest
func (r *RodEngine) OnRequest(handler func(Request)) error {
	// Rod doesn't have request events like colly
	log.Printf("OnRequest called")
	return nil
}

// OnError implements ScraperEngine.OnError
func (r *RodEngine) OnError(handler func(Response, error)) error {
	// Rod doesn't have error events like colly
	log.Printf("OnError called")
	return nil
}

// SetHeaders implements ScraperEngine.SetHeaders
func (r *RodEngine) SetHeaders(headers map[string]string) {
	// Convert headers to extra headers format
	headerList := make([]string, 0, len(headers)*2)
	for name, value := range headers {
		headerList = append(headerList, name, value)
	}

	if len(headerList) > 0 {
		r.page.MustSetExtraHeaders(headerList...)
	}
}

// SetUserAgent implements ScraperEngine.SetUserAgent
func (r *RodEngine) SetUserAgent(userAgent string) {
	r.page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: userAgent,
	})
}

// SetAllowedDomains implements ScraperEngine.SetAllowedDomains
func (r *RodEngine) SetAllowedDomains(domains []string) {
	// Rod doesn't have domain restrictions like colly
	log.Printf("SetAllowedDomains called with: %v", domains)
}

// SetCacheDir implements ScraperEngine.SetCacheDir
func (r *RodEngine) SetCacheDir(dir string) {
	// Rod doesn't have cache like colly
	log.Printf("SetCacheDir called with: %s", dir)
}

// Close implements ScraperEngine.Close
func (r *RodEngine) Close() error {
	if r.page != nil {
		r.page.Close()
	}
	if r.browser != nil {
		r.browser.Close()
	}
	return nil
}

// GetPage returns the underlying rod.Page for direct access
func (r *RodEngine) GetPage() *rod.Page {
	return r.page
}

// GetBrowser returns the underlying rod.Browser for direct access
func (r *RodEngine) GetBrowser() *rod.Browser {
	return r.browser
}

// RodElement wraps rod.Element to implement Element interface
type RodElement struct {
	element *rod.Element
}

// NewRodElement creates a new RodElement
func NewRodElement(element *rod.Element) *RodElement {
	return &RodElement{element: element}
}

// Attr implements Element.Attr
func (e *RodElement) Attr(name string) string {
	attr, err := e.element.Attribute(name)
	if err != nil || attr == nil {
		return ""
	}
	return *attr
}

// Text implements Element.Text
func (e *RodElement) Text() string {
	text, err := e.element.Text()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(text)
}

// Find implements Element.Find
func (e *RodElement) Find(selector string) Element {
	element, err := e.element.Element(selector)
	if err != nil {
		return nil
	}
	return NewRodElement(element)
}

// DOM implements Element.DOM
func (e *RodElement) DOM() interface{} {
	return e.element
}

// ForEach implements Element.ForEach
func (e *RodElement) ForEach(selector string, fn func(int, Element)) error {
	elements, err := e.element.Elements(selector)
	if err != nil {
		return err
	}
	for i, element := range elements {
		fn(i, NewRodElement(element))
	}
	return nil
}

// RodRequest wraps rod.Request to implement Request interface
type RodRequest struct {
	request interface{} // Placeholder since rod doesn't expose Request type
}

// URL implements Request.URL
func (r *RodRequest) URL() string {
	return ""
}

// SetHeader implements Request.SetHeader
func (r *RodRequest) SetHeader(name, value string) {
	// Rod doesn't support setting headers on individual requests
	log.Printf("SetHeader called: %s = %s", name, value)
}

// Context implements Request.Context
func (r *RodRequest) Context() context.Context {
	return context.Background()
}

// RodResponse wraps rod.Response to implement Response interface
type RodResponse struct {
	response interface{} // Placeholder since rod doesn't expose Response type
}

// StatusCode implements Response.StatusCode
func (r *RodResponse) StatusCode() int {
	return 200 // Placeholder
}

// Body implements Response.Body
func (r *RodResponse) Body() []byte {
	return nil // Placeholder
}

// URL implements Response.URL
func (r *RodResponse) URL() string {
	return "" // Placeholder
}
