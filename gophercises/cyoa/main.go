package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func loadJSONData(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Story map[string]StoryArc

func (story Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./adventure.html"))
	t, err := tmpl.Clone()
	if err != nil {
		log.Fatal(err)
	}
	arc := strings.TrimSpace(r.URL.Path)
	if arc == "/" {
		arc = "/intro"
	}
	if strings.HasSuffix(arc, "/") {
		arc = strings.TrimSuffix(arc, "/")
	}
	data := story[arc[1:]]
	t.Execute(w, data)

}

func main() {
	data, err := loadJSONData("./adventure.json")
	if err != nil {
		log.Fatal(err)
	}
	var story Story

	err = json.Unmarshal(data, &story)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", story)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
