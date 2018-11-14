package models

type Metadata map[string]interface{}
type ConektaBase struct {
	Livemode  *bool    `json:"livemode,omitempty"`
	Object    string   `json:"object,omitempty"`
	Metadata  Metadata `json:"metadata,omitempty,omitempty"`
	CreatedAt int64    `json:"created_at,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
}
