package models

import "time"

type Post struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
