package dto

import (
	"github.com/aungsannphyo/ywartalk/internal/domain/models"
	v "github.com/aungsannphyo/ywartalk/pkg/validator"
)

type PersonalDetailDto struct {
	Gender      *models.Gender `json:"gender"`        // Nullable ENUM
	DateOfBirth *string        `json:"date_of_birth"` // Nullable DATE
	Bio         *string        `json:"bio"`           // Nullable TEXT
}

func ValidateCreatePersonalDetails(pd PersonalDetailDto) error {
	var errs v.ValidationErrors

	if *pd.Gender != models.GenderMale && *pd.Gender != models.GenderFemale && *pd.Gender != models.GenderOther {
		errs = append(errs, v.ValidationError{
			Field:   "Status",
			Message: "Status should be MALE or FEMALE or OTHER",
		})
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
