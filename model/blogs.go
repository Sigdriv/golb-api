package model

import "time"

type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags"`
	Views     string    `json:"views"`
}
