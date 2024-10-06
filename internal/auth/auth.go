package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Shoule be added in headers: Authorization: api_key {actual key}
// eg: Authorization: api_key awmo83nm98oj-0m-3qrfqw
func GetApiKey(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	split_auth := strings.Split(authorization, " ")
	if len(split_auth) != 2 && split_auth[0] != "api_key" {
		return "", errors.New("invalid authorization; should be format: api_key {actual key}")
	}

	return split_auth[1], nil

}
