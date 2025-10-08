package visitor

// Visitor defines operations for different elements
type Visitor interface {
	VisitFile(f *File)
	VisitDirectory(d *Directory)
}

// Element interface that all file system objects must implement
type Element interface {
	Accept(v Visitor)
}
