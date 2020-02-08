package main

import (
	"fmt"
	"net/http"
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

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey you made it to the front page"))
}

func main() {
	sampleUrlMap := map[string]string {
		"/http": "https://golang.org/pkg/net/http/",
		"/compress": "https://golang.org/pkg/compress/",
		"/context": "https://golang.org/pkg/context/",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	
	mapHandler := MapHandler(sampleUrlMap, mux)

	http.ListenAndServe(":8080", mapHandler)

}