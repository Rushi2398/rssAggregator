package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rushi2398/rssAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := paramters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create a feed: %s", err))
		return
	}
	responseWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %s", err))
		return
	}
	responseWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
