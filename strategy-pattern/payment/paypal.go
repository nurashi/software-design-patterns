package payment

import "fmt"

// PayPalPayment is a concrete strategy for PayPal payments.
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal account: %s\n", amount, p.Email)
}
