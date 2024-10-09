package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thesujai/aggregator/internal/auth"
)

type Middleware func(http.Handler) http.Handler

func use(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func (cfg *apiConfig) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		user_id, err := cfg.DB.GetUserId(r.Context(), api_key)
		if err != nil {
			http.Error(w, "user doesn't exists", 400)
			return
		}
		w.Header().Set("userID", user_id.String())
		next.ServeHTTP(w, r)
	})
}
