package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		app.notFound(writer)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(writer, err)
		return
	}

	data := &templateData{
		Snippets: snippets,
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

	err = ts.Execute(writer, data)

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

	snippet, err := app.snippets.Get(id)

	if err == models.ErrorNoRecord {
		app.notFound(writer)
		return
	} else if err != nil {
		app.serverError(writer, err)
		return
	}

	data := &templateData{
		Snippet: snippet,
	}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(writer, err)
		return
	}

	err = ts.Execute(writer, data)

	if err != nil {
		app.serverError(writer, err)
	}
}

func (app *application) createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		app.clientError(writer, http.StatusMethodNotAllowed)
		return
	}

	title := "test title"
	content := "here is content"
	expires := "3"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(writer, err)
		return
	}

	// Redirect the user to the relavent page for the snippet.
	http.Redirect(writer, request, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
