package models

type Metadata map[string]interface{}
type ShippingContact struct {
	ID             string   `json:"id,omitempty"`
	Object         string   `json:"object,omitempty"`
	CreatedAt      int64    `json:"created_at,omitempty"`
	UpdatedAt      int64    `json:"updated_at,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	Receiver       string   `json:"receiver,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Address        Address  `json:"address,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}
