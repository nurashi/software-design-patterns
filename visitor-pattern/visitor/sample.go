package visitor

// SampleFileStructure builds a simple file system tree for demo/testing
func SampleFileStructure() *Directory {
	return &Directory{
		Name: "root",
		Children: []Element{
			&File{Name: "main.go", Size: 120},
			&File{Name: "report.txt", Size: 80},
			&Directory{
				Name: "assets",
				Children: []Element{
					&File{Name: "logo.png", Size: 150},
					&File{Name: "photo.jpg", Size: 200},
				},
			},
		},
	}
}
