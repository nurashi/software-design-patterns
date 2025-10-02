package bridge

import "fmt"

// Image - Refined abstraction
type Image struct {
	*Document
}

func NewImage(storage Storage) *Image {
	return &Image{
		Document: NewDocument(storage),
	}
}

func (i *Image) Store(name string) {
	fmt.Println("Optimizing image for storage...")
	i.Document.Store(name)
}
