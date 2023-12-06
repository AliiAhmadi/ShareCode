package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliiAhmadi/ShareCode/pkg/forms"
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

	err := request.ParseForm()
	if err != nil {
		app.clientError(writer, http.StatusBadRequest)
		return
	}

	form := forms.New(request.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(writer, request, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	id, err := app.snippets.Insert(
		form.Get("title"),
		form.Get("content"),
		form.Get("expires"),
	)

	if err != nil {
		app.serverError(writer, err)
		return
	}

	app.session.Put(request, "flash", "Snippet successfully created!")

	// Redirect the user to the relavent page for the snippet.
	http.Redirect(writer, request, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUserForm(writer http.ResponseWriter, request *http.Request) {
	app.render(writer, request, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		app.clientError(writer, http.StatusBadRequest)
	}

	form := forms.New(request.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(
			writer,
			request,
			"signup.page.tmpl",
			&templateData{
				Form: form,
			},
		)
		return
	}

	_, err = app.users.Insert(
		form.Get("name"),
		form.Get("email"),
		form.Get("password"),
	)
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Email address is already in use")
		app.render(writer, request, "signup.page.tmpl", &templateData{
			Form: form,
		})
		return
	} else if err != nil {
		app.serverError(writer, err)
		return
	}

	app.session.Put(request, "flash", "Your signup was successful. Please log in.")
	http.Redirect(writer, request, "/user/login", http.StatusSeeOther)

}

func (app *application) loginUserForm(writer http.ResponseWriter, request *http.Request) {

}

func (app *application) loginUser(writer http.ResponseWriter, request *http.Request) {

}

func (app *application) logoutUser(writer http.ResponseWriter, request *http.Request) {

}
