package wallet

import "fmt"

type Wallet interface {
	Send(receiver string, amount float64) *Transaction
	SetFeeStrategy(strategy FeeStrategy)
	GetOwner() string
}

type BaseWallet struct {
	Owner       string
	FeeStrategy FeeStrategy // strategy to pay between each other
}


// Setter
func (w *BaseWallet) SetFeeStrategy(strategy FeeStrategy) {
	w.FeeStrategy = strategy
}

// Getter
func (w *BaseWallet) GetOwner() string {
	return w.Owner
}


// 
func (w *BaseWallet) Send(receiver string, amount float64) *Transaction {
	fee := 0.0 // fee - платеж
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
