package visitor

// SizeCalculatorVisitor calculates the total size of files
type SizeCalculatorVisitor struct {
	Total int
}

func (v *SizeCalculatorVisitor) VisitFile(f *File) {
	v.Total += f.Size
}


// at this particular case logic is not needed, for simplicity
func (v *SizeCalculatorVisitor) VisitDirectory(d *Directory) {}
