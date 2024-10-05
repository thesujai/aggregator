package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serv := &http.Server{
		Addr:    ":8012",
		Handler: registerRoutes(),
	}
	fmt.Printf("Server Running and Listening on %v\n", serv.Addr)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected Error", err)
	}
}
