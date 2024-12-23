package dto

import "github.com/sunjin110/pdf_manager/core/domain/model"

type Passwords []Password

func (passwords Passwords) ToModel() model.Passwords {
	models := make(model.Passwords, 0, len(passwords))
	for _, p := range passwords {
		models = append(models, p.ToModel())
	}
	return models
}

type Password struct {
	ID         string `db:"id"`
	TargetName string `db:"target_name"`
	Password   string `db:"password"`
}

func NewPassword(m model.Password) Password {
	return Password{
		ID:         m.ID,
		TargetName: m.TargetName,
		Password:   m.Password,
	}
}

func (p Password) ToModel() model.Password {
	return model.Password{
		ID:         p.ID,
		TargetName: p.TargetName,
		Password:   p.Password,
	}
}
