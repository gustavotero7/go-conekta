package models

type DiscountLine struct {
	ConektaBase
	ID       string `json:"id,omitempty"`
	Code     string `json:"code,omitempty"`
	Type     string `json:"type,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
}
