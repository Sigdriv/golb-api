package db

import "time"

type Blog struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Author    string    `db:"authorName"`
	CreatedAt time.Time `db:"created_at"`
	Tag       string    `db:"tagName"`
	Views     string    `db:"views"`
}

type Author struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	// Email string `db:"email"`
}

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
