package file

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	uploadDocumentMaxMemory = 5242880 // 5 << 20
)

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

// v will handle the upload file request based on the given writer and request.
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
