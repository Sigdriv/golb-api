package db

import (
	"encoding/base64"
	"fmt"
)

func (db *DB) UploadBlogPhoto(blogID string, file []byte) (err error) {
	query := `INSERT INTO file (blog_id, data) VALUES (:blogID, :file)`

	args := map[string]any{
		"blogID": blogID,
		"file":   base64.StdEncoding.EncodeToString(file),
	}

	_, err = Exec(db, query, args)
	if err != nil {
		err = fmt.Errorf("error uploading blog photo >> %w", err)
		return
	}

	return
}
