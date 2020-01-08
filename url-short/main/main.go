package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	urlshort "github.com/ankitjena/30daysofgo/url-short"
)

func main() {
	yamlFile := flag.String("yaml", "url.yaml", "yaml file containing the paths and urls")
	jsonFile := flag.String("json", "url.json", "json file containing the paths and urls")
	flag.Parse()

	yaml, err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		log.Fatalf("Cannot open file : %s", *yamlFile)
	}

	json, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		log.Fatalf("Cannot open file : %s", *jsonFile)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), []byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
