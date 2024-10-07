package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func use(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}
