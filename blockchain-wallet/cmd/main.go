package main

import (
	"blockchain-wallet/wallet"
)

func main() {
	facade := wallet.NewWalletFacade()

	w := facade.CreateWallet("Nurasyl", "hot")

	// Decorator for MFA
	w = wallet.NewMFAWallet(w, "email")

	// Strategy for Bitcoin
	w.SetFeeStrategy(wallet.BitcoinFeeStrategy{})

	// Observer
	facade.Subscribe(wallet.LoggerObserver{})

	// Send
	facade.SendTransaction(w, "receiver123", 10.0)
}
