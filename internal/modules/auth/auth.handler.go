package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	email := fmt.Sprintf(`the email is %s`, u.Email)
	password := u.Password

	log.Printf("email: %v, password: %v", email, password)

	err = json.NewEncoder(w).Encode(User{email, password})
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// write the handler for /register with email & password

// write the handler for /refresh with refresh token

// write the handler for /logout with refresh token
