package database

import (
	"database/sql"
	"file_storage_service/internal/models"
	"log"
	"time"
)

const (
	getAllFilesQuery = `
		SELECT
			file_id,
			url,
			uploader,
			upload_time
		FROM files LIMIT 100
	`

	uploadFileQuery = `
		INSERT INTO files (url, uploader, upload_time)
		VALUES (?, ?, ?)
	`
)

// GetAll will get all the file from database.
//
// Return models.File and nil error when succeed.
// Otherwise, will return empty row of file and non-nil error.
func (a Repository) GetAll() ([]models.File, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = a.db.Query(getAllFilesQuery)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		_ = rows.Scan(&file.FileID, &file.URL, &file.Uploader, &file.UploadTime)
		files = append(files, file)
	}

	return files, nil
}

// SetFile will set and save file on DB based on the given url path and username parameter.
//
// Return nil error when succeed.
// Otherwise, will return non-nil error.
func (a Repository) SetFile(url string, username string) error {
	var (
		err error
	)

	now := time.Now()

	_, err = a.db.Exec(uploadFileQuery, url, username, now)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
