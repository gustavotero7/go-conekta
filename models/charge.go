package models

type Charge struct {
	ConektaBase
	ID                  string        `json:"id,omitempty"`
	Description         string        `json:"description,omitempty"`
	ExpiresAt           int64         `json:"expires_at,omitempty"`
	Currency            string        `json:"currency,omitempty"`
	Amount              float64       `json:"amount,omitempty"`
	MonthlyInstallments float64       `json:"monthly_installments,omitempty"`
	Status              string        `json:"status,omitempty"`
	Fee                 float64       `json:"fee,omitempty"`
	OrderID             string        `json:"order_id,omitempty"`
	CustomerID          string        `json:"customer_id,omitempty"`
	DeviceFingerprint   string        `json:"device_fingerprint,omitempty"`
	PaidAt              int64         `json:"paid_at,omitempty"`
	PaymentMethod       PaymentMethod `json:"payment_method,omitempty"`
}

type Charges struct {
	ConektaBase
	HasMore *bool    `json:"has_more,omitempty"`
	Total   int64    `json:"total,omitempty"`
	Data    []Charge `json:"data,omitempty"`
}
