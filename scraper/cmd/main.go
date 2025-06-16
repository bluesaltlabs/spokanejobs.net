package main

import (
	"log"
	"os"
  "gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

func main() {
	if len(os.Args) == 1 {
	  scrapers.RunAll()
	} else if len(os.Args) == 2 {
		scrapers.RunScraper(os.Args[1])
		// todo: error checking.
	} else {
		log.Println(len(os.Args), os.Args) // todo: use Printf instead
	}

}
