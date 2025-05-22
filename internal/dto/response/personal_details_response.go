package response

import "github.com/aungsannphyo/ywartalk/internal/domain/models"

type PersonalDetailsResponse struct {
	UserID       string         `json:"userId"`
	Gender       *models.Gender `json:"gender"`      // Nullable ENUM
	DateOfBirth  *string        `json:"dateOfBirth"` // Nullable DATE
	Bio          *string        `json:"bio"`         // Nullable TEXT
	ProfileImage *string        `json:"profileImage"`
}

func NewPersonalDetailsResponse(ps *models.PersonalDetails) *PersonalDetailsResponse {
	return &PersonalDetailsResponse{
		UserID:       ps.UserID,
		Gender:       ps.Gender,
		DateOfBirth:  ps.DateOfBirth,
		Bio:          ps.Bio,
		ProfileImage: ps.ProfileImage,
	}
}
