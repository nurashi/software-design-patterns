package bridge

// Document - Abstraction side of the bridge
type Document struct {
	storage Storage
}

func NewDocument(storage Storage) *Document {
	return &Document{storage: storage}
}

func (d *Document) Store(name string) {
	d.storage.SaveDocument(name)
}
