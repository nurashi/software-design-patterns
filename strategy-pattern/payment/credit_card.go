package payment

import "fmt"

// CreditCardPayment is a concrete strategy for credit card payments.
type CreditCardPayment struct {
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card (%s)\n", amount, c.CardNumber)
}
