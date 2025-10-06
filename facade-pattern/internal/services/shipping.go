package services

import "fmt"

// ShippingService defines the interface for delivery systems.
type ShippingService interface {
	ShipItem(address string)
}

// CourierShipping is a concrete implementation of a shipping service.
type CourierShipping struct{}

// ShipItem simulates delivering the product to the provided address.
func (c *CourierShipping) ShipItem(address string) {
	fmt.Printf("Shipping item to address: %s\n", address)
}
