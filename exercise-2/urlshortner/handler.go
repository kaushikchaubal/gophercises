package urlshortner

import (
	"io/ioutil"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for path := range pathsToUrls {
			if r.URL.Path == path {
				http.Redirect(w, r, pathsToUrls[path], 301)
			}
		}

		fallback.ServeHTTP(w, r)

	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
func YAMLHandler(yamlFileName string, fallback http.Handler) http.HandlerFunc {
	pathsToUrls := parseYAML(yamlFileName)
	return MapHandler(pathsToUrls, fallback)
}

func parseYAML(yamlFileName string) map[string]string {
	yamlFile, err := ioutil.ReadFile(yamlFileName)
	if err != nil {
		panic("Incorrect file!")
	}

	pathsToUrls := make(map[string]string)

	err = yaml.Unmarshal(yamlFile, pathsToUrls)
	if err != nil {
		panic("Incorrect yaml format")
	}

	return pathsToUrls

}
