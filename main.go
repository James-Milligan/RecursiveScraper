package main

import (
	"encoding/json"
	"flag"
	"fmt"
	app "github.com/James-Milligan/RecursiveScraper/app"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No arguments provided")
	}

	dataCmd := flag.NewFlagSet("scan", flag.ExitOnError)
	var url = dataCmd.String("u", "", "Required - URL path from which scrape will begin")
	var path = dataCmd.String("p", "output.json", "Optional - Output path for JSON. default is output.json")
	err := dataCmd.Parse(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if *url == "" {
		log.Fatal("no argument for URL")
	}
	if *path == "output.json" {
		fmt.Println("No output path provided, will save output to output.json")
	}

	results := app.ScrapeURL(*url)

	fmt.Printf("Scrape complete.\n%d pages found app to the domain\n%d pages found external to the domain\n---\n",
		len(results.Pages), len(results.ExternalPages))

	fmt.Println("Links to external pages:")
	for x, o := range results.ExternalPages {
		fmt.Printf("External link %d: %s \n", x+1, o)
	}

	fmt.Printf("Saving all pages within domain to %s\n", *path)
	file, _ := json.MarshalIndent(results, "", " ")

	if err = ioutil.WriteFile(*path, file, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")
}
