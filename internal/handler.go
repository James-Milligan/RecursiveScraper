package app

import (
	"fmt"
	"log"
	"sync"
)

type HrefList []string

type Output struct {
	Url           string   `json:"url"`
	Pages         HrefList `json:"pages"`
	ExternalPages HrefList `json:"-"`
}

func ScrapeURL(url string) Output {
	var wg sync.WaitGroup

	if err := validateURL(url); err != nil {
		log.Fatal(err)
	}

	nextCycle := []string{""}
	inDomainCycleOutput := []string{}
	outDomainCycleOutput := []string{}

	masterOutputIn := HrefList{}
	masterOutputOut := HrefList{}

	firstRun := true

	for len(nextCycle) != 0 || firstRun == true {

		if firstRun {
			firstRun = false
		}

		inDomainCycleOutput = []string{}
		outDomainCycleOutput = []string{}

		for _, x := range nextCycle {
			wg.Add(1)
			go reciprocalScrape(fmt.Sprintf("%s%s", url, x), &inDomainCycleOutput, &outDomainCycleOutput, &wg)
		}
		wg.Wait()

		nextCycle = []string{}
		for _, href := range inDomainCycleOutput {
			if ok := masterOutputIn.addToHrefList(href); ok {
				nextCycle = append(nextCycle, href)
			}
		}
		for _, href := range outDomainCycleOutput {
			masterOutputOut.addToHrefList(href)
		}

	}

	for i, x := range masterOutputIn {
		masterOutputIn[i] = fmt.Sprintf("%s%s", url, x)
	}

	return Output{
		Url:           url,
		Pages:         masterOutputIn,
		ExternalPages: masterOutputOut,
	}
}
