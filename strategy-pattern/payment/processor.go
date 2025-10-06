package payment

// PaymentProcessor acts as the context, using different payment strategies.
type PaymentProcessor struct {
	strategy PaymentStrategy
}

// SetStrategy allows changing the payment method at runtime.
func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// ProcessPayment delegates the payment process to the selected strategy.
func (p *PaymentProcessor) ProcessPayment(amount float64) {
	if p.strategy == nil {
		panic("no payment strategy selected")
	}
	p.strategy.Pay(amount)
}
