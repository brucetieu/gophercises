package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/brucetieu/gophercises/ex02/urlshortener"
)

func main() {
	yamlPtr := flag.String("yaml", "", "A yaml file with a set of paths and urls")

	flag.Parse()

	isFlag := false
	yamlFile := []byte{}
	var err error
	var yamlHandler http.HandlerFunc

	if *yamlPtr != "" {
		yamlFile, err = ioutil.ReadFile(*yamlPtr)
		if err != nil {
			panic(err)
		}

		isFlag = true
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
  - path: /urlshort
    url: https://github.com/gophercises/urlshort
  - path: /urlshort-final
    url: https://github.com/gophercises/urlshort/tree/solution
  `
	
	if isFlag {
		yamlHandler, err = urlshortener.YAMLHandler(yamlFile, mapHandler)
		
	} else {
		yamlHandler, err = urlshortener.YAMLHandler([]byte(yaml), mapHandler)
	}

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
