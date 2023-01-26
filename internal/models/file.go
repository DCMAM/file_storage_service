package models

// File represents the entity of file
type File struct {
	FileID     int64  `json:"file_id"`
	URL        string `json:"url"`
	Uploader   string `json:"uploader"`
	UploadTime string `json:"upload_time"`
}
