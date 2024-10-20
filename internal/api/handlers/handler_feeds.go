package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/auth"
	"github.com/thesujai/aggregator/internal/database"
	"github.com/thesujai/aggregator/internal/utils"
)

type Feeds struct {
	ID            uuid.UUID    `json:"id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Name          string       `json:"name"`
	Url           string       `json:"url"`
	UserID        uuid.UUID    `json:"user_id"`
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
}

func (cfg *Config) CreateFeed(w http.ResponseWriter, r *http.Request) {
	feed := Feeds{}
	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	userID, err := uuid.Parse(w.Header().Get("userID"))
	if err != nil {
		http.Error(w, "user doesn't exist", http.StatusNotFound)
		return
	}

	newFeed, err := cfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    userID,
	})
	if err != nil {
		http.Error(w, "feed already exists", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, Feeds(newFeed))
}

func (cfg *Config) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *Config) GetFeedByUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r)
	if err != nil {
		http.Error(w, "user doesn't exist", http.StatusNotFound)
		return
	}

	feeds, err := cfg.DB.GetFeedByUser(r.Context(), apiKey)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.ConvertDBStructSliceToResponseStructSlice(feeds, Feeds{}))
}

func (cfg *Config) GetFollowedFeeds(w http.ResponseWriter, r *http.Request) {
	userID, _ := uuid.Parse(w.Header().Get("userID"))
	feeds, err := cfg.DB.GetFollowedFeeds(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, 400, "no feeds following")
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.ConvertDBStructSliceToResponseStructSlice(feeds, Feeds{}))
}
