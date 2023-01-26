package database

import (
	"database/sql"
	"log"

	"file_storage_service/internal/models"
)

const (
	getUserByUsernameQuery = `
		SELECT 
			user_id,
			name,
			username,
			password
		FROM users
		WHERE username = ?
	`
	registerUserQuery = `
		INSERT INTO users (name, username, password)
		VALUES (?, ?, ?)
	`
)

// GetUserByUsername will get user based on the given username parameter.
//
// Return models.User and nil error when succeed.
// Otherwise, will return empty row of user and non-nil error.
func (a Repository) GetUserByUsername(username string) (models.User, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = a.db.Query(getUserByUsernameQuery, username)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User
	for rows.Next() {
		_ = rows.Scan(&user.UserID, &user.Name, &user.Username, &user.Password)
	}

	return user, nil
}

// RegisterUser will register user based on the given name, username, password parameter.
//
// Return models.User and nil error when succeed.
// Otherwise, will return empty row of user and non-nil error.
func (a Repository) RegisterUser(name string, username string, password string) error {
	var (
		err error
	)

	_, err = a.db.Exec(registerUserQuery, name, username, password)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
