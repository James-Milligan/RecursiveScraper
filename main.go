package main

import (
	"fmt"
	"github.com/James-Milligan/RecursiveScraper/internal"
)

func main() {
	url := "https://monzo.co.uk"
	fmt.Printf("Scraping %s\n----\n", url)
	in, out, err := internal.Scrape(url)
	if err != nil {
		panic(err)
	}
	for x, o := range out {
		fmt.Printf("External link %d: %s \n", x, o)
	}
	fmt.Println("----")
	for x, i := range in {
		fmt.Printf("Internal link %d: %s%s \n", x, url, i)
	}

}
