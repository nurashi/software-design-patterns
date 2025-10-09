package wallet

import "fmt"

// Transaction represents a simple transfer record.
type Transaction struct {
	ID       string
	Amount   float64
	Receiver string
}

// String returns a short human-readable representation of the transaction.
func (t *Transaction) String() string {
	return fmt.Sprintf("[TX %s] -> %.2f to %s", t.ID, t.Amount, t.Receiver)
}
