package payment

// PaymentStrategy defines a common interface for all payment methods.
type PaymentStrategy interface {
	Pay(amount float64)
}
