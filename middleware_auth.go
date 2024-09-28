package main

import (
	"fmt"
	"net/http"

	"github.com/minzhoudu/rss-aggregator/internal/auth"
	"github.com/minzhoudu/rss-aggregator/internal/database"
)

type authenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfiguration) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error getting the user by api key: %v", err))
			return
		}

		handler(w, r, user)
	}
}
