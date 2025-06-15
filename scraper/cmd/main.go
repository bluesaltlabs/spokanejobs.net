package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	scrapeUrl := "https://memos.bluesaltlabs.com/"

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("memos.bluesaltlabs.com", "bluesaltlabs.com"))


  // A simple check to prove the library was imported correctly
  c.OnRequest(func(r *colly.Request) {
    fmt.Printf("Colly is visiting: %s\n", r.URL)
  })


  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    fmt.Printf("Error while scraping: %s\n", e.Error())
  })



  // Example selector
  c.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Printf("%s\n", h.Text)
  })

  // Example selector 2
  c.OnHTML("div.w-full.grid div.group", func(h *colly.HTMLElement) {
 		fmt.Printf("%s\n", h.Text)
  })



  fmt.Printf("\n-----\n\nColly instance created: %+v\n\n", c)

  // Visit the scrapeUrl site (initate the script. )
  c.Visit(scrapeUrl)


  fmt.Printf("\n-----\n\nColly instance done: %+v\n\n", c)
}
