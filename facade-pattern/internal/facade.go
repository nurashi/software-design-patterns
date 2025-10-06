package internal

import (
	"fmt"
	"facade-pattern/internal/services"
)

// OrderFacade acts as a simplified interface to the complex order system.
// It combines payment, shipping, and inventory services into one easy-to-use API.
type OrderFacade struct {
	payment   services.PaymentService
	shipping  services.ShippingService
	inventory services.InventoryService
}

// NewOrderFacade constructs and initializes an OrderFacade with given servicee
func NewOrderFacade(payment services.PaymentService, shipping services.ShippingService, inventory services.InventoryService) *OrderFacade {
	return &OrderFacade{
		payment:   payment,
		shipping:  shipping,
		inventory: inventory,
	}
}

// PlaceOrder hides the complexity of multiple subsystems
// The client just calls this single method to complete the entire order process
func (o *OrderFacade) PlaceOrder(item string, address string, amount float64) {
	fmt.Println("Starting order process...")

	if !o.inventory.CheckStock(item) {
		fmt.Println("Item is out of stock")
		return
	}

	o.inventory.ReserveItem(item)
	o.payment.ProcessPayment(amount)
	o.shipping.ShipItem(address)

	fmt.Println("Order completed successffuly")
}
