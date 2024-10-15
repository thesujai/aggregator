package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thesujai/aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func GetAPIConfig() *apiConfig {
	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apiCgf := apiConfig{
		DB: dbQueries,
	}

	return &apiCgf
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable is not set")
	}

	serv := &http.Server{
		Addr:    ":" + port,
		Handler: registerRoutes(),
	}
	go processFeeds(GetAPIConfig().DB, 10, time.Minute)
	fmt.Printf("Server Running and Listening on %v\n", serv.Addr)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected Error", err)
	}
}
