package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
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
