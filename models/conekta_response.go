package models

type ConektaResponse struct {
	Livemode        *bool           `json:"livemode,omitempty"`
	Amount          int64           `json:"amount,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	PaymentStatus   string          `json:"payment_status,omitempty"`
	AmountRefunded  int64           `json:"amount_refunded,omitempty"`
	CustomerInfo    CustomerInfo    `json:"customer_info,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Object          string          `json:"object,omitempty"`
	ID              string          `json:"id,omitempty"`
	Metadata        Metadata        `json:"metadata,omitempty,omitempty"`
	CreatedAt       int64           `json:"created_at,omitempty"`
	UpdatedAt       int64           `json:"updated_at,omitempty"`
	LineItems       LineItems       `json:"line_items,omitempty"`
	ShippingLines   ShippingLines   `json:"shipping_lines,omitempty"`
	Charges         Charges         `json:"charges,omitempty"`
}
