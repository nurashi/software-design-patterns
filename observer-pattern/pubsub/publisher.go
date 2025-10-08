package pubsub

// Subscriber defines the interface for all observers
type Subscriber interface {
	Update(event string)
}

// Publisher defines the subject that manages subscribers
type Publisher struct {
	subscribers []Subscriber
}

// Subscribe adds a new subscriber
func (p *Publisher) Subscribe(s Subscriber) {
	p.subscribers = append(p.subscribers, s)
}

// Unsubscribe removes a subscriber
func (p *Publisher) Unsubscribe(s Subscriber) {
	for i, sub := range p.subscribers {
		if sub == s {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
}

// NotifySubscribers sends event notifications to all subscribers
func (p *Publisher) NotifySubscribers(event string) {
	for _, s := range p.subscribers {
		s.Update(event)
	}
}
