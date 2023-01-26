package user

import (
	"file_storage_service/internal/models"
)

//go:generate mockgen -source=./handler.go -destination=./handler_mock.go -package=user

// UserProvider provides user repository's methods for user usecase
type UserProvider interface {
	// GetUserByUsername will get user based on the given username parameter.
	//
	// Return models.User and nil error when succeed.
	// Otherwise, will return empty row of user and non-nil error.
	Login(username string, password string) (models.User, error)

	// Register will register user based on the given name, username, password parameter.
	//
	// Return models.User and nil error when succeed.
	// Otherwise, will return empty row of user and non-nil error.
	Register(name string, username string, password string) error
}

// Handler represents all the needed usecase for users
type Handler struct {
	user UserProvider
}

// NewHandler will initiate user's usecase
func NewHandler(user UserProvider) Handler {
	return Handler{
		user: user,
	}
}
