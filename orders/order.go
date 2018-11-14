package orders

import (
	"encoding/json"
	"fmt"

	"github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/models"
)

const basePath = "/orders"

// Create Creates a new Order
func Create(order models.Order) (*models.Order, error) {
	response, err := client.Post(basePath, order)
	if err != nil {
		return nil, err
	}

	var cr models.Order
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Update Updates an existing Order
func Update(order models.Order) (*models.Order, error) {
	response, err := client.Put(fmt.Sprintf("%s/%s", basePath, order.ID), order)
	if err != nil {
		return nil, err
	}

	var cr models.Order
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Capture Process a pre-authorized order.
func Capture(orderID string) (*models.Order, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/capture", basePath, orderID), nil)
	if err != nil {
		return nil, err
	}

	var cr models.Order
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Refund A Refund details the amount and reason why an order was refunded.
// Reasons:
//	requested_by_client
//	cannot_be_fulfilled
//	duplicated_transaction
//	suspected_fraud
//	other
func Refund(order models.Order) (*models.Order, error) {
	response, err := client.Post(fmt.Sprintf("%s/%s/refunds", basePath, order.ID), order)
	if err != nil {
		return nil, err
	}

	var cr models.Order
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}
