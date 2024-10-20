package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/database"
	"github.com/thesujai/aggregator/internal/utils"
)

type FeedID struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func (cfg *Config) FollowFeed(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.Parse(w.Header().Get("userID"))
	feedId := FeedID{}
	err := json.NewDecoder(r.Body).Decode(&feedId)
	if err != nil {
		utils.RespondWithError(w, 400, "feed_id should be a uuid")
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
		utils.RespondWithError(w, 400, err.Error())
		return
	}
	utils.RespondWithJSON(w, 201, struct{}{})

}
