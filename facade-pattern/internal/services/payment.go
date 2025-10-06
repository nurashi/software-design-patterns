package services

import "fmt"

// PaymentService defines the interface for payment processing systems.
type PaymentService interface {
	ProcessPayment(amount float64)
}

// CreditCardPayment is a concrete implementation of the payment system.
type CreditCardPayment struct{}

// ProcessPayment simulates processing a credit card transaction.
func (c *CreditCardPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f...\n", amount)
}
