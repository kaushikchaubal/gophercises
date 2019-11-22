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
		"/u": "https://godoc.org/github.com/gophercises/urlshort",
		"/y": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mapHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
