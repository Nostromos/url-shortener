package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/Nostromos/url-shortener"
)

const (
	defaultYAMLPath = "urls.yaml"
	defaultJSONPath = "urls.json"
)

func main() {
	defaultPath := flag.String("p", defaultJSONPath, "Path to file to use for shortening")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Read YAML data from path and
	// convert to bytesliceable data
	urlData, err := os.ReadFile(*defaultPath)
	if err != nil {
		panic(err)
	}

	filetype := path.Ext(*defaultPath)

	if filetype == ".yaml" {
		// Build the YAMLHandler using the mapHandler as the
		// fallback
		yamlHandler, err := urlshort.YAMLHandler([]byte(urlData), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("YAML: Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else if filetype == ".json" {
		// Build the JSONHandler using the mapHandler as the
		// fallback
		jsonHandler, err := urlshort.JSONHandler([]byte(urlData), mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("JSON: Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else {
		fmt.Println("Mux: Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)
	}

	fmt.Println("Starting the server on :8080")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
