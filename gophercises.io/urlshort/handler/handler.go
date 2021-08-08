// Package handler implements urlshort http.Handlers
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathData struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if newDest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, newDest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	pathMap := buildPathsMap(parsedYaml)

	return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//       [
// 			  {
// 				 "path" : "/http",
//        		  "url" : "https://pkg.go.dev/net/http#HandleFunc"
//             }
//       ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
func JSONHandler(j []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonPath, err := parseJSON(j)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	paths := buildPathsMap(jsonPath)

	return MapHandler(paths, fallback), nil
}

func parseJSON(j []byte) ([]pathData, error) {
	var path []pathData
	err := json.Unmarshal(j, &path)
	if err != nil {
		return nil, err
	}
	return path, nil
}

func parseYaml(y []byte) ([]pathData, error) {
	var pats []pathData
	err := yaml.Unmarshal([]byte(y), &pats)
	if err != nil {
		return nil, err
	}

	return pats, err
}

func buildPathsMap(p []pathData) map[string]string {
	pathMap := make(map[string]string)
	for _, e := range p {
		pathMap[e.Path] = e.URL
	}
	return pathMap
}

