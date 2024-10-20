package handlers

import (
	"net/http"

	"github.com/thesujai/aggregator/internal/utils"
)

func SystemHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct {
		Message string `json:"msg"`
	}{Message: "Server Running"})
}
