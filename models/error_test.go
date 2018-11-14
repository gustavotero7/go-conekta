package models

import "testing"

func TestConektaError_Error(t *testing.T) {
	type fields struct {
		Object  string
		Type    string
		LogId   string
		Details []Detail
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				Object:  "Some Object",
				Type:    "Some Type",
				Details: []Detail{},
			},
			want: "Conekta error. Object: Some Object\nType: Some Type \nDetails:[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConektaError{
				Object:  tt.fields.Object,
				Type:    tt.fields.Type,
				LogId:   tt.fields.LogId,
				Details: tt.fields.Details,
			}
			if got := c.Error(); got != tt.want {
				t.Errorf("ConektaError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
