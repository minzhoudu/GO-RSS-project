package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minzhoudu/rss-aggregator/internal/database"
)

func (apiCfg *apiConfiguration) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	respondWithJson(w, 201, convertDbUserToUserDto(user))
}

func (apiCfg *apiConfiguration) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, convertDbUserToUserDto(user))
}
