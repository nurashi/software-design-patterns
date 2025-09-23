package main

import (
	"fmt"

	"adapter-pattern/internal/service"
)

func main() {
	// Declare a variable of type PaymentProcessor interface
	var processor service.PaymentProcessor

	// Use PayPal adapter to process a payment
	processor = service.NewPayPalAdapter()
	fmt.Println(processor.Pay(50.0))

	// Use Forte adapter to process a payment
	processor = service.NewForteAdapter()
	fmt.Println(processor.Pay(19.99))

	// Use Kaspi adapter to process a payment
	processor = service.NewKaspiAdapter()
	fmt.Println(processor.Pay(5000.0))
}
