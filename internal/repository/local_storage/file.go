package local_storage

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

// DonwloadFile will download file based on the path parameter
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (repo Repository) DonwloadFile(path string) (*os.File, error) {
	var err error

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return f, nil
}

// UploadFile will will upload file and save it to local storage
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (repo Repository) UploadFile(file multipart.File, fileLocation string) error {
	var err error

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
