

[![pipeline status](https://gitlab.com/gustavotero7/go-conekta/badges/feature/setup-ci/pipeline.svg)](https://gitlab.com/gustavotero7/go-conekta/commits/feature/setup-ci) [![coverage report](https://gitlab.com/gustavotero7/go-conekta/badges/feature/setup-ci/coverage.svg)](https://gitlab.com/gustavotero7/go-conekta/commits/feature/setup-ci)

### Go-Conekta
A Wrapper for use conekta's api v2 in golang inspired in [sait/go-conekta](https://github.com/sait/go-conekta)

This tutorial assumes the next:

* Have a conekta account
* Have a frontend that tokenize cards
* Knowledge of conekta's api

### Usage

First get the package

```
go get -u github.com/gustavotero7/go-conekta/...
```

Import in your project

```go
import (
    "github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/customers"
	"github.com/gustavotero7/go-conekta/models"
)
```

### Simple Order Example

```go
package main

import (
	"log"

	"github.com/gustavotero7/go-conekta/client"
	"github.com/gustavotero7/go-conekta/customers"
	"github.com/gustavotero7/go-conekta/models"
)

func main() {
	client.APIKey = "your_private_key"
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
```
**You only need to set client.APIKey once**

### Resources

https://developers.conekta.com/libraries/javascript

https://developers.conekta.com/api

https://developers.conekta.com/tutorials/card
