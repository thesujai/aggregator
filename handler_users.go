package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thesujai/aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
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
