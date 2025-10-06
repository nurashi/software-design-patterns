package services

import "fmt"

// InventoryService defines the interface for inventory management systems.
type InventoryService interface {
	CheckStock(item string) bool
	ReserveItem(item string)
}

// WarehouseInventory is a concrete implementation of inventory control.
type WarehouseInventory struct{}

// CheckStock simulates checking if an item is available in the warehouse.
func (w *WarehouseInventory) CheckStock(item string) bool {
	fmt.Printf("Checking stock for item: %s\n", item)
	// Assume item is always in stock for simplicity
	return true
}

// ReserveItem simulates reserving an item for a customer.
func (w *WarehouseInventory) ReserveItem(item string) {
	fmt.Printf("Reserving item: %s\n", item)
}
