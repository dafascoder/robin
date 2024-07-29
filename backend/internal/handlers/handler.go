package handlers

import "backend/internal/services"

func NewAuthHandler(s *services.AuthServices) *AuthHandler {
	return &AuthHandler{
		AuthServices: s,
	}
}
