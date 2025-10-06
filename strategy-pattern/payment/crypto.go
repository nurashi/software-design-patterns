package payment

import "fmt"

// CryptoPayment is a concrete strategy for cryptocurrency payments.
type CryptoPayment struct {
	WalletAddress string
}

func (c *CryptoPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Crypto wallet: %s\n", amount, c.WalletAddress)
}
