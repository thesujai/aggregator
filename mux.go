package main

import "net/http"

type AggregatorMux struct {
	http.ServeMux
}

func (am *AggregatorMux) GET(pattern string, handler http.Handler) {
	am.Handle(pattern, methodHandler(http.MethodGet, handler))
}

func (am *AggregatorMux) POST(pattern string, handler http.Handler) {
	am.Handle(pattern, methodHandler(http.MethodPost, handler))
}

func (am *AggregatorMux) PUT(pattern string, handler http.Handler) {
	am.Handle(pattern, methodHandler(http.MethodPut, handler))
}

func (am *AggregatorMux) DELETE(pattern string, handler http.Handler) {
	am.Handle(pattern, methodHandler(http.MethodDelete, handler))
}

func methodHandler(method string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
