package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"./urlshort"
)

func loadFileData(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	var handler http.Handler

	mux := defaultMux()

	filePath := flag.String("yamlFile", "urls.yaml", "the path to the file to shorten")
	flag.Parse()

	ext := filepath.Ext(*filePath)
	if ext == ".yaml" {
		yml, err := loadFileData(*filePath)
		if err != nil {
			panic(err)
		}

		handler, err = urlshort.YAMLHandler(yml, mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".json" {
		json, err := loadFileData(*filePath)
		if err != nil {
			panic(err)
		}
		handler, err = urlshort.JSONHandler(json, mux)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
