package services

import (
	"backend/internal/forms"
	"backend/internal/mail"
	model "backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
)

type AuthServices struct {
	repo       *repositories.AuthRepository
	mailClient *mail.MailClient
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

	err = s.mailClient.SendEmail(form.Email, "verify-email.tmpl", "You have successfully created an account")
	if err != nil {
		return model.CreateAccountRow{}, err
	}

	return account, nil
}
