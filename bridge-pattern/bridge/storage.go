package bridge

import "fmt"

// Storage interface - Implementation side of the bridge
type Storage interface {
	SaveDocument(name string)
}

// LocalDiskStorage - Concrete implementation
type LocalDiskStorage struct{}

func (l *LocalDiskStorage) SaveDocument(name string) {
	fmt.Printf("Saving '%s' to local disk storage\n", name)
}

// S3Storage - Concrete implementation
type S3Storage struct{}

func (s *S3Storage) SaveDocument(name string) {
	fmt.Printf("Saving '%s' to AWS S3 storage\n", name)
}

// GoogleDriveStorage - Concrete implementation
type GoogleDriveStorage struct{}

func (g *GoogleDriveStorage) SaveDocument(name string) {
	fmt.Printf("Saving '%s' to Google Drive storage\n", name)
}
