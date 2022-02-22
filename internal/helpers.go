package app

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func validateURL(input string) error {
	_, err := url.ParseRequestURI(input)
	return err
}

func reciprocalScrape(url string, inDomain *[]string, outDomain *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	hrefs, err := getHrefsFromUrl(url)
	if err != nil {
		log.Fatalf("Error within goroutine: %s", err.Error())
	}

	for _, href := range hrefs {
		if href == "/" || href == "#" {
			continue
		} else if href[0:1] != "/" {
			*outDomain = append(*outDomain, href)
		} else {
			*inDomain = append(*inDomain, href)
		}
	}
}

func (l *HrefList) addToHrefList(href string) bool {
	for _, x := range *l {
		if x == href {
			return false
		}
	}
	*l = append(*l, href)
	return true
}

func getHrefsFromUrl(url string) ([]string, error) {
	output := []string{}

	res, err := http.Get(url)
	if err != nil {
		return output, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return output, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return output, err
	}

	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			if href, ok := getHrefFromNode(node); ok {
				output = append(output, href)
			}
		}
	})

	return output, nil

}

func getHrefFromNode(node *html.Node) (string, bool) {
	for _, v := range node.Attr {
		if v.Key == "href" {
			return v.Val, true
		}
	}
	return "", false
}
