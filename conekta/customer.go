package conekta

import (
	"encoding/json"

	"github.com/gustavotero7/go-conekta/models"
)

type Customer struct {
	models.Customer
}

// Create Creates a new Customer.
func (c *Customer) Create() (*models.ConektaResponse, error) {
	response, err := request("POST", "/customers", c)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Update Updates an existing Customer.
func (c *Customer) Update() (*models.ConektaResponse, error) {
	response, err := request("PUT", "/customers/"+c.CustomerID, c)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Delete Deletes an existing customer.
func (c *Customer) Delete() (*models.ConektaResponse, error) {
	response, err := request("DELETE", "/customers/"+c.CustomerID, nil)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {

	}
	return &cr, nil
}

// CreateSubscription Creates a new subscription using tokenized data.
func (c *Customer) CreateSubscription(plan string) (*models.Subscription, error) {
	response, err := request("POST", "/customers/"+c.CustomerID+"/subscription", models.Body{"plan": plan})
	if err != nil {
		return nil, err
	}

	var s models.Subscription
	if err := json.Unmarshal(response, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// UpdateSubscription Updates a subscription with a different card or plan.
func (c *Customer) UpdateSubscription(plan string) (*models.Subscription, error) {
	response, err := request("PUT", "/customers/"+c.CustomerID+"/subscription", models.Body{"plan": plan})
	if err != nil {
		return nil, err
	}

	var s models.Subscription
	if err := json.Unmarshal(response, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// PauseSubscription Pauses a subscription.
func (c *Customer) PauseSubscription() (*models.Subscription, error) {
	response, err := request("POST", "/customers/"+c.CustomerID+"/subscription/pause", nil)
	if err != nil {
		return nil, err
	}

	var s models.Subscription
	if err := json.Unmarshal(response, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// ResumeSubscription Resumes a subscription.
func (c *Customer) ResumeSubscription() (*models.Subscription, error) {
	response, err := request("POST", "/customers/"+c.CustomerID+"/subscription/resume", nil)
	if err != nil {
		return nil, err
	}

	var s models.Subscription
	if err := json.Unmarshal(response, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// CancelSubscription Cancels a subscription.
func (c *Customer) CancelSubscription() (*models.Subscription, error) {
	response, err := request("POST", "/customers/"+c.CustomerID+"/subscription/cancel", nil)
	if err != nil {
		return nil, err
	}

	var s models.Subscription
	if err := json.Unmarshal(response, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// CreatePaymentSource Creates a new Payment Source
func (c *Customer) CreatePaymentSource(paymentSource models.PaymentSource) (*models.PaymentSource, error) {
	response, err := request("POST", "/customers/"+c.CustomerID+"/payment_sources/", paymentSource)
	if err != nil {
		return nil, err
	}

	var ps models.PaymentSource
	if err := json.Unmarshal(response, &ps); err != nil {
		return nil, err
	}
	return &ps, nil
}

// DeletePaymentSource Deletes an existing Payment Source.
func (c *Customer) DeletePaymentSource(paymentSourceID string) (*models.PaymentSource, error) {
	response, err := request("DELETE", "/customers/"+c.CustomerID+"/payment_sources/"+paymentSourceID, nil)
	if err != nil {
		return nil, err
	}

	var ps models.PaymentSource
	if err := json.Unmarshal(response, &ps); err != nil {
		return nil, err
	}
	return &ps, nil
}
