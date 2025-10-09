package wallet

// CreateWallet is a tiny factory that returns a Wallet by type.
func CreateWallet(walletType, owner string) Wallet {
	switch walletType {
	case "hot":
		return &BaseWallet{Owner: owner}
	case "cold":
		return &BaseWallet{Owner: owner}
	default:
		return &BaseWallet{Owner: owner}
	}
}
