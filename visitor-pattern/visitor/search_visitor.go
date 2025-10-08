package visitor

import "strings"

// SearchVisitor finds files by their extension (e.g. ".go", ".png")
type SearchVisitor struct {
	Ext     string
	Results []string
}

func (v *SearchVisitor) VisitFile(f *File) {
	if strings.HasSuffix(f.Name, v.Ext) {
		v.Results = append(v.Results, f.Name)
	}
}

func (v *SearchVisitor) VisitDirectory(d *Directory) {}
