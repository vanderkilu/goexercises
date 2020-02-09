package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
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

type Story map[string] StoryArc

func loadStory() []byte  {
	file, err := ioutil.ReadFile("cyao/story.json")
	if err != nil {
		fmt.Print(err)
		log.Fatal("there was a problem reading the file, make sure it exists")
	}
	return file
}

func parseStory(storyMap []byte) (story Story) {
	err := json.Unmarshal(storyMap, &story)
	if err != nil {
		log.Fatal("error parsing story json file")
	}
	return
}

func makeTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

func templateHandler(key string, story Story, w http.ResponseWriter) error {

	t := makeTemplate("cyao/cyao.html")
	template, err := t.Clone()
	if err != nil {
		log.Fatal("error reading the template file")
	}
	storyContent := story[key] 
	return template.Execute(w, storyContent)
}

func storyHandler(story Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSpace(r.URL.Path)
		if path == "/" {
			path = "/intro"
		}
		key := path[1:]
		if err := templateHandler(key, story, w); err != nil {
			log.Fatal("couldn't read story with given arc path")
		}

	}
}

func main() {
	file := loadStory()
	storyMap := parseStory(file)

	mux := http.NewServeMux()
	story := storyHandler(storyMap)
	mux.HandleFunc("/", story)

	fmt.Println("listening and serving on port 3000..........")
	http.ListenAndServe(":3000", mux)
	
}