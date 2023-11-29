package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(writer, err)
		return
	}

	err = ts.Execute(writer, nil)

	if err != nil {
		app.serverError(writer, err)
		return
	}
}

func (app *application) showSnippet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(writer)
		return
	}

	fmt.Fprintf(writer, "Display snippet with ID %d ...", id)
}

func (app *application) createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	writer.Write([]byte("Creating a new snippet..."))
}
