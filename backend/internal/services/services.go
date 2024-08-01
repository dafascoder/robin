package services

import (
	"backend/internal/datastore"
	"backend/internal/mail"
	"backend/internal/repositories"
)

func NewAuthServices(authRepo *repositories.AuthRepository, mailClient *mail.MailClient, redis *datastore.RedisStore) *AuthServices {
	return &AuthServices{
		repo:       authRepo,
		mailClient: mailClient,
		redis:      redis,
	}
}
