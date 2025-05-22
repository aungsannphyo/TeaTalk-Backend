package models

import "time"

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

type PersonalDetails struct {
	UserID       string    `json:"user_id" db:"user_id"`
	ProfileImage *string   `json:"profile_image,omitempty" db:"profile_image"` // Nullable
	Gender       *Gender   `json:"gender,omitempty" db:"gender"`               // Nullable ENUM
	DateOfBirth  *string   `json:"date_of_birth,omitempty" db:"date_of_birth"` // Nullable DATE
	Bio          *string   `json:"bio,omitempty" db:"bio"`                     // Nullable TEXT
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	IsOnline     bool      `json:"isOnline"`
	LastSeen     time.Time `json:"lastSeen"`
}
