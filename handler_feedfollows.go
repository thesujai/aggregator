package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/database"
)

type FeedID struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func (cfg *apiConfig) followFeed(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(w.Header().Get("userID"))
	feedId := FeedID{}
	err := json.NewDecoder(r.Body).Decode(&feedId)
	if err != nil {
		RespondWithError(w, 400, "feed_id should be a uuid")
		return
	}
	err = cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
		FeedID:    feedId.FeedId,
	})
	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	RespondWithJSON(w, 201, struct{}{})

}
