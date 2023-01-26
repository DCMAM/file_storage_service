package user

import (
	"errors"
	"log"

	jwtHelper "file_storage_service/internal/helper/jwt"
	"file_storage_service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

const (
	errAlreadyExist = "User already exist"
	errNotMatch     = "username didn't match with any data in database"
	errPassNotMatch = "password didn't match"
)

// Login will handle the login request based on the given username and password
//
// Return models.User and nil error when succeed.
// Otherwise, will return empty row of user and non-nil error.
func (usecase Usecase) Login(username string, password string) (models.User, error) {
	// TODO: span the context

	resp, err := usecase.user.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	if resp.UserID == 0 {
		log.Println(err)
		return models.User{}, errors.New(errNotMatch)
	}

	// password check
	if err := bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(password)); err != nil {
		log.Println(err)
		return models.User{}, errors.New(errPassNotMatch)
	}

	resp.Token, err = jwtHelper.GenerateToken(resp.Username)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	resp.Password = ""
	return resp, nil
}

// Register will handle the register request based on the given name, username and password
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (usecase Usecase) Register(name string, username string, password string) error {
	// TODO: span the context

	resp, err := usecase.user.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		return err
	}

	if resp.Username == username {
		log.Println(resp.UserID)
		return errors.New(errAlreadyExist)
	}

	// encrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	err = usecase.user.RegisterUser(name, username, string(hashPassword))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
