package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GETAPIKey extracts an API Key
// from the headers on an HTTP Request
// Authorization : ApiKey {insert API Key Here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Auth Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Auth Header")
	}
	return vals[1], nil
}
