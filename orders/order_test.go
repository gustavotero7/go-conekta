package orders

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/models"
)

var (
	s = func(serveStatus int, serveResponse interface{}) *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, _ := json.Marshal(serveResponse)
			w.WriteHeader(serveStatus)
			w.Write(b)
		}))
	}
)

func TestOrderCreate(t *testing.T) {
	tests := []struct {
		useStatus       int
		useResponse     interface{}
		name            string
		wantNilResponse bool
		wantErr         bool
	}{
		{
			name:            "OK",
			useStatus:       200,
			useResponse:     models.Order{},
			wantNilResponse: false,
			wantErr:         false,
		},
		{
			name:            "Bad status code (Not 200)",
			useStatus:       400,
			useResponse:     nil,
			wantNilResponse: true,
			wantErr:         true,
		},
		{
			name:            "Invalid response body (int instead Order struct)",
			useStatus:       200,
			useResponse:     0,
			wantNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := Create(models.Order{})
			if (err != nil) != tt.wantErr {
				t.Errorf("order.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (got == nil) != tt.wantNilResponse {
				t.Errorf("order.Create() got = %v, wantNilResponse %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrderUpdate(t *testing.T) {
	tests := []struct {
		useStatus       int
		useResponse     interface{}
		name            string
		wantNilResponse bool
		wantErr         bool
	}{
		{
			name:            "OK",
			useStatus:       200,
			useResponse:     models.Order{},
			wantNilResponse: false,
			wantErr:         false,
		},
		{
			name:            "Bad status code (Not 200)",
			useStatus:       400,
			useResponse:     nil,
			wantNilResponse: true,
			wantErr:         true,
		},
		{
			name:            "Invalid response body (int instead Order struct)",
			useStatus:       200,
			useResponse:     0,
			wantNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := Update(models.Order{})
			if (err != nil) != tt.wantErr {
				t.Errorf("order.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (got == nil) != tt.wantNilResponse {
				t.Errorf("order.Update() got = %v, wantNilResponse %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrderCapture(t *testing.T) {
	tests := []struct {
		useStatus       int
		useResponse     interface{}
		name            string
		wantNilResponse bool
		wantErr         bool
	}{
		{
			name:            "OK",
			useStatus:       200,
			useResponse:     models.Order{},
			wantNilResponse: false,
			wantErr:         false,
		},
		{
			name:            "Bad status code (Not 200)",
			useStatus:       400,
			useResponse:     nil,
			wantNilResponse: true,
			wantErr:         true,
		},
		{
			name:            "Invalid response body (int instead Order struct)",
			useStatus:       200,
			useResponse:     0,
			wantNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)
			got, err := Capture("order_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("order.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (got == nil) != tt.wantNilResponse {
				t.Errorf("order.Capture() got = %v, wantNilResponse %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrderRefund(t *testing.T) {
	tests := []struct {
		useStatus       int
		useResponse     interface{}
		name            string
		wantNilResponse bool
		wantErr         bool
	}{
		{
			name:            "OK",
			useStatus:       200,
			useResponse:     models.Order{},
			wantNilResponse: false,
			wantErr:         false,
		},
		{
			name:            "Bad status code (Not 200)",
			useStatus:       400,
			useResponse:     nil,
			wantNilResponse: true,
			wantErr:         true,
		},
		{
			name:            "Invalid response body (int instead Order struct)",
			useStatus:       200,
			useResponse:     0,
			wantNilResponse: true,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := Refund(models.Order{})
			if (err != nil) != tt.wantErr {
				t.Errorf("order.Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (got == nil) != tt.wantNilResponse {
				t.Errorf("order.Refund() got = %v, wantNilResponse %v", err, tt.wantErr)
				return
			}
		})
	}
}
