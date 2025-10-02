package bridge

import "fmt"

// Report - Refined abstraction
type Report struct {
	*Document
}

func NewReport(storage Storage) *Report {
	return &Report{
		Document: NewDocument(storage),
	}
}

func (r *Report) Store(name string) {
	fmt.Println("Preparing report for storage...")
	r.Document.Store(name)
}
