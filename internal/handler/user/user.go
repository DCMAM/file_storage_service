package user

import (
	"encoding/json"
	"file_storage_service/internal/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// error message
	errEmptyName     = "name is required"
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
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := user.Username
	if username == "" {
		_, _ = w.Write([]byte(errEmptyEmail))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := user.Password
	if password == "" {
		_, _ = w.Write([]byte(errEmptyPassword))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := a.user.Login(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body, _ = json.Marshal(map[string]interface{}{
			"message": err,
			"code":    http.StatusInternalServerError,
		})
		_, _ = w.Write(body)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    resp.Token,
		HttpOnly: true,
	})

	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
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
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := user.Name
	if name == "" {
		_, _ = w.Write([]byte(errEmptyName))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := user.Username
	if username == "" {
		_, _ = w.Write([]byte(errEmptyEmail))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := user.Password
	if password == "" {
		_, _ = w.Write([]byte(errEmptyPassword))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.user.Register(name, username, password)
	if err != nil {
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
		body, _ = json.Marshal(map[string]interface{}{
			"message": err,
			"code":    http.StatusInternalServerError,
		})
		_, _ = w.Write(body)
		return
	}

	body, _ = json.Marshal(map[string]interface{}{
		"message": "success",
		"code":    200,
	})
	_, _ = w.Write(body)
}
