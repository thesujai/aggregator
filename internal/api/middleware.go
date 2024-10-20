package api

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

func (cfg *Config) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func (cfg *Config) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		userID, err := cfg.DB.GetUserId(r.Context(), apiKey)
		if err != nil {
			http.Error(w, "user doesn't exist", 400)
			return
		}

		w.Header().Set("userID", userID.String())
		next.ServeHTTP(w, r)
	})
}
