package main

import (
	"flag"
	"io"
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"github.com/vanderkilu/goexercises/exercises/link"
)

func extractLinks(r io.Reader, base string) []string {
	var hrefs []string
	links := link.LinkParser(r)
	for _, link := range links {
		if strings.HasPrefix(link.Href,"/") {
			hrefs = append(hrefs, base + link.Href)
		}
		if strings.HasPrefix(link.Href, "http") {
			hrefs = append(hrefs, link.Href)
		}
		
	}
	return hrefs
}

func filterLinks(links []string, base string) []string {
	var filteredLinks []string
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			filteredLinks = append(filteredLinks, link)
		}
	}
	return filteredLinks
}

func get(inputUrl string) []string {
	resp, err := http.Get(inputUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	base := &url.URL{
		Host: resp.Request.URL.Host,
		Scheme: resp.Request.URL.Scheme,
	}
	links := extractLinks(resp.Body, base.String())
	links = filterLinks(links, base.String())
	return links
}

func getAllLinks(url string, maxDepth int) []string {
	cache := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		url: struct{}{},
	}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for k,_ := range q {
			if _, ok := cache[k]; ok {
				continue
			}
			cache[k] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	var links []string
	for k, _ := range cache {
		links = append(links,k)
	}
	return links
}

func main() {

	//1. Get all the links from a page
	//2. Go through each link and get all associated links recursively
	//3. Build a sitemap by encoding the result in an xml

	inputUrl := flag.String("url", "http://gophercises.com", "the url you want to parse")
	flag.Parse()

	links := getAllLinks(*inputUrl,5)

	for _, link := range links {
		fmt.Println(link)
	}

	
}