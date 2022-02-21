package internal

func Scrape(url string) (HrefList, HrefList, error) {
	inDomain := HrefList{}
	outDomain := HrefList{}

	err := reciprocalScrape(url, &inDomain, &outDomain)

	return inDomain, outDomain, err
}
