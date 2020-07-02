package downloader

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Downloader struct {
	query    string
	filePath string
	hostname string
}

func (d *Downloader) GenerateHash() string {

	f, err := os.Open(d.filePath)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error reading the file")
	}
	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal("There was an error processing file")
	}
	hasBytes := h.Sum(nil)
	hash := hex.EncodeToString(hasBytes[:])
	return hash
}

func NewDownloader(path string) *Downloader {
	return &Downloader{query: "", filePath: path, hostname: "http://api.thesubdb.com"}
}

func (d *Downloader) GetLanguages() string {
	url := d.hostname + "/?action=languages"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request")
	}
	req.Header.Set("user-agent", "SubDB/1.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error processing request")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Couldn't read response data")
	}
	return string(content)
}

func (d *Downloader) DownloadSubtitle() {
	lang := "en"
	hash := d.GenerateHash()
	url := d.hostname + "/?action=download&hash=" + hash + "&language=" + lang
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error setting headers for downloading")
	}
	req.Header.Set("user-agent", "SubDB/1.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error processing request")
	}
	defer resp.Body.Close()

	fileName := filepath.Base(d.filePath)
	strings.TrimSuffix(fileName, filepath.Ext(fileName))
	subtitleFile := fileName + ".srt"

	f, err := os.Create(subtitleFile)
	defer f.Close()
	if err != nil {
		log.Fatal("error creating file")
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatal("Error downloading subtitle")
	}

}
