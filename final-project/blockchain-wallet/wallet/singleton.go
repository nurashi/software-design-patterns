package wallet

import "fmt"

type Network struct{}

var instance *Network

func GetNetworkInstance() *Network {
	if instance == nil {
		instance = &Network{}
		fmt.Println("Initialized blockchain network...")
	}
	return instance
}

func (n *Network) Broadcast(tx *Transaction) {
	fmt.Println("Broadcasting transaction:", tx)
}
