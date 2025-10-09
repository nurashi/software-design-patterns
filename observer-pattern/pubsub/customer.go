package pubsub

import "fmt"

// Customer is a concrete subscriber
type Customer struct {
	Name string 
}

// Update is called when store notifies about new products
func (c *Customer) Update(event string) {
	fmt.Printf("%s received update: %s\n", c.Name, event)
}
