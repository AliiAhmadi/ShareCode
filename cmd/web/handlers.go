package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	writer.Write([]byte("Hello from ShareCode"))
}

func showSnippet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "Display snippet with ID %d ...", id)
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		http.Error(writer, "Method Not Allowed", 405)
		return
	}

	writer.Write([]byte("Creating a new snippet..."))
}
