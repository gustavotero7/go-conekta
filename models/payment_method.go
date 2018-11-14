package models

type PaymentMethod struct {
	ConektaBase
	Type                   string  `json:"type,omitempty"`
	TokenId                string  `json:"token_id,omitempty"`
	PaymentSourceID        string  `json:"payment_source_id,omitempty"`
	ServiceName            string  `json:"service_name,omitempty"`
	BarcodeURL             string  `json:"barcode_url,omitempty"`
	ExpiresAt              int64   `json:"expires_at,omitempty"`
	StoreName              string  `json:"store_name,omitempty"`
	Reference              string  `json:"reference,omitempty"`
	Name                   string  `json:"name,omitempty"`
	ExpMonth               string  `json:"exp_month,omitempty"`
	ExpYear                string  `json:"exp_year,omitempty"`
	AuthCode               string  `json:"auth_code,omitempty"`
	Last4                  string  `json:"last4,omitempty"`
	Brand                  string  `json:"brand,omitempty"`
	Issuer                 string  `json:"issuer,omitempty"`
	AccountType            string  `json:"account_type,omitempty"`
	Country                string  `json:"country,omitempty"`
	FraudScore             float64 `json:"fraud_score,omitempty"`
	ReceivingAccountBank   string  `json:"receiving_account_bank,omitempty"`
	ReceivingAccountNumber string  `json:"receiving_account_number,omitempty"`
}
