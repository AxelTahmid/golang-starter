package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/AxelTahmid/golang-starter/internal/utils/bcrypt"
	"github.com/AxelTahmid/golang-starter/internal/utils/message"
	"github.com/AxelTahmid/golang-starter/internal/utils/request"
	"github.com/AxelTahmid/golang-starter/internal/utils/respond"
	t "github.com/AxelTahmid/golang-starter/internal/utils/tokens"
	"github.com/AxelTahmid/golang-starter/internal/utils/validate"
)

func (handler AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(w)

	var req LoginRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	req.Email = strings.ToLower(req.Email)

	validationErrs := validate.Check(handler.v, req)
	if validationErrs != nil {
		reply.Status(http.StatusBadRequest).WithErrs(validationErrs)
		return
	}

	fetchedUser, err := handler.user.getOne(r.Context(), req.Email)
	if err != nil {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	isPasswordValid := bcrypt.VerifyPassword(req.Password, fetchedUser.Password)
	if !isPasswordValid {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	tokens, err := t.IssueToken(t.UserClaims{
		Id:    fetchedUser.Id,
		Email: fetchedUser.Email,
		Role:  fetchedUser.Role,
	})
	if err != nil {
		reply.Status(http.StatusInternalServerError).WithErr(err)
		return
	}

	reply.Status(http.StatusOK).WithJson(respond.Standard{
		Message: fmt.Sprintf("%s %s", message.SuccessLogin, fetchedUser.Email),
		Data:    LoginResponse{tokens},
	})
}

func (handler AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(w)

	var req RegisterRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	validationErrs := validate.Check(handler.v, req)
	if validationErrs != nil {
		reply.Status(http.StatusBadRequest).WithErrs(validationErrs)
		return
	}

	hash, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(err)
		return
	}

	req.Email = strings.ToLower(req.Email)
	req.Password = hash

	err = handler.user.insertOne(r.Context(), req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(err)
		return
	}

	reply.Status(http.StatusCreated).WithJson(respond.Standard{
		Message: fmt.Sprintf("%s %s", message.SuccessRegister, req.Email),
	})
}

func (handler AuthHandler) me(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(w)

	userClaim, ok := r.Context().Value(t.AuthReqCtxKey).(*jwt.RegisteredClaims)
	if !ok {
		respond.Write(w).Status(http.StatusUnauthorized).WithErr(message.ErrUnauthorized)
		return
	}

	fetchedUser, err := handler.user.getOne(r.Context(), userClaim.Subject)
	if err != nil {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrNoRecord)
		return
	}

	reply.Status(http.StatusOK).WithJson(respond.Standard{
		Message: fmt.Sprintf("%s %s", message.SuccessLogin, fetchedUser.Email),
		Data:    fetchedUser,
	})
}
