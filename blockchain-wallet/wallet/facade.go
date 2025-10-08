package wallet

type WalletFacade struct {
	Notifier *TransactionNotifier
	Network  *Network
}

func NewWalletFacade() *WalletFacade {
	return &WalletFacade{
		Notifier: &TransactionNotifier{},
		Network:  GetNetworkInstance(),
	}
}

func (f *WalletFacade) CreateWallet(owner, walletType string) Wallet {
	return CreateWallet(walletType, owner)
}

func (f *WalletFacade) SendTransaction(w Wallet, receiver string, amount float64) {
	tx := w.Send(receiver, amount)
	f.Network.Broadcast(tx)
	f.Notifier.Notify(tx)
}

func (f *WalletFacade) Subscribe(o Observer) {
	f.Notifier.Register(o)
}
