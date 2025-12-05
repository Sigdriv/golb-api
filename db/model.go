package db

import "time"

type Blog struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Author    string    `db:"authorName"`
	CreatedAt time.Time `db:"created_at"`
	Tag       string    `db:"tagName"`
	Views     string    `db:"views"`
	File      *string   `db:"file,omitempty"`
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

type Statistics struct {
	ID        string    `db:"id"`
	BlogID    string    `db:"blog_id"`
	StartedAt time.Time `db:"started_at"`
	EndedAt   time.Time `db:"ended_at"`
	Uuid      string    `db:"uuid"`
}
