package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AxelTahmid/golang-starter/internal/utils"
)

// routing & validation

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Password = hash

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// write the handler for /register with email & password

// write the handler for /refresh with refresh token

// write the handler for /logout with refresh token
