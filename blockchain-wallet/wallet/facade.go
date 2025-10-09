package wallet

// WalletFacade provides a simplified API for common wallet flows.
type WalletFacade struct {
	Notifier *TransactionNotifier
	Network  *Network
}

// NewWalletFacade constructs a facade with default components.
func NewWalletFacade() *WalletFacade {
	return &WalletFacade{
		Notifier: &TransactionNotifier{},
		Network:  GetNetworkInstance(),
	}
}

// CreateWallet creates a wallet of the given type for owner.
func (f *WalletFacade) CreateWallet(owner, walletType string) Wallet {
	return CreateWallet(walletType, owner)
}

// SendTransaction sends, broadcasts and notifies observers about the tx.
func (f *WalletFacade) SendTransaction(w Wallet, receiver string, amount float64) {
	tx := w.Send(receiver, amount)
	f.Network.Broadcast(tx)
	f.Notifier.Notify(tx)
}

// Subscribe registers an observer for transaction notifications.
func (f *WalletFacade) Subscribe(o Observer) {
	f.Notifier.Register(o)
}
