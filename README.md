[![pipeline status](https://gitlab.com/gustavotero7/go-conekta/badges/develop/pipeline.svg)](https://gitlab.com/gustavotero7/go-conekta/commits/develop) [![coverage report](https://gitlab.com/gustavotero7/go-conekta/badges/develop/coverage.svg)](https://gitlab.com/gustavotero7/go-conekta/commits/develop)

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

### Simple Example

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

### License

MIT License

Copyright (c) 2018 Gustavo Otero

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

### Resources

https://godoc.org/github.com/gustavotero7/go-conekta

https://developers.conekta.com/libraries/javascript

https://developers.conekta.com/api

https://developers.conekta.com/tutorials/card
