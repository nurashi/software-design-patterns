package main

import "bridge-pattern/bridge"

func main() {
	
	// Create different storage implementations
	localStorage := &bridge.LocalDiskStorage{}
	s3Storage := &bridge.S3Storage{}
	driveStorage := &bridge.GoogleDriveStorage{}

	// Reports
	report1 := bridge.NewReport(localStorage)
	report1.Store("quarterly_report.txt")

	report2 := bridge.NewReport(s3Storage)
	report2.Store("annual_report.txt")

	// PDFs
	pdf1 := bridge.NewPDF(driveStorage)
	pdf1.Store("presentation.pdf")

	pdf2 := bridge.NewPDF(localStorage)
	pdf2.Store("invoice.pdf")

	// Images
	image1 := bridge.NewImage(s3Storage)
	image1.Store("photo.jpg")

	image2 := bridge.NewImage(driveStorage)
	image2.Store("diagram.png")
}
