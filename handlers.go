package main

import "net/http"

// middleware is a function that wraps http.Handlers
// proving functionality before and after execution
// of the h handler.
type middleware func(h http.Handler) http.Handler

func newServer() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	handler = addMiddleware(handler)
	return handler
}

func addRoutes(mux http.Handler) http.Handler {
	return mux
}

func addMiddleware(handler http.Handler) http.Handler {
	return handler
}
