package file

import (
	"encoding/json"
	"errors"
	"file_storage_service/internal/handler"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// max size
	uploadDocumentMaxMemory = 5242880 // 5 << 20

	// error message
	errInvalidBody = "invalid body"
	errEmptyPath   = "path is required"
	errMaxSize     = "file is too large"
)

// DonwloadFile will handle the download file request based on the given writer and request.
func (h Handler) DonwloadFile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	body, _ := ioutil.ReadAll(r.Body)

	type filePath struct {
		Path string `json:"path"`
	}

	var path filePath

	err = json.Unmarshal(body, &path)
	if err != nil {
		log.Printf("empty password - HNDL.DF00 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	if path.Path == "" {
		err := errors.New(errEmptyPath)
		log.Printf("empty path - HNDL.DF01 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := h.file.DonwloadFile(path.Path)
	if err != nil {
		log.Printf("h.file.DonwloadFile() - HNDL.DF02 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if f != nil {
		defer f.Close()
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Name()))

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllFiles will handle the get all files request based on the given writer and request.
func (a Handler) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	resp, err := a.file.GetAllFiles()
	if err != nil {
		log.Printf("empty password - HNDL.GAF00 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.Response(w, resp, http.StatusOK)
}

// UploadFile will handle the upload file request based on the given writer and request.
func (a Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Printf("empty password - HNDL.UF00 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	if fileHeader.Size > (uploadDocumentMaxMemory) {
		err := errors.New(errMaxSize)
		log.Printf("empty password - HNDL.UF01 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpCookie, _ := r.Cookie("username")
	username := httpCookie.Value

	err = a.file.UploadFile(file, username)
	if err != nil {
		log.Printf("empty password - HNDL.R02 - %s\n", err.Error())
		handler.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.Response(w, "success", http.StatusOK)
}
