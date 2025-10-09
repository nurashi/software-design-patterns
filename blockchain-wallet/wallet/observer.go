package wallet

import "fmt"

// Observer receives transaction updates.
type Observer interface {
	Update(tx *Transaction)
}

// Subject can register observers and notify them.
type Subject interface {
	Register(o Observer)
	Notify(tx *Transaction)
}

// TransactionNotifier holds observers and notifies on new transactions.
type TransactionNotifier struct {
	observers []Observer
}

// Register adds an observer to the notifier.
func (n *TransactionNotifier) Register(o Observer) {
	n.observers = append(n.observers, o)
}

// Notify informs all registered observers about a transaction.
func (n *TransactionNotifier) Notify(tx *Transaction) {
	for _, o := range n.observers {
		o.Update(tx)
	}
}

// LoggerObserver logs transaction updates to stdout.
type LoggerObserver struct{}

func (LoggerObserver) Update(tx *Transaction) {
	fmt.Println("Logger: transaction confirmed ", tx)
}
