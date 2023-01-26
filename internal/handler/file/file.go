package file

import (
	"encoding/json"
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
		fmt.Println(err)
		_, _ = w.Write([]byte(errInvalidBody))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if path.Path == "" {
		_, _ = w.Write([]byte(errEmptyPath))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := h.file.DonwloadFile(path.Path)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(map[string]interface{}{
		"message": "success",
		"code":    200,
		"body":    resp,
	})
	_, _ = w.Write(body)
}

// UploadFile will handle the upload file request based on the given writer and request.
func (a Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if fileHeader.Size > (uploadDocumentMaxMemory) {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.file.UploadFile(file)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(map[string]interface{}{
		"message": "success",
		"code":    200,
	})
	_, _ = w.Write(body)
}
