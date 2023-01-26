package local_storage

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

// UploadFile will will upload file and save it to local storage
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (a Repository) UploadFile(file multipart.File, fileLocation string) error {

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
