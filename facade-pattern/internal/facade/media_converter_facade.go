// Package facade implements the Facade design patternpackage facade

// Clean Code: The facade provides a simplified interface to complex subsystems
package facade

import (
	"fmt"
	"path/filepath"
	"strings"

	"facade-pattern/internal/converters"
	"facade-pattern/internal/processors"
	"facade-pattern/internal/validators"
)

// MediaConverterFacade provides a simplified interface for file/media conversion
// Clean Code: Class name clearly indicates its purpose and pattern used
type MediaConverterFacade struct {
	validator         *validators.FileValidator
	processor         *processors.FileProcessor
	imageConverter    *converters.ImageConverter
	videoConverter    *converters.VideoConverter
	documentConverter *converters.DocumentConverter
}

// NewMediaConverterFacade creates a new facade instance
// Clean Code: Constructor pattern with dependency injection
func NewMediaConverterFacade() *MediaConverterFacade {
	return &MediaConverterFacade{
		validator:         validators.NewFileValidator(),
		processor:         processors.NewFileProcessor(),
		imageConverter:    converters.NewImageConverter(),
		videoConverter:    converters.NewVideoConverter(),
		documentConverter: converters.NewDocumentConverter(),
	}
}

// ConvertFile is the main facade method that simplifies the entire conversion process
// Clean Code: Single public method that encapsulates complex operations
func (f *MediaConverterFacade) ConvertFile(inputPath, outputPath, targetFormat string) error {
	fmt.Printf("🚀 Starting file conversion process...\n")
	fmt.Printf("════════════════════════════════════════\n")

	// Step 1: Validate input file
	if err := f.validateInputFile(inputPath); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Step 2: Process and optimize file
	if err := f.processFile(inputPath, outputPath); err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}

	// Step 3: Convert file based on type
	if err := f.convertBasedOnType(inputPath, outputPath, targetFormat); err != nil {
		return fmt.Errorf("conversion failed: %w", err)
	}

	fmt.Printf("════════════════════════════════════════\n")
	fmt.Printf("✨ Conversion completed successfully!\n")
	fmt.Printf("   Output: %s\n", outputPath)

	return nil
}

// QuickConvert provides a simplified conversion with automatic output naming
// Clean Code: Convenience method for common use case
func (f *MediaConverterFacade) QuickConvert(inputPath, targetFormat string) error {
	outputPath := f.generateOutputPath(inputPath, targetFormat)
	return f.ConvertFile(inputPath, outputPath, targetFormat)
}

// BatchConvert handles multiple file conversions
// Clean Code: Method name clearly indicates batch processing capability
func (f *MediaConverterFacade) BatchConvert(inputFiles []string, targetFormat string) error {
	fmt.Printf("📦 Starting batch conversion of %d files...\n", len(inputFiles))

	successCount := 0
	for i, inputFile := range inputFiles {
		fmt.Printf("\n[%d/%d] Processing: %s\n", i+1, len(inputFiles), filepath.Base(inputFile))

		if err := f.QuickConvert(inputFile, targetFormat); err != nil {
			fmt.Printf("❌ Failed to convert %s: %v\n", filepath.Base(inputFile), err)
			continue
		}
		successCount++
	}

	fmt.Printf("\n📊 Batch conversion summary:\n")
	fmt.Printf("   ✅ Successful: %d/%d\n", successCount, len(inputFiles))
	fmt.Printf("   ❌ Failed: %d/%d\n", len(inputFiles)-successCount, len(inputFiles))

	return nil
}

// Clean Code: Private helper methods with clear, single responsibilities
func (f *MediaConverterFacade) validateInputFile(inputPath string) error {
	fmt.Printf("🔍 Validating input file...\n")

	if err := f.validator.ValidateFile(inputPath); err != nil {
		return err
	}

	fmt.Printf("   ✅ File validation passed\n")
	return nil
}

func (f *MediaConverterFacade) processFile(inputPath, outputPath string) error {
	fmt.Printf("\n🔧 Processing and optimizing file...\n")

	if err := f.processor.ProcessFile(inputPath, outputPath); err != nil {
		return err
	}

	if err := f.processor.OptimizeForConversion(inputPath); err != nil {
		return err
	}

	fmt.Printf("   ✅ Processing completed\n")
	return nil
}

func (f *MediaConverterFacade) convertBasedOnType(inputPath, outputPath, targetFormat string) error {
	fmt.Printf("\n🔄 Converting file...\n")

	switch {
	case f.validator.IsImageFile(inputPath):
		return f.imageConverter.Convert(inputPath, outputPath, targetFormat)
	case f.validator.IsVideoFile(inputPath):
		return f.videoConverter.Convert(inputPath, outputPath, targetFormat)
	case f.validator.IsDocumentFile(inputPath):
		return f.documentConverter.Convert(inputPath, outputPath, targetFormat)
	default:
		return fmt.Errorf("unsupported file type for conversion")
	}
}

func (f *MediaConverterFacade) generateOutputPath(inputPath, targetFormat string) string {
	dir := filepath.Dir(inputPath)
	baseName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return filepath.Join(dir, fmt.Sprintf("%s_converted.%s", baseName, targetFormat))
}

// GetSupportedFormats returns information about supported formats
// Clean Code: Informational method for client convenience
func (f *MediaConverterFacade) GetSupportedFormats() map[string][]string {
	return map[string][]string{
		"images":    {"jpg", "jpeg", "png", "gif", "bmp"},
		"videos":    {"mp4", "avi", "mov", "mkv"},
		"documents": {"pdf", "doc", "docx", "txt"},
	}
}
