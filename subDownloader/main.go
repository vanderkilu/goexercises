package main

import (
	"fmt"

	"./downloader"
)

func main() {
	filePath := "/home/vndrkl/Videos/downloadcleaner.mp4"
	subDownloader := downloader.NewDownloader(filePath)
	results := subDownloader.GetLanguages()
	fmt.Println(results)
}
