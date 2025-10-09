package wallet

import "fmt"

// Network is a very small singleton representing the blockchain network.
type Network struct{}

var instance *Network

// GetNetworkInstance returns a singleton Network.
func GetNetworkInstance() *Network {
	if instance == nil {
		instance = &Network{}
		fmt.Println("Initialized blockchain network...")
	}
	return instance
}

// Broadcast simulates broadcasting a transaction to the network.
func (n *Network) Broadcast(tx *Transaction) {
	fmt.Println("Broadcasting transaction:", tx)
}
