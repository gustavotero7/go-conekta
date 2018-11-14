package models

type Order struct {
	ID              string           `json:"id,omitempty"`
	Object          string           `json:"object,omitempty"`
	CreatedAt       int64            `json:"created_at,omitempty"`
	UpdatedAt       int64            `json:"updated_at,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	LineItems       []LineItem       `json:"line_items,omitempty"`
	ShippingLines   []ShippingLine   `json:"shipping_lines,omitempty"`
	TaxLines        []TaxLine        `json:"tax_lines,omitempty"`
	DiscountLines   []DiscountLine   `json:"discount_lines,omitempty"`
	Livemode        *bool            `json:"livemode,omitempty"`
	PreAuthorize    *bool            `json:"pre_authorize,omitempty"`
	ShippingContact *ShippingContact `json:"shipping_contact,omitempty"`
	Amunt           float64          `json:"amount,omitempty"`
	Reason          string           `json:"reason,omitempty"`
	AmountRefunded  float64          `json:"amount_refunded,omitempty"`
	PaymentStatus   string           `json:"payment_status,omitempty"`
	CustomerInfo    Customer         `json:"customer_info,omitempty"`
	Charges         []Charge         `json:"charges,omitempty"`
	Metadata        Metadata         `json:"metadata,omitempty"`
}
