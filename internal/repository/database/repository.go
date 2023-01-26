package database

import (
	"database/sql"
)

// Repository represents all the needed depedencies for users
type Repository struct {
	db *sql.DB
}

// NewRepository will initiate UserRepository's provider
func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}
