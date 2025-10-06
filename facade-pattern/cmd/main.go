package main

import (
	"facade-pattern/internal"
	"facade-pattern/internal/services"
)

func main() {
	// Create concrete service implementations
	payment := &services.CreditCardPayment{}
	shipping := &services.CourierShipping{}
	invertorty := &services.WarehouseInventory{}


	// Create the facade by injecting all dependencies
	orderFacade := internal.NewOrderFacade(payment, shipping, invertorty)

	// The client interacts only with the Facade
	orderFacade.PlaceOrder("Laptop", "123 Turkistan st, Astana", 1400.99)
}