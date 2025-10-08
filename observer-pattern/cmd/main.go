package main

import "observer-pattern/pubsub"

func main() {
	store := &pubsub.Store{Name: "TechStore"}

	// Create customers (subscribers)
	alex := &pubsub.Customer{Name: "Alex"}
	emma := &pubsub.Customer{Name: "Emma"}
	john := &pubsub.Customer{Name: "John"}

	// Subscribe customers to store updates
	store.Subscribe(alex)
	store.Subscribe(emma)
	store.Subscribe(john)

	// Store adds a new product
	store.AddNewProduct("MacBook Pro 2025")

	// Emma unsubscribes
	store.Unsubscribe(emma)

	// Store adds another product
	store.AddNewProduct("iPhone 17")
}
