package models

type Detail struct {
	Debug_message string `json:"debug_message,omitempty"`
	Message       string `json:"message,omitempty"`
	Code          string `json:"code,omitempty"`
}
