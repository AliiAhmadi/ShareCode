package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
	"github.com/justinas/nosurf"
)

var Middlewares = []func(http.Handler) http.Handler{}

func (app *application) secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("X-XSS-Protection", "1; mode=block")
		writer.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(writer, request)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", request.RemoteAddr, request.Proto, request.Method, request.RequestURI)

		next.ServeHTTP(writer, request)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer.Header().Set("Connection", "close")
				// Internal server error
				app.serverError(writer, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(writer, request)
	})
}

func (app *application) requiredAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if app.authenticatedUser(request) == nil {
			http.Redirect(writer, request, "/user/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		exists := app.session.Exists(request, "userID")
		if !exists {
			next.ServeHTTP(writer, request)
			return
		}

		user, err := app.users.Get(app.session.GetInt(request, "userID"))
		if err == models.ErrorNoRecord {
			app.session.Remove(request, "userID")
			next.ServeHTTP(writer, request)
			return
		} else if err != nil {
			app.serverError(writer, err)
			return
		}

		ctx := context.WithValue(request.Context(), contextKeyUser, user)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
