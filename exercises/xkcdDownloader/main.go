package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func downloadImage(index int, element *goquery.Selection) {
	href, exists := element.Attr("src")
	if exists {
		saveImage(href)
	}
}

func saveImage(path string) {
	imgName := filepath.Base(path)
	fmt.Printf("downloading image:%v\n", imgName)
	image, err := http.Get("https:" + path)
	if err != nil {
		log.Fatal("error downloading the image")
	}
	file, err := os.Create(imgName)
	if err != nil {
		log.Fatal("error creating image file name")
	}
	fmt.Printf("downloading image:%v\n", imgName)

	_, err = io.Copy(file, image.Body)
	if err != nil {
		log.Fatal("error writing image to disk")
	}
	fmt.Println("image written to disk successfully")
}

func extractPrevLink(document *goquery.Document) string {
	href := ""
	selector := "a[rel=prev]"
	document.Find(selector).Each(func(index int, element *goquery.Selection) {
		attr, exist := element.Attr("href")
		if exist {
			href = attr
		}
	})
	return href
}

func initDownload(document *goquery.Document) {
	imgSelector := "#comic > img"
	document.Find(imgSelector).Each(downloadImage)
}

func main() {
	url := "https://xkcd.com"
	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("there was an error getting the html page")
		}
		defer resp.Body.Close()
		document, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal("there was an error passing the page")
		}
		initDownload(document)
		prevLink := extractPrevLink(document)
		url = "https://xkcd.com" + prevLink
		fmt.Println(url)
	}
	
	
}