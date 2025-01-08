package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the API key from the request headers
// Example:
// Bearer {api_key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}
	if vals[0] != "Bearer" {
		return "", errors.New("malformed first part of Authorization header")
	}
	return vals[1], nil
}