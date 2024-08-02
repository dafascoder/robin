package handlers

import (
	"backend/internal/forms"
	"backend/internal/services"
	"backend/internal/utils"
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthServices *services.AuthServices
}

func (h *AuthHandler) HandleSignUp(res http.ResponseWriter, req *http.Request) {

	// Decode the request body

	var form forms.SignUpForm
	form, err := utils.Decode[forms.SignUpForm](req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create account")
		err := utils.Encode(res, req, http.StatusBadRequest, utils.ResponseError{Message: err})
		if err != nil {
			return
		}
	}

	// Validate the form
	if form.Email == "" {
		log.Error().Err(err).Msg("failed to create account")
		err := utils.Encode(res, req, http.StatusBadRequest, "email and password are required")
		if err != nil {
			return
		}
	}

	_, err = h.AuthServices.CreateAccount(context.Background(), form)
	if err != nil {
		log.Error().Err(err).Msg("failed to create account")
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	// Send the response
	err = utils.Encode(res, req, http.StatusOK, utils.Response{Message: "account created"})

	return
}

func (h *AuthHandler) HandleVerifyAccount(res http.ResponseWriter, req *http.Request) {
	var form forms.VerifyAccountForm

	form, err := utils.Decode[forms.VerifyAccountForm](req)
	if err != nil {
		log.Error().Err(err).Msg("failed to verify account")
		err := utils.Encode(res, req, http.StatusBadRequest, utils.ResponseError{Message: err})
		if err != nil {
			return
		}
	}

	err = h.AuthServices.VerifyAccount(context.Background(), form.Email, form.Token)
	if err != nil {
		log.Error().Err(err).Msg("failed to verify account")
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) HandleSignIn(res http.ResponseWriter, req *http.Request) {
	var form forms.SignUpForm

	form, err := utils.Decode[forms.SignUpForm](req)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign in")
		err := utils.Encode(res, req, http.StatusBadRequest, utils.ResponseError{Message: err})
		if err != nil {
			return
		}
	}

	tokens, err := h.AuthServices.Login(context.Background(), form)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign in")
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	cookieRefresh := http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(res, &cookie)

	http.SetCookie(res, &cookieRefresh)

	err = utils.Encode(res, req, http.StatusOK, utils.Response{Message: "successfully logged in"})

	return
}

func (h *AuthHandler) HandleForgotPassword(res http.ResponseWriter, req *http.Request) {
	return
}

func (h *AuthHandler) HandleRefresh(res http.ResponseWriter, req *http.Request) {
	return
}

func (h *AuthHandler) HandleSignOut(res http.ResponseWriter, req *http.Request) {
	return
}
