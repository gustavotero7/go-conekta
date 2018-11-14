package conekta

import (
	"encoding/json"

	"github.com/gustavotero7/go-conekta/models"
)

type Order struct {
	models.Order
}

// Create Creates a new Order
func (o *Order) Create() (*models.ConektaResponse, error) {
	response, err := request("POST", "/orders", o)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Update Updates an existing Order
func (o *Order) Update() (*models.ConektaResponse, error) {
	response, err := request("PUT", "/orders/"+o.ID, o)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Capture Process a pre-authorized order.
func (o *Order) Capture() (*models.ConektaResponse, error) {
	response, err := request("POST", "/orders/"+o.ID+"/capture", nil)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

// Refund A Refund details the amount and reason why an order was refunded.
func (o *Order) Refund() (*models.ConektaResponse, error) {
	response, err := request("POST", "/orders/"+o.ID+"/refunds", o)
	if err != nil {
		return nil, err
	}

	var cr models.ConektaResponse
	if err := json.Unmarshal(response, &cr); err != nil {
		return nil, err
	}
	return &cr, nil
}
