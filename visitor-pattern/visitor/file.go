package visitor

type File struct {
	Name string
	Size int
}

func (f *File) Accept(v Visitor) {
	v.VisitFile(f)
}
