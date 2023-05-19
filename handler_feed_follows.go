package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/poompichayout/rss-aggregator/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't create feed follow:", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't get feeds:", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
