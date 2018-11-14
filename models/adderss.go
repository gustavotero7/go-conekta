package models

type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"string,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Residential bool   `json:"residential,omitempty"`
	Object      string `json:"object,omitempty"`
}
