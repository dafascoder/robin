package services

import (
	"backend/internal/interfaces"
	"backend/internal/mail"
)

func NewAuthServices(authInterface interfaces.AuthInterface, mailClient *mail.MailClient) *AuthServices {
	return &AuthServices{
		AccountInterface: authInterface,
		mailClient:       mailClient,
	}
}
