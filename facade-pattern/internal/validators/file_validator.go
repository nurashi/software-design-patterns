// Package validators provides file validation functionalitypackage validators

// Following Clean Code principle: Single Responsibility - each validator has one job
package validators

import (
	"fmt"
	"path/filepath"
	"strings"
)

// FileValidator validates file properties and formats
type FileValidator struct {
	supportedFormats map[string]bool
}

// NewFileValidator creates a new FileValidator instance
// Clean Code: Meaningful names and clear constructor
func NewFileValidator() *FileValidator {
	return &FileValidator{
		supportedFormats: map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".bmp":  true,
			".mp4":  true,
			".avi":  true,
			".mov":  true,
			".mkv":  true,
			".pdf":  true,
			".doc":  true,
			".docx": true,
			".txt":  true,
		},
	}
}

// ValidateFile checks if the file format is supported
// Clean Code: Small function with clear purpose
func (v *FileValidator) ValidateFile(filename string) error {
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return fmt.Errorf("file must have an extension")
	}

	if !v.supportedFormats[ext] {
		return fmt.Errorf("unsupported file format: %s", ext)
	}

	return nil
}

// IsImageFile checks if the file is an image
// Clean Code: Intention-revealing function name
func (v *FileValidator) IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	imageFormats := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

	for _, format := range imageFormats {
		if ext == format {
			return true
		}
	}
	return false
}

// IsVideoFile checks if the file is a video
func (v *FileValidator) IsVideoFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	videoFormats := []string{".mp4", ".avi", ".mov", ".mkv"}

	for _, format := range videoFormats {
		if ext == format {
			return true
		}
	}
	return false
}

// IsDocumentFile checks if the file is a document
func (v *FileValidator) IsDocumentFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	documentFormats := []string{".pdf", ".doc", ".docx", ".txt"}

	for _, format := range documentFormats {
		if ext == format {
			return true
		}
	}
	return false
}
