package auth

import (
	"fmt"
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
		respond.Write(w).Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	req.Email = strings.ToLower(req.Email)

	validationErrs := validate.Validate(v, req)
	if validationErrs != nil {
		respond.Write(w).Status(http.StatusBadRequest).WithErrs(validationErrs)
		return
	}

	fetchedUser, err := handler.user.getOne(r.Context(), req.Email)
	if err != nil {
		respond.Write(w).Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	isPasswordValid := bcrypt.VerifyPassword(req.Password, fetchedUser.Password)
	if !isPasswordValid {
		respond.Write(w).Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	respond.Write(w).Status(http.StatusOK).WithJson(respond.Standard{
		Message: fmt.Sprintf("%s %s", message.SuccessLogin, fetchedUser.Email),
		Data:    fetchedUser,
	})
}

func (handler AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		respond.Write(w).Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	validationErrs := validate.Validate(v, req)
	if validationErrs != nil {
		respond.Write(w).Status(http.StatusBadRequest).WithErrs(validationErrs)
		return
	}

	hash, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		respond.Write(w).Status(http.StatusBadRequest).WithErr(err)
		return
	}

	req.Email = strings.ToLower(req.Email)
	req.Password = hash

	err = handler.user.insertOne(r.Context(), req)
	if err != nil {
		respond.Write(w).Status(http.StatusBadRequest).WithErr(err)
		return
	}

	respond.Write(w).Status(http.StatusCreated).WithJson(respond.Standard{
		Message: fmt.Sprintf("%s %s", message.SuccessRegister, req.Email),
	})
}
