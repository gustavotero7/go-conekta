package models

type PaymentSource struct {
	ID        string `json:"id,omitempty"`
	Object    string `json:"object,omitempty"`
	TokenID   string `json:"token_id,omitempty"`
	Type      string `json:"type,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	Last4     string `json:"last4,omitempty"`
	Name      string `json:"name,omitempty"`
	ExpMonth  string `json:"exp_month,omitempty"`
	ExpYear   string `json:"exp_year,omitempty"`
	Brand     string `json:"brand,omitempty"`
	ParentID  string `json:"parent_id,omitempty"`
}
