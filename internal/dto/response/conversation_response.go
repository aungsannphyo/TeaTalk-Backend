package response

import (
	"time"

	"github.com/aungsannphyo/ywartalk/internal/domain/models"
)

type Conversation struct {
	ID        string    `json:"id"`
	IsGroup   bool      `json:"is_group"`
	Name      *string   `json:"name"`       // NULLABLE
	CreatedBy *string   `json:"created_by"` // NULLABLE
	CreatedAt time.Time `json:"created_at"`
}

func NewConversationResponse(c *models.Conversation) *Conversation {
	return &Conversation{
		ID:        c.ID,
		IsGroup:   c.IsGroup,
		Name:      c.Name,
		CreatedBy: c.CreatedBy,
		CreatedAt: c.CreatedAt,
	}
}
