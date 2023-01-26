package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"file_storage_service/internal/handler"
	"file_storage_service/internal/models"
)

const (
	// error message
	errEmptyName     = "name is required"
	errEmptyUsername = "username is required"
	errEmptyEmail    = "email is required"
	errEmptyPassword = "password is required"
)

// Login will handle the login request based on the given username and password
func (a Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user models.User
	)

	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("json.Unmarshal() got error - HNDL.L00 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" {
		err := errors.New(errEmptyUsername)
		log.Printf("empty name - HNDL.L01 - %s\n", err.Error())
		handler.Response(w, errEmptyName, http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		err := errors.New(errEmptyPassword)
		log.Printf("empty password - HNDL.L02 - %s\n", err.Error())
		handler.Response(w, errEmptyPassword, http.StatusBadRequest)
		return
	}

	resp, err := a.user.Login(user.Username, user.Password)
	if err != nil {
		log.Printf("a.user.Login() got error - HNDL.L03\n")
		handler.Response(w, errEmptyPassword, http.StatusInternalServerError)
		return
	}

	// set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    resp.Token,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Path:     "/",
		Value:    user.Username,
		HttpOnly: true,
	})

	handler.Response(w, resp, http.StatusOK)
}

// Register will handle the register request based on the given name, username and password
func (a Handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user models.User
	)

	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("json.Unmarshal() got error - HNDL.R00 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := user.Name
	if name == "" {
		err := errors.New(errEmptyName)
		log.Printf("empty name - HNDL.R01 - %s\n", err.Error())
		handler.Response(w, errEmptyName, http.StatusBadRequest)
		return
	}

	username := user.Username
	if username == "" {
		err := errors.New(errEmptyUsername)
		log.Printf("empty username - HNDL.R02 - %s\n", err.Error())
		handler.Response(w, errEmptyEmail, http.StatusBadRequest)
		return
	}

	password := user.Password
	if password == "" {
		err := errors.New(errEmptyPassword)
		log.Printf("empty password - HNDL.R03 - %s\n", err.Error())
		handler.Response(w, errEmptyPassword, http.StatusBadRequest)
		return
	}

	err = a.user.Register(name, username, password)
	if err != nil {
		log.Printf("empty password - HNDL.R04 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.Response(w, "success", 200)
}
