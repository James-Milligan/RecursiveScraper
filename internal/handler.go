package internal

import (
	"fmt"
	"sync"
)

func Scrape(url string) (HrefList, error) {
	var wg sync.WaitGroup

	cycleOutputIn := HrefList{}
	masterOutputIn := HrefList{}
	toCycle := HrefList{}

	wg.Add(1)
	if err := reciprocalScrape(url, &cycleOutputIn, &wg); err != nil {
		panic(err)
	}
	wg.Wait()

	masterOutputIn.addNewToMaster(cycleOutputIn, &toCycle)

	for true {
		fmt.Println("Cycle...")
		for _, x := range toCycle {
			wg.Add(1)
			go reciprocalScrape(fmt.Sprintf("%s%s", url, x), &cycleOutputIn, &wg)
		}
		wg.Wait()
		toCycle = HrefList{}
		masterOutputIn.addNewToMaster(cycleOutputIn, &toCycle)
		if len(toCycle) == 0 {
			break
		}
	}

	return masterOutputIn, nil
}

func (l *HrefList) addNewToMaster(cycleOutput HrefList, nextCycle *HrefList) {
	for _, x := range cycleOutput {
		in := false
		for _, y := range *l {
			if x == y {
				in = true
				break
			}
		}
		if !in {
			*l = append(*l, x)
			*nextCycle = append(*nextCycle, x)
		}
	}
}
