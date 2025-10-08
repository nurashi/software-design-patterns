package wallet

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
