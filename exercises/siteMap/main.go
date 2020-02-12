package main

import (
	"flag"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func parseLinks(baseUrl string) []Link {
	links := LinkParser(strings.NewReader(baseUrl))
	return links
}

func main() {

	//1. Get all the links from a page
	//2. Go through each link and get all associated links recursively
	//3. Build a sitemap by encoding the result in an xml

	baseUrl := flag.String("url", "http://gophercises.com", "the url you want to parse")
	flag.Parse()
	parseLinks(*baseUrl)
}