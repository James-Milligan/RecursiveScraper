package internal

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

type HrefList []string

func (l *HrefList) addToHrefList(href string) bool {
	for _, x := range *l {
		if x == href {
			return false
		}
	}
	*l = append(*l, href)
	return true
}

func reciprocalScrape(url string, inDomain *HrefList, outDomain *HrefList) error {
	hrefs, err := getHrefsFromUrl(url)
	if err != nil {
		return err
	}

	for _, href := range hrefs {
		if href == "/" || href == "#" {
			continue
		} else if href[0:1] != "/" {
			outDomain.addToHrefList(href)
		} else {
			if ok := inDomain.addToHrefList(href); ok {
				err := reciprocalScrape(fmt.Sprintf("%s%s", url, href), inDomain, outDomain)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
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

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
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
