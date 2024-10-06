package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
