package main

import "net/http"

func registerRoutes() http.Handler {
	cfg := GetAPIConfig()

	publicMux := http.NewServeMux()
	publicMux.Handle("GET /healthz", http.HandlerFunc(systemHealth))
	publicMux.Handle("POST /users", http.HandlerFunc(cfg.createUser))
	publicMux.Handle("GET /users", http.HandlerFunc(cfg.getUser))

	protectedMux := http.NewServeMux()
	protectedMux.Handle("POST /feeds", http.HandlerFunc(cfg.createFeed))

	mux := http.NewServeMux()
	mux.Handle("/", use(protectedMux, cfg.authUser))
	mux.Handle("/public/", http.StripPrefix("/public", publicMux))

	v1_mux := http.NewServeMux()
	v1_mux.Handle("/v1/", http.StripPrefix("/v1", use(mux, cfg.logger)))

	return v1_mux
}
