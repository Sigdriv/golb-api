package db

import "fmt"

func (db *DB) IncrementBlogViews(blogID string) (err error) {
	query := `UPDATE statistics SET views = views + 1 WHERE blog_id = :blog_id`
	args := map[string]any{
		"blog_id": blogID,
	}

	_, err = Exec(db, query, args)
	if err != nil {
		err = fmt.Errorf("error incrementing blog views >> %w", err)
		return
	}

	return
}

func createStatistics(db *DB, blogID string) (err error) {
	query := `
		INSERT INTO statistics (blog_id)
		VALUES (:blogID)
	`
	args := map[string]any{
		"blogID": blogID,
	}

	_, err = Exec(db, query, args)
	if err != nil {
		err = fmt.Errorf("error inserting statistics >> %w", err)
		return
	}

	return
}
