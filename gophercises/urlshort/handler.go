package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		path, ok := pathsToUrls[req.URL.Path]
		if ok {
			http.Redirect(res, req, path, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
	return nil
}

type parsedPath struct {
	PATH string
	URL  string
}

func ParseYaml(yml []byte) ([]parsedPath, error) {
	var yamlData []parsedPath
	err := yaml.Unmarshal(yml, &yamlData)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func BuildMap(yamlData []parsedPath) map[string]string {
	mapData := make(map[string]string)
	for _, data := range yamlData {
		mapData[data.PATH] = data.URL
	}
	return mapData
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data, err := ParseYaml(yml)
	if err != nil {
		return nil, err
	}
	mapData := BuildMap(data)
	return MapHandler(mapData, fallback), nil
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var data []parsedPath
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	mapData := BuildMap(data)
	return MapHandler(mapData, fallback), nil
}
