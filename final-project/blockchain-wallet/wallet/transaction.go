package wallet

import "fmt"

type Transaction struct {
	ID       string
	Amount   float64
	Receiver string
}


// its a transaction 
func (t *Transaction) String() string {
	return fmt.Sprintf("[TX %s] -> %.2f to %s", t.ID, t.Amount, t.Receiver)
}
