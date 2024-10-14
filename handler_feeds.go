package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/auth"
	"github.com/thesujai/aggregator/internal/database"
)

type Feeds struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createFeed(w http.ResponseWriter, r *http.Request) {
	feed := Feeds{}
	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}
	user_id, err := uuid.Parse(w.Header().Get("userID"))
	if err != nil {
		http.Error(w, "user doesn't exists", http.StatusNotFound)
		return
	}
	new_feed, err := cfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    user_id,
	})
	if err != nil {
		http.Error(w, "feed already exists", http.StatusBadRequest)
		return
	}
	RespondWithJSON(w, http.StatusCreated, Feeds(new_feed))
}

func (cfg *apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	RespondWithJSON(w, http.StatusOK, feeds)

}

// TODO: feeds should be type casted to the Feed defined in this file for json compatibility
// so should be done with a loop, make sure the loop is concurrent

func (cfg *apiConfig) getFeedByUser(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetApiKey(r)
	if err != nil {
		http.Error(w, "user doesn't exists", http.StatusNotFound)
		return
	}
	feeds, err := cfg.DB.GetFeedByUser(r.Context(), api_key)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	RespondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *apiConfig) getFollowedFeeds(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(w.Header().Get("userID"))
	feeds, err := cfg.DB.GetFollowedFeeds(r.Context(), userId)
	if err != nil {
		RespondWithError(w, 400, "no feeds following")
	}
	RespondWithJSON(w, http.StatusOK, feeds)
}
