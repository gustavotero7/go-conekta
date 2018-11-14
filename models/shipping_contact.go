package models

type ShippingContact struct {
	ConektaBase
	ID             string  `json:"id,omitempty"`
	Phone          string  `json:"phone,omitempty"`
	Receiver       string  `json:"receiver,omitempty"`
	BetweenStreets string  `json:"between_streets,omitempty"`
	Address        Address `json:"address,omitempty"`
}
