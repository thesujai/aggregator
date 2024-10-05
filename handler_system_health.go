package main

import (
	"io"
	"net/http"
)

func systemHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "SystemUp")
}
