package auth

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/AxelTahmid/golang-starter/internal/utils/bcrypt"
	"github.com/AxelTahmid/golang-starter/internal/utils/message"
	"github.com/AxelTahmid/golang-starter/internal/utils/request"
	"github.com/AxelTahmid/golang-starter/internal/utils/respond"
	"github.com/AxelTahmid/golang-starter/internal/utils/validate"
)

var v = validator.New()
var authService = AuthService{}

func (a AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
		return
	}

	req.Email = strings.ToLower(req.Email)

	validationErrs := validate.Validate(v, req)
	if validationErrs != nil {
		respond.Errors(w, http.StatusBadRequest, validationErrs)
		return
	}

	fetchedUser, err := authService.GetUser(r.Context(), a.postgres.DB, req.Email)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, message.ErrPassOrUserIncorrect)
		return
	}

	isPasswordValid := bcrypt.VerifyPassword(req.Password, fetchedUser.Password)
	if !isPasswordValid {
		respond.Error(w, http.StatusUnauthorized, message.ErrPassOrUserIncorrect)
		return
	}

	respond.Json(w, http.StatusOK, fetchedUser)

}

func (a AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
		return
	}

	validationErrs := validate.Validate(v, req)
	if validationErrs != nil {
		respond.Errors(w, http.StatusBadRequest, validationErrs)
		return
	}

	hash, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		respond.Errors(w, http.StatusBadRequest, validationErrs)
		return
	}

	req.Email = strings.ToLower(req.Email)
	req.Password = hash

	err = authService.InsertUser(r.Context(), a.postgres.DB, req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Status(w, http.StatusCreated)
}
