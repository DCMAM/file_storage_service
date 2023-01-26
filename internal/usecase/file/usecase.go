package file

import (
	"file_storage_service/internal/models"
	"mime/multipart"
)

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=file

// fileProvider provides repository's methods for file usecase in local storage repo
type fileProvider interface {
	// UploadFile will will upload file and save it to local storage
	//
	// Return nil error when succeed.
	// Otherwise, will return non-nil error.
	UploadFile(file multipart.File, fileLocation string) error
}

// fileDBProvider provides repository's methods for file usecase
type fileDBProvider interface {
	// GetAll will get all the file from database.
	//
	// Return models.File and nil error when succeed.
	// Otherwise, will return empty row of file and non-nil error.
	GetAll() ([]models.File, error)

	// SetFile will set and save file on DB based on the given url path and username parameter.
	//
	// Return nil error when succeed.
	// Otherwise, will return non-nil error.
	SetFile(url string, username string) error
}

// Usecase represents all the needed repositories for files
type Usecase struct {
	file   fileProvider
	fileDB fileDBProvider
}

// NewUsecase will initiate file's repositories
func NewUsecase(file fileProvider, fileDB fileDBProvider) Usecase {
	return Usecase{
		file:   file,
		fileDB: fileDB,
	}
}
