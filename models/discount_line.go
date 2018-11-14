package models

type DiscountLine struct {
	ID       string   `json:"id,omitempty"`
	Object   string   `json:"object,omitempty"`
	Code     string   `json:"code,omitempty"`
	Type     string   `json:"type,omitempty"`
	Amount   int64    `json:"amount,omitempty"`
	ParentID string   `json:"parent_id,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}
