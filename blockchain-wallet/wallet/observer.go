package wallet

import "fmt"

type Observer interface {
	Update(tx *Transaction)
}

type Subject interface {
	Register(o Observer)
	Notify(tx *Transaction)
}

type TransactionNotifier struct {
	observers []Observer
}

func (n *TransactionNotifier) Register(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *TransactionNotifier) Notify(tx *Transaction) {
	for _, o := range n.observers {
		o.Update(tx)
	}
}

type LoggerObserver struct{}

func (LoggerObserver) Update(tx *Transaction) {
	fmt.Println("Logger: transaction confirmed →", tx)
}
