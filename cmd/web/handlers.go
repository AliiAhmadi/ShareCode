package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(writer, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server Error", 500)
		return
	}
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
