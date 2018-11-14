package models

import "fmt"

type ConektaError struct {
	Object  string   `json:"object,omitempty"`
	Type    string   `json:"type,omitempty"`
	LogId   string   `json:"log_id,omitempty"`
	Details []Detail `json:"details,omitempty"`
}

func (c *ConektaError) Error() string {
	return fmt.Sprintf("Conekta error. Object: %s\nType: %s \nDetails:%s", c.Object, c.Type, c.Details)
}
