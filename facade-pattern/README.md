# Facade Pattern Implementation - File/Media Converter

## Overview

This project demonstrates the **Facade Design Pattern** through a simplified file/media conversion system implemented in Go. The facade pattern provides a unified, simplified interface to a complex subsystem of file processing components.

## 🎯 Assignment Requirements

**Assignment #5**: Implement Facade for file/media conversion (simplified)
- ✅ **Go Implementation**: Complete Go codebase with proper structure
- ✅ **UML Diagram**: Class diagram showing facade structure
- ✅ **Clean Code Principles**: Applied throughout the implementation
- ✅ **Report**: Comprehensive documentation
- ✅ **GitHub Submission**: Code available in repository

## 🏗️ Architecture

### Facade Pattern Structure

```
Client → MediaConverterFacade → [Validators, Processors, Converters]
```

The facade hides the complexity of:
- **File Validation** (format checking, type detection)
- **File Processing** (optimization, preparation)
- **File Conversion** (format transformation by type)

### Project Structure

```
facade-pattern/
├── cmd/
│   └── main.go                 # Main application demonstrating the facade
├── internal/
│   ├── facade/
│   │   └── media_converter_facade.go  # Main Facade implementation
│   ├── validators/
│   │   └── file_validator.go         # File validation subsystem
│   ├── processors/
│   │   └── file_processor.go         # File processing subsystem
│   └── converters/
│       └── file_converters.go        # File conversion subsystem
├── go.mod
└── README.md
```

## 🧹 Clean Code Principles Applied

### 1. Single Responsibility Principle (SRP)
- **FileValidator**: Only handles file validation
- **FileProcessor**: Only handles file processing/optimization
- **ImageConverter**: Only handles image conversions
- **VideoConverter**: Only handles video conversions
- **DocumentConverter**: Only handles document conversions

### 2. Meaningful Names
```go
// ✅ Clear, intention-revealing names
func (f *MediaConverterFacade) ConvertFile(inputPath, outputPath, targetFormat string) error
func (v *FileValidator) IsImageFile(filename string) bool
func (p *FileProcessor) OptimizeForConversion(filename string) error
```

### 3. Small Functions
- Each function has a single, well-defined purpose
- Functions are kept short and focused
- Complex operations are broken down into smaller helper functions

### 4. Clear Comments and Documentation
- Package-level documentation explaining purpose
- Function documentation describing behavior
- Inline comments for complex logic

### 5. Error Handling
```go
// ✅ Proper error handling with context
if err := f.validateInputFile(inputPath); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}
```

### 6. Consistent Formatting and Structure
- Consistent naming conventions (camelCase for Go)
- Proper package organization
- Clear separation of concerns

## 🔧 Features

### Core Functionality
1. **Single File Conversion**: Convert individual files with custom output paths
2. **Quick Conversion**: Simplified conversion with auto-generated output names
3. **Batch Conversion**: Process multiple files in one operation
4. **Format Support**: Images, videos, and documents

### Supported Formats

| Category | Formats |
|----------|---------|
| **Images** | JPG, JPEG, PNG, GIF, BMP |
| **Videos** | MP4, AVI, MOV, MKV |
| **Documents** | PDF, DOC, DOCX, TXT |

## 🚀 Usage Examples

### Basic Usage
```go
// Create facade instance
converter := facade.NewMediaConverterFacade()

// Single file conversion
err := converter.ConvertFile(
    "input/document.pdf", 
    "output/document.docx", 
    "docx"
)

// Quick conversion (auto-naming)
err := converter.QuickConvert("photo.jpg", "png")

// Batch conversion
files := []string{"video1.avi", "video2.mkv", "image.jpg"}
err := converter.BatchConvert(files, "mp4")
```

## 🎭 Facade Pattern Benefits

### Without Facade (Complex):
```go
// Client needs to know about all subsystems
validator := validators.NewFileValidator()
processor := processors.NewFileProcessor()
imageConverter := converters.NewImageConverter()

// Client must orchestrate all operations
err := validator.ValidateFile(inputFile)
err = processor.ProcessFile(inputFile, outputFile)
err = processor.OptimizeForConversion(inputFile)
err = imageConverter.Convert(inputFile, outputFile, format)
```

### With Facade (Simple):
```go
// Client only needs to know about the facade
converter := facade.NewMediaConverterFacade()

// Single method call handles everything
err := converter.ConvertFile(inputFile, outputFile, format)
```

### Key Benefits:
- ✅ **Simplified Interface**: One class to rule them all
- ✅ **Reduced Coupling**: Client doesn't depend on subsystems
- ✅ **Easy to Use**: Complex operations made simple
- ✅ **Maintainable**: Changes to subsystems don't affect client
- ✅ **Testable**: Easy to mock and test

## 🏃‍♂️ Running the Application

```bash
# Navigate to the project directory
cd facade-pattern

# Run the main application
go run cmd/main.go

# Or build and run
go build -o converter cmd/main.go
./converter
```

## 🧪 Testing the Implementation

The main application demonstrates three key scenarios:

1. **Single File Conversion**: Shows detailed conversion process
2. **Quick Conversion**: Demonstrates simplified API
3. **Batch Processing**: Shows handling multiple files

## 📊 UML Class Diagram

```
┌─────────────────────────────┐
│     MediaConverterFacade    │
├─────────────────────────────┤
│ - validator: FileValidator  │
│ - processor: FileProcessor  │
│ - imageConverter: ImageConv │
│ - videoConverter: VideoConv │
│ - docConverter: DocumentConv│
├─────────────────────────────┤
│ + ConvertFile()             │
│ + QuickConvert()            │
│ + BatchConvert()            │
│ + GetSupportedFormats()     │
└─────────────────────────────┘
              │
              │ uses
              ▼
┌─────────────────────────────┐
│      Subsystem Classes      │
├─────────────────────────────┤
│ FileValidator               │
│ FileProcessor               │
│ ImageConverter              │
│ VideoConverter              │
│ DocumentConverter           │
└─────────────────────────────┘
```

## 📝 Report Summary

### Pattern Implementation Success
- ✅ **Facade Pattern**: Successfully implemented with clear separation
- ✅ **Clean Code**: Applied principles throughout the codebase
- ✅ **Go Best Practices**: Proper package structure and naming
- ✅ **Documentation**: Comprehensive comments and README

### Technical Highlights
1. **Encapsulation**: Complex subsystem operations hidden behind simple interface
2. **Modularity**: Each component has clear responsibility
3. **Extensibility**: Easy to add new file types or converters
4. **Error Handling**: Robust error propagation and context
5. **User Experience**: Simple API for complex operations

### Educational Value
This implementation demonstrates how the Facade pattern can transform a complex system of multiple interacting components into a simple, easy-to-use interface, while maintaining clean code principles and Go best practices.

## 👨‍💻 Author

**Nurashi** - Software Design Patterns Assignment #5

**Date**: October 2025

**Course**: Software Design Patterns