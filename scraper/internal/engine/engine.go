package engine

import (
	"context"
)

// ScraperEngine defines the interface for different scraping engines
type ScraperEngine interface {
	Visit(url string) error
	OnHTML(selector string, handler func(Element)) error
	OnScraped(handler func()) error
	OnRequest(handler func(Request)) error
	OnError(handler func(Response, error)) error
	SetHeaders(headers map[string]string)
	SetUserAgent(userAgent string)
	SetAllowedDomains(domains []string)
	SetCacheDir(dir string)
	Close() error
}

// Element represents an HTML element that can be queried
type Element interface {
	Attr(name string) string
	Text() string
	Find(selector string) Element
	DOM() interface{} // Returns the underlying DOM element
	ForEach(selector string, fn func(int, Element)) error
}

// Request represents an HTTP request
type Request interface {
	URL() string
	SetHeader(name, value string)
	Context() context.Context
}

// Response represents an HTTP response
type Response interface {
	StatusCode() int
	Body() []byte
	URL() string
}

// EngineType represents the type of scraping engine
type EngineType string

const (
	EngineColly EngineType = "colly"
	EngineRod   EngineType = "rod"
)
