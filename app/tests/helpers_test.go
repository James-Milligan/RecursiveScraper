package tests

import (
	"fmt"
	"github.com/James-Milligan/RecursiveScraper/app"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"sync"
	"testing"
)

func Test_GetHrefsFromDocument(t *testing.T) {
	f, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	output := app.GetHrefsFromDocument(*doc)

	if len(output) != 1 || output[0] != "href-1.com" {
		fmt.Println(output)
		t.Errorf("Unexpected response from GetHrefsFromDocument on test.html document")
	}
}

func Test_StepThroughURL(t *testing.T) {
	in := []string{}
	out := []string{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go app.StepThroughURL("https://jamesmilligan.net", &in, &out, &wg)
	wg.Wait()

	inCheck := []string{
		"https://jamesmilligan.net/websiteIcon.png",
		"https://jamesmilligan.net/manifest.json",
		"https://jamesmilligan.net/static/css/main.20b290b3.chunk.css",
	}
	outCheck := []string{
		"https://fonts.gstatic.com",
		"https://fonts.googleapis.com/css2?family=Sarala:wght@400;700\u0026family=Shanti\u0026display=swap",
		"https://fonts.googleapis.com",
		"https://fonts.googleapis.com/css2?family=Domine:wght@400;500;600;700\u0026family=Inter:wght@100;200;300;400;500;600;700;800;900\u0026display=swap",
	}

	inTest := app.HrefList{}
	for _, href := range in {
		inTest.AddToHrefList(href)
	}
	if len(inCheck) != len(inTest) {
		t.Errorf("Incorrect number of urls found for in-domain pages. Found %d want %d", len(inTest), len(inCheck))
	}
	for _, x := range inTest {
		found := false
		for _, y := range inCheck {
			if y == fmt.Sprintf("https://jamesmilligan.net%s", x) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Url been found by the scraper that has not been provided in the inCheck array %s", x)
		}
	}

	outTest := app.HrefList{}
	for _, href := range out {
		outTest.AddToHrefList(href)
	}
	if len(outCheck) != len(outTest) {
		t.Errorf("Incorrect number of urls found for in-domain pages. Found %d want %d", len(out), len(outTest))
	}
	for _, x := range outTest {
		found := false
		for _, y := range outCheck {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Url been found by the scraper that has not been provided in the outCheck array %s", x)
		}
	}
}
