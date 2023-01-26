package file

import (
	"file_storage_service/internal/models"
	"os"

	"mime/multipart"
)

//go:generate mockgen -source=./handler.go -destination=./handler_mock.go -package=file

// fileProvicer provides file usecase's methods for user handler
type fileProvicer interface {
	// DonwloadFile will download file based on the path parameter
	//
	// Return nil error when succeed.
	// Otherwise, will return non-nil error.
	DonwloadFile(path string) (*os.File, error)

	// GetAllFiles will get all the files in Database
	//
	// Return list of models.File and nil error when succeed.
	// Otherwise, will return empty list of models.File and non-nil error.
	GetAllFiles() ([]models.File, error)

	// UploadFile will upload given file to local storage
	//
	// Return nil error when succeed.
	// Otherwise, will return non-nil error.
	UploadFile(file multipart.File) error
}

// Handler represents all the needed usecase for users
type Handler struct {
	file fileProvicer
}

// NewHandler will initiate user's usecase
func NewHandler(file fileProvicer) Handler {
	return Handler{
		file: file,
	}
}
