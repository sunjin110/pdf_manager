package repository

import (
	"context"

	model "github.com/sunjin110/pdf_manager/core/domain/model"
)

type Password interface {
	Insert(ctx context.Context, password model.Password) error
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error
	Update(ctx context.Context, id string, password model.Password) error
	GetAll(ctx context.Context) ([]model.Password, error)
	// Err: NotFound
	GetByTargetName(ctx context.Context, targetName string) (model.Password, error)
}
