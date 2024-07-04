package auth

import (
	"net/http"
	"strings"

	"github.com/AxelTahmid/golang-starter/internal/utils/bcrypt"
	"github.com/AxelTahmid/golang-starter/internal/utils/message"
	"github.com/AxelTahmid/golang-starter/internal/utils/request"
	"github.com/AxelTahmid/golang-starter/internal/utils/respond"
	"github.com/AxelTahmid/golang-starter/internal/utils/validate"
)

func (handler AuthHandler) login(w http.ResponseWriter, r *http.Request) {
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

	fetchedUser, err := handler.user.getOne(r.Context(), req.Email)
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

func (handler AuthHandler) register(w http.ResponseWriter, r *http.Request) {
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
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	req.Email = strings.ToLower(req.Email)
	req.Password = hash

	err = handler.user.insertOne(r.Context(), req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Status(w, http.StatusCreated)
}
