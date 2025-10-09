package main

import (
	"blockchain-wallet/wallet"
)

func main() {
	// Setup facade with adapter network
	network := &wallet.NetworkAdapter{ThirdParty: &wallet.ThirdPartyNetwork{}}
	facade := wallet.WalletFacade{Network: network}

	// Create and decorate wallet
	base := facade.CreateWallet("Nurasyl", "hot")
	mfaWallet := &wallet.MFAWalletDecorator{WalletDecorator: wallet.WalletDecorator{Wrapped: base}, Method: "Email"}

	// Subscribe logger observer
	facade.Subscribe(wallet.LoggerObserver{})

	// Perform transaction
	facade.SendTransaction(mfaWallet, "Alice", 50.0)
}
