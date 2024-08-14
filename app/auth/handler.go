package auth

import (
	"net/http"
	"strings"

	"github.com/AxelTahmid/tinker/db"
	"github.com/AxelTahmid/tinker/pkg/bcrypt"
	"github.com/AxelTahmid/tinker/pkg/jwt"
	"github.com/AxelTahmid/tinker/pkg/message"
	"github.com/AxelTahmid/tinker/pkg/request"
	"github.com/AxelTahmid/tinker/pkg/respond"
	"github.com/AxelTahmid/tinker/pkg/validate"
)

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(&w)

	var req LoginRequest

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	req.Email = strings.ToLower(req.Email)

	validationErrs := validate.Check(a.validator, req)
	if validationErrs != nil {
		reply.Status(http.StatusBadRequest).WithErrs(validationErrs)
		return
	}

	fetchedUser, err := a.db.User.GetByEmail(r.Context(), req.Email)
	if err != nil {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	isPasswordValid := bcrypt.VerifyPassword(req.Password, fetchedUser.Password)
	if !isPasswordValid {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrPassOrUserIncorrect)
		return
	}

	tokens, err := jwt.IssueTokenPair(jwt.UserClaims{
		Id:    fetchedUser.Id,
		Email: fetchedUser.Email,
		Role:  fetchedUser.Role,
	})
	if err != nil {
		reply.Status(http.StatusInternalServerError).WithErr(err)
		return
	}

	reply.Status(http.StatusOK).WithJson(respond.Standard{
		Message: message.SuccessLogin,
		Data:    LoginResponse{tokens},
	})
}

func (a *Auth) register(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(&w)

	var req db.InsertUser

	err := request.DecodeJSON(w, r, &req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	validationErrs := validate.Check(a.validator, req)
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

	err = a.db.User.InsertOne(r.Context(), req)
	if err != nil {
		reply.Status(http.StatusBadRequest).WithErr(err)
		return
	}

	reply.Status(http.StatusCreated).WithJson(respond.Standard{
		Message: message.SuccessRegister,
	})
}

func (a *Auth) me(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(&w)
	ctx := r.Context()

	userClaim, ok := jwt.ParseClaimsCtx(ctx)
	if !ok {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrUnauthorized)
		return
	}

	fetchedUser, err := a.db.User.GetByEmail(ctx, userClaim.Subject)
	if err != nil {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrNoRecord)
		return
	}

	reply.Status(http.StatusOK).WithJson(respond.Standard{
		Message: message.SuccessMe,
		Data:    fetchedUser,
	})
}

func (a *Auth) refresh(w http.ResponseWriter, r *http.Request) {
	reply := respond.Write(&w)
	ctx := r.Context()

	userClaim, ok := jwt.ParseClaimsCtx(ctx)
	if !ok {
		reply.Status(http.StatusBadRequest).WithErr(message.ErrBadRequest)
		return
	}

	fetchedUser, err := a.db.User.GetByEmail(ctx, userClaim.Subject)
	if err != nil {
		reply.Status(http.StatusUnauthorized).WithErr(message.ErrNoRecord)
		return
	}

	token, err := jwt.IssueAccessToken(jwt.UserClaims{
		Id:    fetchedUser.Id,
		Email: fetchedUser.Email,
		Role:  fetchedUser.Role,
	})
	if err != nil {
		reply.Status(http.StatusInternalServerError).WithErr(err)
		return
	}

	reply.Status(http.StatusOK).WithJson(respond.Standard{
		Message: message.SucessRefresh,
		Data:    RefreshResponse{token},
	})
}
