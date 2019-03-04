package models

type Order struct {
	ConektaBase
	ID              string           `json:"id,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	LineItems       []LineItem       `json:"line_items,omitempty"`
	ShippingLines   []ShippingLine   `json:"shipping_lines,omitempty"`
	TaxLines        []TaxLine        `json:"tax_lines,omitempty"`
	DiscountLines   []DiscountLine   `json:"discount_lines,omitempty"`
	PreAuthorize    *bool            `json:"pre_authorize,omitempty"`
	ShippingContact *ShippingContact `json:"shipping_contact,omitempty"`
	Amount          float64          `json:"amount,omitempty"`
	Reason          string           `json:"reason,omitempty"`
	AmountRefunded  float64          `json:"amount_refunded,omitempty"`
	PaymentStatus   string           `json:"payment_status,omitempty"`
	CustomerInfo    Customer         `json:"customer_info,omitempty"`
	Charges         []Charge         `json:"charges,omitempty"`
}

type OrderResponse struct {
	ConektaBase
	Order
	LineItems     LineItems     `json:"line_items,omitempty"`     // Overwrite Order.LineItems
	Charges       Charges       `json:"charges,omitempty"`        // Overwrite Order.Charges
	ShippingLines ShippingLines `json:"shipping_lines,omitempty"` // Overwrite Order.ShippingLines
}
