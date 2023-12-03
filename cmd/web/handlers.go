package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(writer, err)
		return
	}

	app.render(writer, request, "home.page.tmpl", &templateData{
		Snippets: snippets,
	})
}

func (app *application) showSnippet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get(":id"))

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

	app.render(writer, request, "show.page.tmpl", &templateData{
		Snippet: snippet,
	})
}

func (app *application) createSnippet(writer http.ResponseWriter, request *http.Request) {

	title := "test title"
	content := "here is content"
	expires := "3"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(writer, err)
		return
	}

	// Redirect the user to the relavent page for the snippet.
	http.Redirect(writer, request, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Creating a new snippet ...."))
}
