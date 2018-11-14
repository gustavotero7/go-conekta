package main

import (
	"log"

	"github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/customers"
	"github.com/gustavotero7/go-conekta/models"
)

func main() {
	client.APIKey = "your_conekta_private_key"
	customer, err := customers.Create(models.Customer{
		Name:  "Fulano Perez",
		Email: "fulano@example.com",
		Phone: "+5215555555555",
	})
	if err != nil {
		log.Println("Err: ", err)
		return
	}
	log.Println("Response: ", customer)

	// customers.Update(models.Customer{})
	// customers.Delete("customer_id")
	// customers.CreatePaymentSource("customer_id", models.PaymentSource{})
	// orders.Create(models.Order{})
	// ...
}