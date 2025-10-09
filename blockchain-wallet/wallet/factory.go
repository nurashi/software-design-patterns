package wallet

// WalletFactory creates wallets based on type
type WalletFactory struct{}

// CreateWallet returns wallet instance by type
func (f WalletFactory) CreateWallet(walletType, owner string) Wallet {
	switch walletType {
	case "hot":
		w := &BaseWallet{Owner: owner, FeeStrategy: BitcoinFeeStrategy{}}
		return w
	case "cold":
		w := &BaseWallet{Owner: owner, FeeStrategy: EthereumFeeStrategy{}}
		return w
	default:
		return nil
	}
}
