package models

// WebhookBody model
type WebhookBody struct {
	ConektaBase
	ID            string       `json:"id"`
	Type          string       `json:"type"`
	Data          WebhookData  `json:"data"`
	WebhookStatus string       `json:"webhook_status"`
	WebhookLogs   []WebhookLog `json:"webhook_logs"`
}

// WebhookLog info about the webhook
type WebhookLog struct {
	ID                     string `json:"id"`
	URL                    string `json:"url"`
	FailedAttempts         int    `json:"failed_attempts"`
	LastHTTPResponseStatus int    `json:"last_http_response_status"`
	Object                 string `json:"object"`
	LastAttemptedAt        int    `json:"last_attempted_at"`
}

// WebhookData holds hook body data
type WebhookData struct {
	Object             WebhookDataObject         `json:"object"`
	PreviousAttributes WebhookPreviousAttributes `json:"previous_attributes"`
}

// WebhookPreviousAttributes previous hook status
type WebhookPreviousAttributes struct {
	Status string `json:"status"`
}

// WebhookDataObject holds data object
type WebhookDataObject struct {
	ID        string `json:"id"`
	Livemode  bool   `json:"livemode"`
	CreatedAt int    `json:"created_at"`
	Currency  string `json:"currency"`
	// This structure mismatch payment_method model, must be checked
	// TODO define an struct for this
	PaymentMethod interface{}              `json:"payment_method"`
	Details       WebhookDataObjectDetails `json:"details"`
	Object        string                   `json:"object"`
	Status        string                   `json:"status"`
	Amount        int                      `json:"amount"`
	PaidAt        int                      `json:"paid_at"`
	Fee           int                      `json:"fee"`
	CustomerID    string                   `json:"customer_id"`
	OrderID       string                   `json:"order_id"`
}

// WebhookDataObjectDetails holds data object details
type WebhookDataObjectDetails struct {
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Email           string          `json:"email"`
	LineItems       []LineItem      `json:"line_items"`
	ShippingContact ShippingContact `json:"shipping_contact"`
	Object          string          `json:"object"`
}
