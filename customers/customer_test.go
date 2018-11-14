package customers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestCreate(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Customer
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Customer{},
			want:        &models.Customer{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead ConektaResponse struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := Create(models.Customer{})
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Customer
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Customer{},
			want:        &models.Customer{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead ConektaResponse struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)
			got, err := Update(models.Customer{})
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Customer
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Customer{},
			want:        &models.Customer{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead ConektaResponse struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := Delete("customer_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSubscription(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Subscription
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Subscription{},
			want:        &models.Subscription{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead Subscription struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := CreateSubscription("customer_id", "plan X")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.CreateSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.CreateSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSubscription(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Subscription
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Subscription{},
			want:        &models.Subscription{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead Subscription struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)
			got, err := UpdateSubscription("customer_id", "plan X")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.UpdateSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.UpdateSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPauseSubscription(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Subscription
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Subscription{},
			want:        &models.Subscription{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead Subscription struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := PauseSubscription("customer_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.PauseSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.PauseSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResumeSubscription(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Subscription
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Subscription{},
			want:        &models.Subscription{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead Subscription struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := ResumeSubscription("customer_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.ResumeSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.ResumeSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelSubscription(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.Subscription
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.Subscription{},
			want:        &models.Subscription{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead Subscription struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := CancelSubscription("customer_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.CancelSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.CancelSubscription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePaymentSource(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.PaymentSource
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.PaymentSource{},
			want:        &models.PaymentSource{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead PaymentSource struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := CreatePaymentSource("customer_id", models.PaymentSource{})
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.CreatePaymentSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.CreatePaymentSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeletePaymentSource(t *testing.T) {
	tests := []struct {
		useStatus   int
		useResponse interface{}
		name        string
		want        *models.PaymentSource
		wantErr     bool
	}{
		{
			name:        "OK",
			useStatus:   200,
			useResponse: models.PaymentSource{},
			want:        &models.PaymentSource{},
			wantErr:     false,
		},
		{
			name:        "Bad status code (Not 200)",
			useStatus:   400,
			useResponse: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "Invalid response body (int instead PaymentSource struct)",
			useStatus:   200,
			useResponse: 0,
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := s(tt.useStatus, tt.useResponse)
			client.SetClientURL(server.URL)

			got, err := DeletePaymentSource("customer_id", "payment_source_id")
			if (err != nil) != tt.wantErr {
				t.Errorf("customer.DeletePaymentSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer.DeletePaymentSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCreate() {
	// You only need to set client.APIKey once (then you can do all the calls you want with customer or orders packages)
	client.APIKey = "your_conekta_private_key"

	// Only name and email are mandatory fields, other fields are optional
	// See https://developers.conekta.com/api?language=bash#create-customer
	customer, err := Create(models.Customer{
		Name:  "Fulano Perez",
		Email: "fulano@example.com",
		Phone: "+5215555555555",
	})

	// Handle error
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	// Do something with the response
	log.Println("Response: ", customer)
}

func ExampleUpdate() {
	// You only need to set client.APIKey once (then you can do all the calls you want with customer or orders packages)
	client.APIKey = "your_conekta_private_key"

	// Only id, name and email are mandatory fields, other fields are optional
	// See https://developers.conekta.com/api?language=bash#update-customer
	customer, err := Update(models.Customer{
		ID:    "cus_123456789",
		Name:  "New Name",
		Email: "new@email.com",
		Phone: "+5215555555555",
	})

	// Handle error
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	// Do something with the response
	log.Println("Response: ", customer)
}

func ExampleDelete() {
	// You only need to set client.APIKey once (then you can do all the calls you want with customer or orders packages)
	client.APIKey = "your_conekta_private_key"

	// See https://developers.conekta.com/api?language=bash#update-customer
	customer, err := Delete("cus_123456789")

	// Handle error
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	// Do something with the response
	log.Println("Response: ", customer)
}
