package file

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"file_storage_service/internal/models"
)

// GetAllFiles will get all the files in Database
//
// Return list of models.File and nil error when succeed.
// Otherwise, will return empty list of models.File and non-nil error.
func (usecase Usecase) GetAllFiles() ([]models.File, error) {
	// TODO: span the context

	resp, err := usecase.fileDB.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

// UploadFile will upload given file to local storage
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (usecase Usecase) UploadFile(file multipart.File) error {
	// TODO: span the context

	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		log.Println(err)
		return err
	}

	filename := "IMAGE-" + hex.EncodeToString(uuid)
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	fileLocation := filepath.Join(dir, "files", filename)

	err = usecase.file.UploadFile(file, fileLocation)
	if err != nil {
		log.Println(err)
		return err
	}

	// TODO: remove this harcoded username
	err = usecase.fileDB.SetFile(fileLocation, "david123")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}