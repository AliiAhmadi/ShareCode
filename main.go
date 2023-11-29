package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "Snippet with id %d", id)
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		// writer.WriteHeader(405)
		// writer.Write([]byte("Method not allowed"))
		http.Error(writer, "Method Not Allowed", 405)
		return
	}

	writer.Write([]byte("Create a new snippet..."))
}
