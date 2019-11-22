package main

import (
	"fmt"
	"gophercises/exercise-2/urlshortner"
	"log"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/u": "https://www.google.com",
		"/y": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	yamlFileName := "short-urls.yaml"
	yamlHandler := urlshortner.YAMLHandler(yamlFileName, mapHandler)

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the default path... be more creative!")
}
