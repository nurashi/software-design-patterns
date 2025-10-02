// Main application demonstrating the Facade Pattern for file/media conversionpackage cmd

// Clean Code: Main package is clean and focused on demonstration
package main

import (
	"fmt"
	"log"

	"facade-pattern/internal/facade"
)

func main() {
	fmt.Printf("🎭 Facade Pattern Demo - File/Media Converter\n")
	fmt.Printf("=============================================\n\n")

	// Create the facade - this is the only complex object client needs to know about
	// Clean Code: Simple client interface, complexity is hidden
	converter := facade.NewMediaConverterFacade()

	// Display supported formats
	displaySupportedFormats(converter)

	// Demonstrate single file conversion
	demonstrateSingleConversion(converter)

	// Demonstrate quick conversion
	demonstrateQuickConversion(converter)

	// Demonstrate batch conversion
	demonstrateBatchConversion(converter)
}

// Clean Code: Small functions with descriptive names
func displaySupportedFormats(converter *facade.MediaConverterFacade) {
	fmt.Printf("📋 Supported file formats:\n")
	formats := converter.GetSupportedFormats()

	for category, formatList := range formats {
		fmt.Printf("   %s: %v\n", category, formatList)
	}
	fmt.Printf("\n")
}

func demonstrateSingleConversion(converter *facade.MediaConverterFacade) {
	fmt.Printf("🔧 Demo 1: Single File Conversion\n")
	fmt.Printf("─────────────────────────────────\n")

	// Client only needs to call one method - the facade handles all complexity
	err := converter.ConvertFile(
		"documents/report.pdf",
		"documents/report_converted.docx",
		"docx",
	)

	if err != nil {
		log.Printf("Conversion failed: %v", err)
	}
	fmt.Printf("\n")
}

func demonstrateQuickConversion(converter *facade.MediaConverterFacade) {
	// Even simpler interface - output path is automatically generated
	err := converter.QuickConvert("images/photo.jpg", "png")

	if err != nil {
		log.Printf("Quick conversion failed: %v", err)
	}
	fmt.Printf("\n")
}

func demonstrateBatchConversion(converter *facade.MediaConverterFacade) {

	// Batch processing made simple through the facade
	files := []string{
		"videos/vacation.mp4",
		"videos/presentation.avi",
		"images/logo.png",
		"documents/manual.pdf",
	}

	err := converter.BatchConvert(files, "mp4") // Convert all to mp4 (videos only)

	if err != nil {
		log.Printf("Batch conversion failed: %v", err)
	}

	fmt.Printf("\n🎉 All demonstrations completed!\n")
	fmt.Printf("\n💡 Benefits of Facade Pattern:\n")
	fmt.Printf("   ✅ Simplified client interface\n")
	fmt.Printf("   ✅ Hidden subsystem complexity\n")
	fmt.Printf("   ✅ Reduced coupling between client and subsystems\n")
	fmt.Printf("   ✅ Easy to use and maintain\n")
	fmt.Printf("   ✅ Flexible internal implementation\n")
}
