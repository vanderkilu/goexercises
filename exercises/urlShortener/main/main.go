package main

import (
	"fmt"
	"net/http"
	"log"
	yaml "gopkg.in/yaml.v2"
)

func MapHandler(urlPaths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		destUrl, ok := urlPaths[path]
		fmt.Println(path)
		if ok {
			http.Redirect(w,r, destUrl, http.StatusFound)
			return 
		}
		fallback.ServeHTTP(w,r)
	}
}

type pathUrl struct {
	Path string `json:path`
	URL string `json:url`
}

func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//parse and convert yaml in bytes to an array of yaml configurations
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	//convert array of yaml configurations to map of path and url
	mapUrls := make(map[string]string) 
	for _, pathObj := range pathUrls {
		mapUrls[pathObj.Path] = pathObj.URL
	}

	return MapHandler(mapUrls, fallback), nil


}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey you made it to the front page"))
}

func main() {
	sampleUrlMap := map[string]string {
		"/http": "https://golang.org/pkg/net/http/",
		"/compress": "https://golang.org/pkg/compress/",
		"/context": "https://golang.org/pkg/context/",
	}
	sampleYaml := `
		- path: /urlshort
		url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
	`

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	
	mapHandler := MapHandler(sampleUrlMap, mux)
	yamlHandler,err := YamlHandler([]byte(sampleYaml), mapHandler)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("listening on port 8080........")
	http.ListenAndServe(":8080", yamlHandler)

}