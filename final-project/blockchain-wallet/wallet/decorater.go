package wallet

import "fmt"

type WalletDecorator struct {
	Wrapped Wallet
}

type MFAWalletDecorator struct {
	WalletDecorator
	Method string
}

func NewMFAWallet(w Wallet, method string) Wallet {
	return &MFAWalletDecorator{
		WalletDecorator{Wrapped: w},
		method,
	}
}

func (m *MFAWalletDecorator) Send(receiver string, amount float64) *Transaction {
	fmt.Printf("Performing MFA (%s)... verified.\n", m.Method)
	return m.Wrapped.Send(receiver, amount)
}

func (m *MFAWalletDecorator) SetFeeStrategy(s FeeStrategy) {
	m.Wrapped.SetFeeStrategy(s)
}

func (m *MFAWalletDecorator) GetOwner() string {
	return m.Wrapped.GetOwner()
}
