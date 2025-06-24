package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitea.bluesaltlabs.com/BlueSaltLabs/bedrock/scraper/internal/scrapers"
)

func main() {
	fmt.Printf("=== Scraper Init ===\n")
    now := time.Now()
    fmt.Printf("Current time: %v\n", now)

	if len(os.Args) == 1 {
	  scrapers.RunAll()
	} else if len(os.Args) == 2 {
		scrapers.RunScraper(os.Args[1])
		// todo: error checking.
	} else {
		log.Println(len(os.Args), os.Args) // todo: use Printf instead
	}

}
