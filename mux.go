package main

import "net/http"

type AggregatorMux struct {
	http.ServeMux
}

func (am *AggregatorMux) GET(pattern string, handler http.Handler, middlewares ...Middleware) {
	am.Handle(pattern, applyMiddleware(methodHandler(http.MethodGet, handler), middlewares...))
}

func (am *AggregatorMux) POST(pattern string, handler http.Handler, middlewares ...Middleware) {
	am.Handle(pattern, applyMiddleware(methodHandler(http.MethodPost, handler), middlewares...))
}

func (am *AggregatorMux) PUT(pattern string, handler http.Handler, middlewares ...Middleware) {
	am.Handle(pattern, applyMiddleware(methodHandler(http.MethodPut, handler), middlewares...))
}

func (am *AggregatorMux) DELETE(pattern string, handler http.Handler, middlewares ...Middleware) {
	am.Handle(pattern, applyMiddleware(methodHandler(http.MethodDelete, handler), middlewares...))
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
