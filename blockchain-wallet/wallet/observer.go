package wallet

import "fmt"

// Observer interface defines how listeners react to transaction events
type Observer interface {
	Update(tx Transaction)
}

// Subject manages observers
type Subject struct {
	observers []Observer
}

// Register adds new observer
func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

// Notify informs all observers about transaction
func (s *Subject) Notify(tx Transaction) {
	for _, o := range s.observers {
		o.Update(tx)
	}
}

// LoggerObserver logs new transactions
type LoggerObserver struct{}

func (l LoggerObserver) Update(tx Transaction) {
	fmt.Println("[Logger] New transaction confirmed:", tx.String())
}
