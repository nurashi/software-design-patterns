package main

import "observer-pattern/pubsub"

func main() {
	store := &pubsub.Store{Name: "TechStore"}

	// Create customers (subscribers)
	nur := &pubsub.Customer{Name: "Nurassyl"}
	bal := &pubsub.Customer{Name: "Balanbay"}
	balen := &pubsub.Customer{Name: "Balanbayev"}

	// Subscribe customers to store updates
	store.Subscribe(nur)
	store.Subscribe(bal)
	store.Subscribe(balen)

	// Store adds a new product
	store.AddNewProduct("MacBook Pro 2025")

	// Emma unsubscribes
	store.Unsubscribe(bal)

	// Store adds another product
	store.AddNewProduct("iPhone 17")
}
