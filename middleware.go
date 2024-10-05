package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func applyMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares); i >= 0; i-- {
		handler = middlewares[0](handler)
	}
	return handler
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v %v \n", r.Method, r.URL.Path)
	})
}
