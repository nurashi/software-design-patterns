package main

import (
	"fmt"
	"visitor-pattern/visitor"
)

func main() {
	root := visitor.SampleFileStructure()

	// Calculate total size
	sizeVisitor := &visitor.SizeCalculatorVisitor{}
	root.Accept(sizeVisitor)
	fmt.Printf("Total size: %d KB\n", sizeVisitor.Total)

	// Search by extension
	searchVisitor := &visitor.SearchVisitor{Ext: ".go"}
	root.Accept(searchVisitor)
	fmt.Printf("Files with '%s': %v\n", searchVisitor.Ext, searchVisitor.Results)
}
