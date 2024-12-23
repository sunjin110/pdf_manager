package usecase

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/sunjin110/pdf_manager/core/domain/model"
	"github.com/sunjin110/pdf_manager/core/domain/repository"
)

type Password interface {
	RegistPasswords(ctx context.Context, passwords model.Passwords) error
	DeleteAllPasswords(ctx context.Context) error
	GetAllPasswords(ctx context.Context) (model.Passwords, error)
	RegistPasswordsByCSV(ctx context.Context, csvReader io.Reader) error
	GetPasswordsByTargetName(ctx context.Context, targetNames []string) (model.Passwords, error)
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

func (p *password) RegistPasswordsByCSV(ctx context.Context, csvReader io.Reader) error {
	lines, err := csv.NewReader(csvReader).ReadAll()
	if err != nil {
		return fmt.Errorf("failed new csv reader. err: %w", err)
	}
	passwords := model.Passwords{}
	for _, line := range lines {
		if len(line) < 2 {
			return fmt.Errorf("invalid csv. columns must be 2")
		}

		u, err := uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("failed generate uuid. err: %w", err)
		}

		passwords = append(passwords, model.Password{
			ID:         u.String(),
			TargetName: line[0],
			Password:   line[1],
		})
	}

	// 全部消して
	if err := p.DeleteAllPasswords(ctx); err != nil {
		return fmt.Errorf("failed delete all. err: %w", err)
	}

	// 全部作り直す
	if err := p.RegistPasswords(ctx, passwords); err != nil {
		return fmt.Errorf("failed regist password. err: %w", err)
	}
	return nil
}

func (p *password) GetPasswordsByTargetName(ctx context.Context, targetNames []string) (model.Passwords, error) {
	return p.passwordRepo.FindByTargetNames(ctx, targetNames)
}
