package sqlitequery

import (
	_ "embed"
)

var (
	//go:embed delete_all_password.sql
	DeleteAllPassword string

	//go:embed insert_password.sql
	InsertPassword string
)
