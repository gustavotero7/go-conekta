package models

type AntifraudInfo map[string]string
type Tags map[string]interface{}
type LineItem struct {
	ID            string        `json:"id,omitempty"`
	Object        string        `json:"object,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	UnitPrice     int64         `json:"unit_price,omitempty"`
	Quantity      int64         `json:"quantity,omitempty"`
	Sku           string        `json:"sku,omitempty"`
	Tags          Tags          `json:"tags,omitempty"`
	Brand         string        `json:"brand,omitempty"`
	ParentID      string        `json:"parent_id,omitempty"`
	Metadata      Metadata      `json:"metadata,omitempty"`
	AntifraudInfo AntifraudInfo `json:"antifraud_info,omitempty"`
}
type LineItems struct {
	Object  string     `json:"object,omitempty"`
	HasMore *bool      `json:"has_more,omitempty"`
	Total   int64      `json:"total,omitempty"`
	Data    []LineItem `json:"data,omitempty"`
}
