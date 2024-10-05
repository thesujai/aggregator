package main

import "net/http"

func registerRoutes() *AggregatorMux {
	mux := &AggregatorMux{}
	mux.GET("/healthz", http.HandlerFunc(systemHealth), logger)
	return mux
}
