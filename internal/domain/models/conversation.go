package models

import "time"

type Conversation struct {
	ID        string    `json:"id"`
	IsGroup   bool      `json:"is_group"`
	Name      *string   `json:"name"`       // NULLABLE
	CreatedBy *string   `json:"created_by"` // NULLABLE
	CreatedAt time.Time `json:"created_at"`
}
