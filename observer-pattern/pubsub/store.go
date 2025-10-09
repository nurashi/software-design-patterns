package pubsub

import "fmt"

// Store is a concrete publisher that notifies customers when a new product arrives
type Store struct {
	Publisher // composition
	Name string
}

// AddNewProduct adds a product and notifies all subscribers
func (s *Store) AddNewProduct(product string) {
	fmt.Printf("\n%s received new product: %s\n", s.Name, product)
	s.NotifySubscribers(fmt.Sprintf("%s is now available in %s!", product, s.Name))
}
