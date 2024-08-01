package services

import (
	"backend/internal/config"
	"backend/internal/datastore"
	"backend/internal/forms"
	"backend/internal/helpers"
	"backend/internal/mail"
	model "backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AuthServices struct {
	repo       *repositories.AuthRepository
	mailClient *mail.MailClient
	redis      *datastore.RedisStore
}

func (s *AuthServices) CreateAccount(ctx context.Context, form forms.SignUpForm) (model.CreateAccountRow, error) {
	_, err := s.repo.GetAccountByEmail(ctx, form.Email)

	if err == nil {
		return model.CreateAccountRow{}, errors.New("email already exists")
	}

	newPassword, err := utils.HashPassword(form.Password)
	if err != nil {
		return model.CreateAccountRow{}, errors.New("failed to hash password")
	}

	newAccount := model.CreateAccountParams{
		Email:    form.Email,
		Password: newPassword,
	}

	account, err := s.repo.CreateAccount(ctx, newAccount)
	if err != nil {
		fmt.Print(err)
		return model.CreateAccountRow{}, errors.New("failed to create account")
	}

	// Generate OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return model.CreateAccountRow{}, errors.New("failed to generate OTP")
	}

	// Store OTP in Redis
	err = s.redis.Set(ctx, fmt.Sprintf("%s_%s", activationPrefix, form.Email), otp.Hash, 0)
	if err != nil {
		return model.CreateAccountRow{}, errors.New("failed to store OTP")
	}

	// Send email
	randomString, err := helpers.GenerateRandomString(32)
	if err != nil {
		return model.CreateAccountRow{}, errors.New("failed to generate random string")
	}

	now := time.Now()
	expiration := now.Add(config.Env.TokenExpiration.Duration)
	exact := expiration.Format(time.RFC1123)
	emailData := map[string]interface{}{
		"token":       otp.Secret,
		"email":       form.Email,
		"frontendURL": fmt.Sprintf("http://localhost:3000/verify?email=%s&token=%s", form.Email, randomString),
		"expiration":  config.Env.TokenExpiration.Duration,
		"exact":       exact,
	}

	err = s.mailClient.SendEmail(form.Email, "verify-email.tmpl", emailData)
	if err != nil {
		return model.CreateAccountRow{}, err
	}

	return account, nil
}

func (s *AuthServices) GetAccountByID(ctx context.Context, id uuid.UUID) (model.Account, error) {
	account, err := s.repo.GetAccountByID(ctx, id)
	if err != nil {
		return model.Account{}, errors.New("failed to get account")
	}

	return account, nil
}

var activationPrefix = "activation_"

func (s *AuthServices) VerifyAccount(ctx context.Context, email string, code string) error {
	hash, err := s.redis.Get(ctx, fmt.Sprintf("%s_%s", activationPrefix, email))

	if err != nil {
		return errors.New("failed to verify Redis")
	}

	tokenHash := fmt.Sprintf("%x\n", sha256.Sum256([]byte(code)))

	if tokenHash != *hash {
		return errors.New("token mismatch")
	}

	_, err = s.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		return errors.New("failed to verify account")
	}

	return nil
}
