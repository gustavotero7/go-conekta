package models

type ShippingLine struct {
	ConektaBase
	ID             string `json:"id,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	Method         string `json:"method,omitempty"`
	ParentID       string `json:"parent_id,omitempty"`
}

type ShippingLines struct {
	ConektaBase
	HasMore *bool          `json:"has_more,omitempty"`
	Total   int64          `json:"total,omitempty"`
	Data    []ShippingLine `json:"data,omitempty"`
}
