package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AxelTahmid/golang-starter/internal/utils"
	"github.com/go-playground/validator/v10"
)

var authService = AuthService{}

func (a AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Email = strings.ToLower(user.Email)

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check for user in database, if not found return 404
	fetchedUser, err := authService.getUser(r.Context(), a.postgres.DB, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if found, verify password
	isPasswordValid := utils.VerifyPassword(user.Password, fetchedUser.Password)
	// if password is incorrect, return 401
	if !isPasswordValid {
		http.Error(w, "Password is incorrect", http.StatusUnauthorized)
		return
	}

	// if password is correct, return user data with token
	err = json.NewEncoder(w).Encode(fetchedUser)

	// if any other error occurs, return 500
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (a AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user UserRegister

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Email = strings.ToLower(user.Email)
	user.Password = hash

	// insert user into database
	err = authService.insertUser(r.Context(), a.postgres.DB, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if successful, return user datam with token
	// if any other error occurs, return 500
	// if unique violation, return already exists

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
