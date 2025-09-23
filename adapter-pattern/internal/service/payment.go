package service

import "fmt"

// Target interface (client works only with this)
type PaymentProcessor interface {
	Pay(amount float64) string
}

// Adaptee 1: PayPal (has its own method)
type PayPal struct{}

func (p *PayPal) SendPayment(amount float64) string {
	return fmt.Sprintf("Paid %.2f USD with PayPal", amount)
}

// Adapter for PayPal
type PayPalAdapter struct {
	paypal *PayPal
}

func NewPayPalAdapter() *PayPalAdapter {
	return &PayPalAdapter{paypal: &PayPal{}}
}

func (a *PayPalAdapter) Pay(amount float64) string {
	return a.paypal.SendPayment(amount)
}

// Adaptee 2: Forte (different API)
type Forte struct{}

func (f *Forte) ProcessTenge(amount float64) string {
	return fmt.Sprintf("Processed %.2f KZT with Forte", amount)
}

// Adapter for Forte
type ForteAdapter struct {
	forte *Forte
}

func NewForteAdapter() *ForteAdapter {
	return &ForteAdapter{forte: &Forte{}}
}

func (a *ForteAdapter) Pay(amount float64) string {
	return a.forte.ProcessTenge(amount)
}

// Adaptee 3: Kaspi (local API)
type Kaspi struct{}

func (k *Kaspi) PayKZT(amount float64) string {
	return fmt.Sprintf("Processed %.2f KZT with Kaspi", amount)
}

// Adapter for Kaspi
type KaspiAdapter struct {
	kaspi *Kaspi
}

func NewKaspiAdapter() *KaspiAdapter {
	return &KaspiAdapter{kaspi: &Kaspi{}}
}

func (a *KaspiAdapter) Pay(amount float64) string {
	return a.kaspi.PayKZT(amount)
}
