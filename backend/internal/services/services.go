package services

import (
	"backend/internal/mail"
	"backend/internal/repositories"
)

func NewAuthServices(authRepo *repositories.AuthRepository, mailClient *mail.MailClient) *AuthServices {
	return &AuthServices{
		repo:       authRepo,
		mailClient: mailClient,
	}
}
