package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/boltdb/bolt"

	"./urlshort"
)

func loadFileData(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setUpDB(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("shortner"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("/urlshort"), []byte("https://github.com/gophercises/urlshort"))
		err = b.Put([]byte("/urlshort-final"), []byte("https://github.com/gophercises/urlshort/tree/solution"))
		return err
	})
}

func main() {
	var handler http.Handler

	mux := defaultMux()

	filePath := flag.String("yamlFile", "urls.yaml", "the path to the file to shorten")
	useDb := flag.Bool("useDB", false, "whether you will like to use a db or not")
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
	} else if ext == ".db" && *useDb {
		db, err := bolt.Open(*filePath, 0600, nil)
		defer db.Close()

		if err != nil {
			panic(err)
		}

		if err := setUpDB(db); err != nil {
			panic(err)
		}
		handler = urlshort.DBHandler(db, mux)

	} else {
		fmt.Println("Path file is invalid")
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
