package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/auth"
	"github.com/thesujai/aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	var p User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Some Error Occured", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      p.Name,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user %v", err), http.StatusBadRequest)
	}
	RespondWithJSON(w, http.StatusCreated, User(user))
}

func (apiCfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetApiKey(r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), api_key)
	if err != nil {
		http.Error(w, "user doesn't exists", 400)
		return
	}

	RespondWithJSON(w, http.StatusFound, User(user))
}
