package handlers

import "backend/internal/services"

func NewAuthHandler(s *services.AuthServices) *AuthHandler {
	return &AuthHandler{
		AuthServices: s,
	}
}

func NewUserHandler(s *services.UserServices) *UserHandler {
	return &UserHandler{
		UserServices: s,
	}
}
