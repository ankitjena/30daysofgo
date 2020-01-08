package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
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
func YAMLHandler(yml []byte, json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls := make(map[string]string)
	pathURLs, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	yamlPathsToUrls := buildMap(pathURLs)

	for path, url := range yamlPathsToUrls {
		pathToUrls[path] = url
	}

	pathURLs, err = parseJSON(json)
	if err != nil {
		return nil, err
	}

	jsonPathsToUrls := buildMap(pathURLs)

	for path, url := range jsonPathsToUrls {
		pathToUrls[path] = url
	}

	return MapHandler(pathToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func parseJSON(data []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := json.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
