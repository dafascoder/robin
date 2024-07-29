package handlers

import (
	"backend/internal/forms"
	"backend/internal/services"
	"backend/internal/utils"
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
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
		err := utils.Encode(res, req, http.StatusBadRequest, err)
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

	return
}

func (h *AuthHandler) HandleSignIn(res http.ResponseWriter, req *http.Request) {
	return
}

func (h *AuthHandler) HandleSignOut(res http.ResponseWriter, req *http.Request) {
	return
}
