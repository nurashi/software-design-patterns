package wallet

import "fmt"

// Wallet represents a minimal wallet interface able to send funds
// and configure a fee strategy.
type Wallet interface {
	Send(receiver string, amount float64) *Transaction
	SetFeeStrategy(strategy FeeStrategy)
	GetOwner() string
}

// BaseWallet is a simple concrete wallet implementation.
type BaseWallet struct {
	Owner       string
	FeeStrategy FeeStrategy // fee calculation strategy
}

// SetFeeStrategy sets the wallet's fee calculation strategy.
func (w *BaseWallet) SetFeeStrategy(strategy FeeStrategy) {
	w.FeeStrategy = strategy
}

// GetOwner returns the wallet owner identifier.
func (w *BaseWallet) GetOwner() string {
	return w.Owner
}

// Send performs a send operation, applies fee strategy and returns a Transaction.
func (w *BaseWallet) Send(receiver string, amount float64) *Transaction {
	fee := 0.0
	if w.FeeStrategy != nil {
		fee = w.FeeStrategy.Calculate(amount)
	}
	tx := &Transaction{
		ID:       "TX123",
		Amount:   amount - fee,
		Receiver: receiver,
	}
	fmt.Printf("%s sent %.2f (fee %.2f) to %s\n", w.Owner, amount, fee, receiver)
	return tx
}
