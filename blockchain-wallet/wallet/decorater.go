package wallet

import "fmt"

// WalletDecorator wraps wallet functionality
type WalletDecorator struct {
	Wrapped Wallet
}

// MFAWalletDecorator adds multi-factor authentication before sending
type MFAWalletDecorator struct {
	WalletDecorator
	Method string
}

func (m *MFAWalletDecorator) Send(receiver string, amount float64) Transaction {
	fmt.Printf("[MFA] Authenticating via %s...\n", m.Method)
	return m.Wrapped.Send(receiver, amount)
}

// Forwarding methods so decorators satisfy Wallet interface by delegating
// calls to the wrapped wallet implementation.
func (w *WalletDecorator) SetFeeStrategy(strategy FeeStrategy) {
	if w.Wrapped != nil {
		w.Wrapped.SetFeeStrategy(strategy)
	}
}

func (w *WalletDecorator) GetOwner() string {
	if w.Wrapped != nil {
		return w.Wrapped.GetOwner()
	}
	return ""
}
