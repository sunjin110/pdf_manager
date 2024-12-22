package usecase

import (
	"context"
	"fmt"

	"github.com/sunjin110/pdf_manager/core/domain/model"
	"github.com/sunjin110/pdf_manager/core/domain/repository"
)

type Password interface {
	RegistPasswords(ctx context.Context, passwords model.Passwords) error
	DeleteAllPasswords(ctx context.Context) error
	GetAllPasswords(ctx context.Context) (model.Passwords, error)
}

type password struct {
	passwordRepo repository.Password
}

func NewPassword(passwordRepo repository.Password) Password {
	return &password{
		passwordRepo: passwordRepo,
	}
}

func (p *password) GetAllPasswords(ctx context.Context) (model.Passwords, error) {
	return p.passwordRepo.GetAll(ctx)
}

func (p *password) DeleteAllPasswords(ctx context.Context) error {
	return p.passwordRepo.DeleteAll(ctx)
}

func (p *password) RegistPasswords(ctx context.Context, passwords model.Passwords) error {
	for _, password := range passwords {
		if err := p.passwordRepo.Insert(ctx, password); err != nil {
			return fmt.Errorf("failed insert password. err: %w", err)
		}
	}
	return nil
}
