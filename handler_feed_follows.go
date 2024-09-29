package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/minzhoudu/rss-aggregator/internal/database"
)

func (apiCfg *apiConfiguration) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type handleFeedParams struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := handleFeedParams{}
	decoder.Decode(&params)

	_, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to follow the feed. Error: %v", err))
		return
	}

	respondWithJson(w, 200, struct{ Status string }{
		Status: "success",
	})
}

func (apiCfg *apiConfiguration) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to get feed follows. Error: %v", err))
		return
	}

	respondWithJson(w, 200, convertToFeedFollowsSliceDto(feedFollows))
}

func (apiCfg *apiConfiguration) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	feedFollowUUID, err := uuid.Parse(feedFollowId)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to parse UUID of the feed follow. Error: %v", err))
		return
	}

	deleteFeedError := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowUUID,
		UserID: user.ID,
	})
	if deleteFeedError != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to remove feed follow. Error: %v", err))
		return
	}

	respondWithJson(w, 200, struct {
		Status string `json:"status"`
	}{Status: "success"})
}
