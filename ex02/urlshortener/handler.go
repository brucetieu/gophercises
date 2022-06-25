package urlshortener

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Request struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := pathsToUrls[r.URL.Path]
		if ok {
			fmt.Println(r.URL.Path, pathsToUrls[r.URL.Path])
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusSeeOther)
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
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler() {
	
}

func parseYAML(yml []byte) ([]Request, error) {
	out := []Request{}

	err := yaml.Unmarshal([]byte(yml), &out)
	if err != nil {
		log.Fatalf("Cannot unmarshall data: %+v", err)
	}
	return out, err
}

func buildMap(data []Request) map[string]string {
	m := make(map[string]string, len(data))

	for _, req := range data {
		m[req.Path] = req.Url
	}

	return m
}
