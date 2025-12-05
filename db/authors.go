package db

import (
	"fmt"
	"golb-api/model"
)

func (db *DB) GetAuthors() (authors []model.Author, err error) {
	query := `SELECT id, name FROM author`

	dbAuthors, err := Query[Author](db, query, nil)
	if err != nil {
		err = fmt.Errorf("error getting authors >> %w", err)
		return
	}

	authors = mapAuthors(dbAuthors)

	return
}

func mapAuthors(dbAuthors []Author) (authors []model.Author) {
	for _, dbAuthor := range dbAuthors {
		author := model.Author{
			ID:   dbAuthor.ID,
			Name: dbAuthor.Name,
		}
		authors = append(authors, author)
	}
	return
}
