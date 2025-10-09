package wallet

// WalletFacade simplifies complex subsystem interaction
type WalletFacade struct {
	Factory  WalletFactory
	Notifier Subject
	Network  BlockchainNetwork
}

// CreateWallet wraps factory creation
func (f *WalletFacade) CreateWallet(owner, walletType string) Wallet {
	return f.Factory.CreateWallet(walletType, owner)
}

// SendTransaction handles sending, broadcasting and notifying
func (f *WalletFacade) SendTransaction(w Wallet, receiver string, amount float64) {
	tx := w.Send(receiver, amount)
	f.Network.Broadcast(tx)
	f.Notifier.Notify(tx)
}

// Subscribe adds new observer
func (f *WalletFacade) Subscribe(o Observer) {
	f.Notifier.Register(o)
}
