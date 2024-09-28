package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extacts the API KEY from the Headers of the HTTP request
// EXAMPLE: Authorization: ApiKey {insert Api Key here}
func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("unauthorized access")
	}

	vals := strings.Split(authHeader, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Auth Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed name of the Auth Header")
	}

	return vals[1], nil
}
