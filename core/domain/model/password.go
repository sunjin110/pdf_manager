package model

type Passwords []Password

type Password struct {
	ID         string // unique_id
	TargetName string
	Password   string
}
