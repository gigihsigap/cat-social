package entities

import (
	"github.com/google/uuid"
	"cat-social/utils/constant"
)

type Cat struct {
	CatID       uuid.UUID        `json:"cat_id" db:"cat_id"`
	UserID      uuid.UUID        `json:"user_id" db:"user_id"`
	Name        string           `json:"name" db:"name"`
	Race        constant.CatRace `json:"race" db:"race"`
	Sex         constant.CatSex  `json:"sex" db:"sex"`
	AgeInMonth  int              `json:"age_in_month" db:"age_in_month"`
	Description string           `json:"description" db:"description"`
	HasMatched  bool             `json:"has_matched" db:"has_matched"`
	ImageURLs   []string         `json:"image_urls" db:"image_urls"`
	Common
}