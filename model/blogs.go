package model

import "time"

type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"publishedAt"`
	Tags      []string  `json:"tags"`
	Views     string    `json:"views"`
	File      *string   `json:"file,omitempty"`
}
