package wallet

import "fmt"

// Transaction represents a blockchain transaction
type Transaction struct {
	ID       string
	Amount   float64
	Receiver string
}

// String returns readable transaction info
func (t Transaction) String() string {
	return fmt.Sprintf("Transaction[ID=%s, Amount=%.2f, Receiver=%s]", t.ID, t.Amount, t.Receiver)
}
