package urlshort

import (
	"encoding/json"
	"net/http"

	yaml "github.com/go-yaml/yaml"
)

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse the provided JSON and return an
// http.HandlerFunc (which also implements http.Handler) that
// will attempt to map any paths to their corresponding URL.
// If the path is not provided in the JSON, then the fallback
// http.Handler will be called instead.
//
// JSON is expected to be in the format:
// 
// [
//   {
//     "path": "/urlshort",
//     "url": "https://github.com/gophercises/urlshort"
//   },
// ]
// 
// The only errors that can be returned all relate to having invalid JSON data.
// 
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}

// parseJSON takes raw JSON input as a byte slice and returns a
// slice of pathURL structs and an error if the JSON is invalid.
func parseJSON(rawJson []byte) ([]pathURL, error) {
	var pathURLs []pathURL

	err := json.Unmarshal(rawJson, &pathURLs)

	if err != nil {
		return nil, err
	}

	return pathURLs, nil
}

// parseYAML takes raw YAML input as a byte slice and returns a
// slice of pathURL structs and an error if the YAML is invalid.
func parseYAML(rawYaml []byte) ([]pathURL, error) {
	var pathURLs []pathURL

	err := yaml.Unmarshal(rawYaml, &pathURLs)

	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

// buildMap takes a slice of pathURL structs and returns a map where the keys
// are the paths and the values are the corresponding URLs.
func buildMap(pathURLs []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)

	for _, p := range pathURLs {
		pathsToUrls[p.Path] = p.URL
	}

	return pathsToUrls
}
