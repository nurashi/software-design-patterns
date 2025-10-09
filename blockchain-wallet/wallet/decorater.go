package wallet

import "fmt"

// WalletDecorator wraps another Wallet to extend behavior.
type WalletDecorator struct {
	Wrapped Wallet
}

// MFAWalletDecorator adds a simple MFA step before sending.
type MFAWalletDecorator struct {
	WalletDecorator
	Method string
}

// NewMFAWallet constructs an MFA decorator around a Wallet.
func NewMFAWallet(w Wallet, method string) Wallet {
	return &MFAWalletDecorator{
		WalletDecorator{Wrapped: w},
		method,
	}
}

// Send performs an MFA check then delegates to the wrapped wallet.
func (m *MFAWalletDecorator) Send(receiver string, amount float64) *Transaction {
	fmt.Printf("Performing MFA (%s)... verified.\n", m.Method)
	return m.Wrapped.Send(receiver, amount)
}

// SetFeeStrategy delegates to the wrapped wallet.
func (m *MFAWalletDecorator) SetFeeStrategy(s FeeStrategy) {
	m.Wrapped.SetFeeStrategy(s)
}

// GetOwner delegates to the wrapped wallet.
func (m *MFAWalletDecorator) GetOwner() string {
	return m.Wrapped.GetOwner()
}
