package customers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/models"
)

const basePath = "/customers"

// Create Creates a new Customer.
func Create(customer models.Customer) (*models.Customer, error) {
	response, err := client.Post(basePath, customer)
	if err != nil {
		return nil, err
	}

	log.Println(string(response))

	var cr models.Customer
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Update Updates an existing Customer.
func Update(customer models.Customer) (*models.Customer, error) {
	response, err := client.Put(fmt.Sprintf("%s/%s", basePath, customer.CustomerID), customer)
	if err != nil {
		return nil, err
	}

	var cr models.Customer
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Delete Deletes an existing customer.
func Delete(customerID string) (*models.Customer, error) {
	response, err := client.Delete(fmt.Sprintf("%s/%s", basePath, customerID))
	if err != nil {
		return nil, err
	}

	var cr models.Customer
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// CreateSubscription Creates a new subscription using tokenized data.
// Status of the subscription. Allowed values are: n_trial, active, past_due, paused, and canceled.
func CreateSubscription(customerID string, plan string) (*models.Subscription, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/subscription", basePath, customerID), map[string]string{"plan": plan})
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
// Status of the subscription. Allowed values are: n_trial, active, past_due, paused, and canceled.
func UpdateSubscription(customerID string, plan string) (*models.Subscription, error) {
	response, err := client.Put(fmt.Sprintf("%s/%s/subscription", basePath, customerID), map[string]string{"plan": plan})
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
func PauseSubscription(customerID string) (*models.Subscription, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/subscription/pause", basePath, customerID), nil)
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
func ResumeSubscription(customerID string) (*models.Subscription, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/subscription/resume", basePath, customerID), nil)
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
func CancelSubscription(customerID string) (*models.Subscription, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/subscription/cancel", basePath, customerID), nil)
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
func CreatePaymentSource(customerID string, paymentSource models.PaymentSource) (*models.PaymentSource, error) {

	response, err := client.Post(fmt.Sprintf("%s/%s/payment_sources", basePath, customerID), paymentSource)
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
func DeletePaymentSource(customerID string, paymentSourceID string) (*models.PaymentSource, error) {
	response, err := client.Delete(fmt.Sprintf("%s/%s/payment_sources/%s", basePath, customerID, paymentSourceID))
	if err != nil {
		return nil, err
	}

	var ps models.PaymentSource
	if err := json.Unmarshal(response, &ps); err != nil {
		return nil, err
	}
	return &ps, nil
}
