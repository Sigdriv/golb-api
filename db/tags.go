package db

import (
	"fmt"
	"golb-api/model"
)

func (db *DB) GetTags() (tags []model.Tag, err error) {
	query := `SELECT id, name FROM tag`

	newTags, err := Query[Tag](db, query, nil)
	if err != nil {
		err = fmt.Errorf("error getting tags >> %w", err)
		return
	}

	tags = mapTags(newTags)

	return
}

func mapTags(dbTags []Tag) (tags []model.Tag) {
	for _, dbTag := range dbTags {
		tag := model.Tag{
			ID:   dbTag.ID,
			Name: dbTag.Name,
		}
		tags = append(tags, tag)
	}

	return
}
