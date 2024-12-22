package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(dbPath string) (*sqlx.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed open sqliteDB, dbPath: %s, err: %w", dbPath, err)
	}

	return sqlx.NewDb(db, "sqlite3"), nil
}
