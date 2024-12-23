package repository

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunjin110/pdf_manager/core/domain/model"
	"github.com/sunjin110/pdf_manager/core/domain/repository"
	"github.com/sunjin110/pdf_manager/core/infrastructure/repository/dto"
	"github.com/sunjin110/pdf_manager/core/infrastructure/repository/query/sqlitequery"
)

type passwordRepo struct {
	db *sqlx.DB
}

func NewPassword(db *sqlx.DB) repository.Password {
	return &passwordRepo{
		db: db,
	}
}

func (p *passwordRepo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (p *passwordRepo) DeleteAll(ctx context.Context) error {
	if _, err := p.db.ExecContext(ctx, sqlitequery.DeleteAllPassword); err != nil {
		return fmt.Errorf("failed deleteAll. err: %w", err)
	}
	return nil
}

func (p *passwordRepo) GetAll(ctx context.Context) ([]model.Password, error) {
	query, args := sqlbuilder.Select("*").From("passwords").OrderBy("id").Asc().BuildWithFlavor(sqlbuilder.SQLite)
	dtos := dto.Passwords{}
	if err := p.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, fmt.Errorf("failed select context. err: %w", err)
	}
	return dtos.ToModel(), nil
}

func (p *passwordRepo) GetByTargetName(ctx context.Context, targetName string) (model.Password, error) {
	sb := sqlbuilder.Select("*").From("passwords")
	sb.Where(
		sb.Equal("target_name", targetName),
	)
	sb.Limit(1)
	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)

	d := dto.Password{}
	if err := p.db.GetContext(ctx, &d, query, args...); err != nil {

		return model.Password{}, fmt.Errorf("failed get. err: %w", err)
	}
	return d.ToModel(), nil
}

func (p *passwordRepo) FindByTargetNames(ctx context.Context, targetNames []string) (model.Passwords, error) {

	targetNameItfs := make([]any, 0, len(targetNames))
	for _, targetName := range targetNames {
		targetNameItfs = append(targetNameItfs, targetName)
	}

	sb := sqlbuilder.Select("*").From("passwords")
	sb.Where(
		sb.In("target_name", targetNameItfs...),
	)
	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)

	d := dto.Passwords{}
	if err := p.db.SelectContext(ctx, &d, query, args...); err != nil {
		return model.Passwords{}, fmt.Errorf("failed FindByTargetNames. err: %w", err)
	}
	return d.ToModel(), nil
}

func (p *passwordRepo) Insert(ctx context.Context, password model.Password) error {
	stmt, err := p.db.PrepareNamedContext(ctx, sqlitequery.InsertPassword)
	if err != nil {
		return fmt.Errorf("failed PrepareNamedContext. err: %w", err)
	}
	if _, err := stmt.ExecContext(ctx, map[string]any{
		"id":          password.ID,
		"target_name": password.TargetName,
		"password":    password.Password,
	}); err != nil {
		return fmt.Errorf("failed insert password. err: %w", err)
	}
	return nil
}

func (p *passwordRepo) Update(ctx context.Context, id string, password model.Password) error {
	panic("unimplemented")
}
