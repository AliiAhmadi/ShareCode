package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) serverError(writer http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(writer http.ResponseWriter, status int) {
	http.Error(writer, http.StatusText(status), status)
}

func (app *application) notFound(writer http.ResponseWriter) {
	app.clientError(writer, http.StatusNotFound)
}

func (app *application) render(writer http.ResponseWriter, request *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		app.serverError(writer, fmt.Errorf("the template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	// Execute template set, passing in any dynamic data.
	err := ts.Execute(writer, app.addDefaultData(td, request))
	if err != nil {
		app.serverError(writer, err)
	}

	buf.WriteTo(writer)
}

func (app *application) addDefaultData(td *templateData, request *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.Flash = app.session.PopString(request, "flash")
	td.CurrentYear = time.Now().Year()
	return td
}
