package wallet

import "fmt"

// ThirdPartyNetwork is a foreign API we can't modify
type ThirdPartyNetwork struct{}

func (n *ThirdPartyNetwork) Push(data string) {
	fmt.Println("[3rd-Party Network] Broadcasting:", data)
}

// BlockchainNetwork defines the expected interface
type BlockchainNetwork interface {
	Broadcast(tx Transaction)
}

// NetworkAdapter adapts ThirdPartyNetwork to BlockchainNetwork
type NetworkAdapter struct {
	ThirdParty *ThirdPartyNetwork
}

func (a *NetworkAdapter) Broadcast(tx Transaction) {
	a.ThirdParty.Push(fmt.Sprintf("Adapted TX: %s", tx.String()))
}
