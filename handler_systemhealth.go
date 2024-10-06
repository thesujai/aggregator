package main

import (
	"net/http"
)

func systemHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, struct {
		Message string `json:"msg"`
	}{Message: "Server Running"})
}
