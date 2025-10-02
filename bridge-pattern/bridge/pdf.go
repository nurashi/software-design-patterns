package bridge

import "fmt"

// PDF - Refined abstraction
type PDF struct {
	*Document
}

func NewPDF(storage Storage) *PDF {
	return &PDF{
		Document: NewDocument(storage),
	}
}

func (p *PDF) Store(name string) {
	fmt.Println("Compressing PDF for storage...")
	p.Document.Store(name)
}
