package main

import "net/http"

func registerRoutes() *AggregatorMux {
	mux := &AggregatorMux{}
	apiCfg := GetAPIConfig()
	mux.GET("/healthz", http.HandlerFunc(systemHealth), logger)
	mux.POST("/users/create", http.HandlerFunc(apiCfg.createUser), logger)
	mux.GET("/users", http.HandlerFunc(apiCfg.getUser), logger)
	return mux
}
