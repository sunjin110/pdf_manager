package core

import (
	"context"
	"fmt"
	"io"

	"github.com/sunjin110/pdf_manager/core/domain/model"
	"github.com/sunjin110/pdf_manager/core/infrastructure/repository"
	"github.com/sunjin110/pdf_manager/core/infrastructure/sqlite"
	"github.com/sunjin110/pdf_manager/core/usecase"
)

type Core interface {
	RegistPasswordByCSV(csvReader io.Reader) error
	GetAllPasswords(ctx context.Context) (model.Passwords, error)
	GetPasswordsByTargetNames(ctx context.Context, targetNames []string) (model.Passwords, error)
	ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error
}

type core struct {
	securityUsecase usecase.Security
	passwordUsecase usecase.Password
}

func NewCore(dbPath string) (Core, error) {
	if err := sqlite.Migrate(dbPath); err != nil {
		return nil, fmt.Errorf("failed migrate sqlite. err: %w", err)
	}

	db, err := sqlite.NewSQLiteDB(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed new sqlite. err: %w", err)
	}

	securityRepo := repository.NewSecurity()
	passwordRepo := repository.NewPassword(db)

	return &core{
		securityUsecase: usecase.NewSecurity(securityRepo),
		passwordUsecase: usecase.NewPassword(passwordRepo),
	}, nil
}

func (c *core) RegistPasswordByCSV(csvReader io.Reader) error {
	return c.passwordUsecase.RegistPasswordsByCSV(context.Background(), csvReader)
}

func (c *core) GetAllPasswords(ctx context.Context) (model.Passwords, error) {
	return c.passwordUsecase.GetAllPasswords(ctx)
}

func (c *core) GetPasswordsByTargetNames(ctx context.Context, targetNames []string) (model.Passwords, error) {
	return c.passwordUsecase.GetPasswordsByTargetName(ctx, targetNames)
}

func (c *core) ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error {
	return c.securityUsecase.ProtectPDF(input, output, userPassword, ownerPassword)
}
