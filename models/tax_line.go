package models

type TaxLine struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	Description string   `json:"description,omitempty"`
	Amount      float64  `json:"amount,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}
