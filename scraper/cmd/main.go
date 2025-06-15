package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	scrapeUrl := "https://memos.bluesaltlabs.com"

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("memos.bluesaltlabs.com", "bluesaltlabs.com"))


  // A simple check to prove the library was imported correctly
  c.OnRequest(func(r *colly.Request) {
    fmt.Printf("Colly is visiting: %s", r.URL)
  })


  // ...
  c.OnHTML("div", func(h *colly.HTMLElement) {
 		fmt.Println(h.Text)
  })




  // Visit the memos.bluesaltlabs.com site.
  c.Visit(scrapeUrl)


  // Error handling
  c.OnError(func(r *colly.Response, e error) {
    fmt.Printf("Error while scraping: %s\n", e.Error())
  })

  fmt.Println("Go project setup is successful!")

  fmt.Printf("Colly instance created: %+v\n", c)
}
