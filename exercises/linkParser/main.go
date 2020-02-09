package main

import (
	"io"
	"golang.org/x/net/html"
	"log"
	"fmt"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func LinkParser(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal("there was an error parsing the html")
	}
	links := buildLinks(doc)
	return links
}

func buildLinks(n *html.Node) []Link {
	var links []Link
	if (n.Type == html.ElementNode && n.Data == "a") {
		for _, attr := range n.Attr {
			if (attr.Key == "href") {
				text := extractText(n)
				links = append(links, Link{Href: attr.Val, Text: text})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, buildLinks(c)...)
	}
	return links
}

func extractText(n *html.Node) string {
	if (n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode) {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

var htmlTemplate string = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</section>
`

func main() {
	links := LinkParser(strings.NewReader(htmlTemplate))
	fmt.Println(links)	
}