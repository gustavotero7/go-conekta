package models

type Customer struct {
	CustomerID               string            `json:"customer_id,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Phone                    string            `json:"phone,omitempty"`
	Email                    string            `json:"email,omitempty"`
	Corporate                bool              `json:"corporate,omitempty"`
	DefaultPaymentSourceID   string            `json:"default_payment_source_id,omitempty"`
	DefaultShippingContactID string            `json:"default_shipping_contact_id,omitempty"`
	PaymentSources           []PaymentSource   `json:"payment_sources,omitempty"`
	ShippingContacts         []ShippingContact `json:"shipping_contacts,omitempty"`
}

type CustomerInfo struct {
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	Corporate  *bool  `json:"corporate,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Object     string `json:"object,omitempty"`
}
