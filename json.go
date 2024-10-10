package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	if code < 500 {
		fmt.Println("Unexpected Error: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	json_body, err := json.Marshal(Error{
		Error: string(err.Error()),
	})
	if err != nil {
		fmt.Println("Unexpected Error: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(json_body)
}

func RespondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json_body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error Occured while parsing json: %v", err)
		w.WriteHeader(500)
	}
	// The sequence was important as per the docs
	w.WriteHeader(code)
	w.Write(json_body)
}
