package entities

import (
	"cat-social/models/enum"

	"github.com/google/uuid"
)

type Cat struct {
	CatID       uuid.UUID `json:"cat_id" db:"cat_id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Race        enum.Race `json:"race" db:"race"`
	Sex         enum.Sex  `json:"sex" db:"sex"`
	AgeInMonth  int       `json:"ageInMonth" db:"age_in_month"`
	Description string    `json:"description" db:"description"`
	HasMatched  bool      `json:"hasMatched" db:"has_matched"`
	ImageURLs   []string  `json:"imageUrls" db:"image_urls"`
	Common
}
