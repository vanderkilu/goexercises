package main

import (
	"fmt"

	"./downloader"
)

func main() {
	filePath := "/home/vndrkl/Downloads/Video/dexter.mp4.mp4"
	subDownloader := downloader.NewDownloader(filePath)
	results := subDownloader.GetLanguages()
	subDownloader.DownloadSubtitle()
	fmt.Println(results)
}
