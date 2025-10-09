package wallet

import "fmt"

// Wallet defines basic wallet interface
type Wallet interface {
	Send(receiver string, amount float64) Transaction
	SetFeeStrategy(strategy FeeStrategy)
	GetOwner() string
}

// BaseWallet is core wallet implementation
type BaseWallet struct {
	Owner      string
	FeeStrategy FeeStrategy
}

func (w *BaseWallet) Send(receiver string, amount float64) Transaction {
	fee := w.FeeStrategy.Calculate(amount)
	fmt.Printf("%s is sending %.2f (+ fee %.2f) to %s\n", w.Owner, amount, fee, receiver)
	return Transaction{ID: "tx123", Amount: amount, Receiver: receiver}
}

func (w *BaseWallet) SetFeeStrategy(strategy FeeStrategy) {
	w.FeeStrategy = strategy
}

func (w *BaseWallet) GetOwner() string {
	return w.Owner
}
