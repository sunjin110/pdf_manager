package core

import (
	"fmt"

	"github.com/sunjin110/pdf_manager/core/infrastructure/repository"
	"github.com/sunjin110/pdf_manager/core/infrastructure/sqlite"
	"github.com/sunjin110/pdf_manager/core/usecase"
)

type Core interface {
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
