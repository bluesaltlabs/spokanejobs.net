package engine

import (
	"context"
	"log"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/types"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// CollyEngine wraps colly.Collector to implement ScraperEngine interface
type CollyEngine struct {
	collector *colly.Collector
	config    types.ScraperConfig
}

// NewCollyEngine creates a new Colly-based scraping engine
func NewCollyEngine(config types.ScraperConfig) *CollyEngine {
	// Set default cache directory
	cacheDir := "./scraper_cache"

	c := colly.NewCollector(
		colly.AllowedDomains(config.AllowedDomains...),
		colly.CacheDir(cacheDir),
	)

	// Set default headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
		log.Printf("visiting: %s\n", r.URL.String())
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Printf("Error while scraping: %s\n", e.Error())
	})

	return &CollyEngine{
		collector: c,
		config:    config,
	}
}

// Visit implements ScraperEngine.Visit
func (c *CollyEngine) Visit(url string) error {
	return c.collector.Visit(url)
}

// OnHTML implements ScraperEngine.OnHTML
func (c *CollyEngine) OnHTML(selector string, handler func(Element)) error {
	c.collector.OnHTML(selector, func(h *colly.HTMLElement) {
		handler(&CollyElement{HTMLElement: h})
	})
	return nil
}

// OnScraped implements ScraperEngine.OnScraped
func (c *CollyEngine) OnScraped(handler func()) error {
	c.collector.OnScraped(func(r *colly.Response) {
		handler()
	})
	return nil
}

// OnRequest implements ScraperEngine.OnRequest
func (c *CollyEngine) OnRequest(handler func(Request)) error {
	c.collector.OnRequest(func(r *colly.Request) {
		handler(&CollyRequest{Request: r})
	})
	return nil
}

// OnError implements ScraperEngine.OnError
func (c *CollyEngine) OnError(handler func(Response, error)) error {
	c.collector.OnError(func(r *colly.Response, e error) {
		handler(&CollyResponse{Response: r}, e)
	})
	return nil
}

// SetHeaders implements ScraperEngine.SetHeaders
func (c *CollyEngine) SetHeaders(headers map[string]string) {
	c.collector.OnRequest(func(r *colly.Request) {
		for name, value := range headers {
			r.Headers.Set(name, value)
		}
	})
}

// SetUserAgent implements ScraperEngine.SetUserAgent
func (c *CollyEngine) SetUserAgent(userAgent string) {
	c.collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})
}

// SetAllowedDomains implements ScraperEngine.SetAllowedDomains
func (c *CollyEngine) SetAllowedDomains(domains []string) {
	// Note: This would require recreating the collector, so we'll just log it
	log.Printf("SetAllowedDomains called with: %v", domains)
}

// SetCacheDir implements ScraperEngine.SetCacheDir
func (c *CollyEngine) SetCacheDir(dir string) {
	// Note: This would require recreating the collector, so we'll just log it
	log.Printf("SetCacheDir called with: %s", dir)
}

// Close implements ScraperEngine.Close
func (c *CollyEngine) Close() error {
	// Colly doesn't have a close method, so we just return nil
	return nil
}

// CollyElement wraps colly.HTMLElement to implement Element interface
type CollyElement struct {
	*colly.HTMLElement
}

// Attr implements Element.Attr
func (e *CollyElement) Attr(name string) string {
	return e.HTMLElement.Attr(name)
}

// Text implements Element.Text
func (e *CollyElement) Text() string {
	return e.HTMLElement.Text
}

// Find implements Element.Find
func (e *CollyElement) Find(selector string) Element {
	selection := e.HTMLElement.DOM.Find(selector)
	if selection.Length() > 0 {
		return &CollyElement{HTMLElement: &colly.HTMLElement{
			DOM: selection,
		}}
	}
	return nil
}

// DOM implements Element.DOM
func (e *CollyElement) DOM() interface{} {
	return e.HTMLElement.DOM
}

// ForEach implements Element.ForEach
func (e *CollyElement) ForEach(selector string, fn func(int, Element)) error {
	e.HTMLElement.DOM.Find(selector).Each(func(i int, s *goquery.Selection) {
		fn(i, &CollyElement{HTMLElement: &colly.HTMLElement{DOM: s}})
	})
	return nil
}

// CollyRequest wraps colly.Request to implement Request interface
type CollyRequest struct {
	*colly.Request
}

// URL implements Request.URL
func (r *CollyRequest) URL() string {
	return r.Request.URL.String()
}

// SetHeader implements Request.SetHeader
func (r *CollyRequest) SetHeader(name, value string) {
	r.Request.Headers.Set(name, value)
}

// Context implements Request.Context
func (r *CollyRequest) Context() context.Context {
	// Return a background context since colly.Context doesn't implement context.Context
	return context.Background()
}

// CollyResponse wraps colly.Response to implement Response interface
type CollyResponse struct {
	*colly.Response
}

// StatusCode implements Response.StatusCode
func (r *CollyResponse) StatusCode() int {
	return r.Response.StatusCode
}

// Body implements Response.Body
func (r *CollyResponse) Body() []byte {
	return r.Response.Body
}

// URL implements Response.URL
func (r *CollyResponse) URL() string {
	return r.Response.Request.URL.String()
}
