package db

import (
	"fmt"
	"golb-api/model"
	"strconv"
)

func (db *DB) RegisterViewToDB(blogID string, body model.Statistics) (message string, err error) {
	checkQuery := ` 
		SELECT id
		FROM statistics
		WHERE uuid = :uuid AND blog_id = :blogID
		`

	blogIDInt, err := strconv.Atoi(blogID)
	if err != nil {
		err = fmt.Errorf("invalid blog ID >> %w", err)
		return
	}

	checkArgs := map[string]any{
		"uuid":   body.Uuid,
		"blogID": blogIDInt,
	}

	res, err := Query[Statistics](db, checkQuery, checkArgs)
	if err != nil {
		err = fmt.Errorf("error checking existing statistics >> %w", err)
		return
	}

	if len(res) > 0 {
		updateQuery := `
			UPDATE statistics
			SET ended_at = CURRENT_TIMESTAMP
			WHERE uuid = :uuid AND blog_id = :blogID
		`
		updateArgs := map[string]any{
			"uuid":   body.Uuid,
			"blogID": blogIDInt,
		}

		_, err = Exec(db, updateQuery, updateArgs)
		if err != nil {
			err = fmt.Errorf("error updating statistics >> %w", err)
			return
		}

		message = "updated"

		return
	}

	insertNewQuery := `
		INSERT INTO statistics (blog_id, uuid)
		VALUES (:blogID, :uuid)
	`
	insertNewArgs := map[string]any{
		"uuid":   body.Uuid,
		"blogID": blogIDInt,
	}

	_, err = Exec(db, insertNewQuery, insertNewArgs)
	if err != nil {
		err = fmt.Errorf("error inserting new statistics >> %w", err)
		return
	}

	return
}
