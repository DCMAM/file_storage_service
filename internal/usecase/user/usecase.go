package user

import (
	"file_storage_service/internal/models"
)

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user

// UserProvider provides user repository's methods for user usecase
type UserProvider interface {
	// GetUserByUsername will get user based on the given username parameter.
	//
	// Return models.User and nil error when succeed.
	// Otherwise, will return empty row of user and non-nil error.
	GetUserByUsername(username string) (models.User, error)

	// RegisterUser will register user based on the given name, username, password parameter.
	//
	// Return models.User and nil error when succeed.
	// Otherwise, will return empty row of user and non-nil error.
	RegisterUser(name string, username string, password string) error
}

// Usecase represents all the needed repositories for users
type Usecase struct {
	user UserProvider
}

// NewUsecase will initiate user's repositories
func NewUsecase(user UserProvider) Usecase {
	return Usecase{
		user: user,
	}
}
