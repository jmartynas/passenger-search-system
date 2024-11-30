package main

import (
	"embed"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static/*
var staticFiles embed.FS

func newServer() http.Handler {
	mux := mux.NewRouter()
	mux = addLoginRoutes(mux)
	addRegisterRoutes(mux)

	var handler http.Handler = mux
	// handler = userEndpoints(handler)
	// handler = addMiddleware(handler)
	return handler
}

func addLoginRoutes(mux *mux.Router) *mux.Router {
	mux.Handle(
		"/static/login.html",
		http.FileServer(http.FS(staticFiles)),
	).Methods("GET")
	mux.Handle("/login", login()).Methods("POST")
	return mux
}

func addRegisterRoutes(mux *mux.Router) {
	mux.Handle(
		"/static/register.html",
		http.FileServer(http.FS(staticFiles)),
	).Methods("GET")
	mux.Handle("/register", register()).Methods("POST")
}
