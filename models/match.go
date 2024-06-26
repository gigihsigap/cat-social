package model

type Match struct {
	ID         int    `json:"id"`
	MatchCatID int    `json:"match_cat_id"`
	UserCatID  int    `json:"user_cat_id"`
	IssuedBy   int    `json:"issued_by"`
	AcceptedBy int    `json:"accepted_by"`
	Message    string `json:"message"`
	IsAproved  bool   `json:"is_approved"`
	IsMatched  bool   `json:"is_matched"`
	Common
}
