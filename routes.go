package main

import "net/http"

func registerRoutes() http.Handler {
	mux := http.NewServeMux()
	cfg := GetAPIConfig()
	mux.Handle("GET /healthz", http.HandlerFunc(systemHealth))
	mux.Handle("POST /users", http.HandlerFunc(cfg.createUser))
	mux.Handle("GET /users", http.HandlerFunc(cfg.getUser))
	mux_with_middleware := use(mux, logger)
	return mux_with_middleware
}
