package visitor

// SizeCalculatorVisitor calculates the total size of files
type SizeCalculatorVisitor struct {
	Total int
}

func (v *SizeCalculatorVisitor) VisitFile(f *File) {
	v.Total += f.Size
}

func (v *SizeCalculatorVisitor) VisitDirectory(d *Directory) {
	// No specific logic for directory itself
}
