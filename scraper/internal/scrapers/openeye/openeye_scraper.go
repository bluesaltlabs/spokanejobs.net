package openeye_scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Main() {

	scrapeUrl := "https://job-boards.greenhouse.io/openeye"

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("job-boards.greenhouse.io"))


  // A simple check to prove the library was imported correctly
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Accept-Language", "en-US")
    r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1")
    fmt.Printf("Colly is visiting: %s\n", r.URL)
  })


  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    fmt.Printf("Error while scraping: %s\n", e.Error())
  })


  // Example selector
  c.OnHTML("div.job-posts--table tr.job-post", func(h *colly.HTMLElement) {
    selection := h.DOM
    url, _ := selection.Find("td.cell a").Attr("href")
    title := selection.Find("p.body.body--medium").Text() // Job Title
    location := selection.Find("p.body.body__secondary.body--metadata").Text() // Location (City, State)

    fmt.Printf("%s (%s) | %s\n", title, location, url)
  })


  fmt.Printf("\n-----\n\nColly instance created: %+v\n\n", c)

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)


  fmt.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
}
