package api

import (
	"net/http"

	"github.com/thesujai/aggregator/internal/api/handlers"
)

func RegisterRoutes(cfg *Config) http.Handler {
	publicMux := http.NewServeMux()
	publicMux.Handle("GET /healthz", http.HandlerFunc(handlers.SystemHealth))
	publicMux.Handle("POST /users", http.HandlerFunc(cfg.CreateUser))
	publicMux.Handle("GET /users", http.HandlerFunc(cfg.GetUser))
	publicMux.Handle("GET /feeds", http.HandlerFunc(cfg.GetAllFeeds))

	protectedMux := http.NewServeMux()
	protectedMux.Handle("POST /feeds", http.HandlerFunc(cfg.CreateFeed))
	protectedMux.Handle("GET /feeds", http.HandlerFunc(cfg.GetFeedByUser))
	protectedMux.Handle("GET /followedfeeds", http.HandlerFunc(cfg.GetFollowedFeeds))
	protectedMux.Handle("POST /followfeed", http.HandlerFunc(cfg.FollowFeed))

	mux := http.NewServeMux()
	mux.Handle("/", use(protectedMux, cfg.AuthUser))
	mux.Handle("/public/", http.StripPrefix("/public", publicMux))

	v1Mux := http.NewServeMux()
	v1Mux.Handle("/v1/", http.StripPrefix("/v1", use(mux, cfg.Logger)))

	return v1Mux
}
