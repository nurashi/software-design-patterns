package main

import (
	"strategy-pattern/payment"
)

func main() {
	processor := payment.PaymentProcessor{}

	// Use Credit Card Payment
	processor.SetStrategy(&payment.CreditCardPayment{CardNumber: "+7 987 612 33 33"})
	processor.ProcessPayment(120.50)

	// Use PayPal Payment
	processor.SetStrategy(&payment.PayPalPayment{Email: "nurashi@gmail.com"})
	processor.ProcessPayment(89.99)

	// Use Crypto Payment
	processor.SetStrategy(&payment.CryptoPayment{WalletAddress: "123123123123"})
	processor.ProcessPayment(0.005)
}
