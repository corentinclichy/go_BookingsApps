package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurfHandler
// add CSRF protection to all requests
func NosurfHandler(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{HttpOnly: true, Path: "/", Secure: app.InProduction, SameSite: http.SameSiteLaxMode})

	return csrfHandler
}

// Session load
// load and save session on every request
func Sessionload(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
