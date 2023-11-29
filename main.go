package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Listen on localhost:4000")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	writer.Write([]byte("Welcome to ShareCode"))
}

func showSnippet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a specific snippet..."))
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		writer.WriteHeader(405)
		writer.Write([]byte("Method not allowed"))
		return
	}

	writer.Write([]byte("Create a new snippet..."))
}
