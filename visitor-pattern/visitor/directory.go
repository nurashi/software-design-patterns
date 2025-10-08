package visitor

type Directory struct {
	Name     string
	Children []Element
}

func (d *Directory) Accept(v Visitor) {
	v.VisitDirectory(d)
	for _, child := range d.Children {
		child.Accept(v)
	}
}
