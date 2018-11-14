package models

type TaxLine struct {
	ConektaBase
	ID          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	ParentID    string  `json:"parent_id,omitempty"`
}
