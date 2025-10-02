# UML Class Diagram - Facade Pattern Implementation

## Detailed UML Diagram

```
                                  ┌─────────────────────────────────────┐
                                  │              Client                 │
                                  │           (main.go)                 │
                                  └──────────────┬──────────────────────┘
                                                 │
                                                 │ uses
                                                 ▼
                    ┌─────────────────────────────────────────────────────────┐
                    │                MediaConverterFacade                     │
                    ├─────────────────────────────────────────────────────────┤
                    │ - validator: *FileValidator                             │
                    │ - processor: *FileProcessor                             │
                    │ - imageConverter: *ImageConverter                       │
                    │ - videoConverter: *VideoConverter                       │
                    │ - documentConverter: *DocumentConverter                 │
                    ├─────────────────────────────────────────────────────────┤
                    │ + NewMediaConverterFacade(): *MediaConverterFacade      │
                    │ + ConvertFile(input, output, format string): error      │
                    │ + QuickConvert(input, format string): error             │
                    │ + BatchConvert(files []string, format string): error    │
                    │ + GetSupportedFormats(): map[string][]string            │
                    │ - validateInputFile(path string): error                 │
                    │ - processFile(input, output string): error              │
                    │ - convertBasedOnType(input, output, format string): err │
                    │ - generateOutputPath(input, format string): string      │
                    └──────────────┬──────────────┬──────────────┬────────────┘
                                   │              │              │
                            delegates to   delegates to   delegates to
                                   │              │              │
                                   ▼              ▼              ▼
            ┌─────────────────────────┐  ┌─────────────────────────┐  ┌─────────────────────────┐
            │     FileValidator       │  │     FileProcessor       │  │      Converters         │
            ├─────────────────────────┤  ├─────────────────────────┤  ├─────────────────────────┤
            │ - supportedFormats:     │  │                         │  │                         │
            │   map[string]bool       │  │                         │  │                         │
            ├─────────────────────────┤  ├─────────────────────────┤  ├─────────────────────────┤
            │ + NewFileValidator():   │  │ + NewFileProcessor():   │  │                         │
            │   *FileValidator        │  │   *FileProcessor        │  │                         │
            │ + ValidateFile(name     │  │ + ProcessFile(input,    │  │                         │
            │   string): error        │  │   output string): error │  │                         │
            │ + IsImageFile(name      │  │ + OptimizeForConversion │  │                         │
            │   string): bool         │  │   (filename string): err│  │                         │
            │ + IsVideoFile(name      │  │ - optimizeImage(file    │  │                         │
            │   string): bool         │  │   string): error        │  │                         │
            │ + IsDocumentFile(name   │  │ - optimizeVideo(file    │  │                         │
            │   string): bool         │  │   string): error        │  │                         │
            └─────────────────────────┘  │ - optimizeDocument(file │  │                         │
                                         │   string): error        │  │                         │
                                         └─────────────────────────┘  │                         │
                                                                      │                         │
                                                                      │                         │
                                         ┌─────────────────────────┐  │                         │
                                         │    ImageConverter       │  │                         │
                                         ├─────────────────────────┤  │                         │
                                         │ + NewImageConverter():  │  │                         │
                                         │   *ImageConverter       │  │                         │
                                         │ + Convert(input, output,│  │                         │
                                         │   format string): error │  │                         │
                                         │ - validateConversion    │  │                         │
                                         │   (input, format string)│  │                         │
                                         │   : error               │  │                         │
                                         └─────────────────────────┘  │                         │
                                                                      │                         │
                                         ┌─────────────────────────┐  │                         │
                                         │    VideoConverter       │  │                         │
                                         ├─────────────────────────┤  │                         │
                                         │ + NewVideoConverter():  │  │                         │
                                         │   *VideoConverter       │  │                         │
                                         │ + Convert(input, output,│  │                         │
                                         │   format string): error │  │                         │
                                         │ - validateConversion    │  │                         │
                                         │   (input, format string)│  │                         │
                                         │   : error               │  │                         │
                                         └─────────────────────────┘  │                         │
                                                                      │                         │
                                         ┌─────────────────────────┐  │                         │
                                         │   DocumentConverter     │  │                         │
                                         ├─────────────────────────┤  │                         │
                                         │ + NewDocumentConverter()│  │                         │
                                         │   : *DocumentConverter  │  │                         │
                                         │ + Convert(input, output,│  │                         │
                                         │   format string): error │  │                         │
                                         │ - validateConversion    │  │                         │
                                         │   (input, format string)│  │                         │
                                         │   : error               │  │                         │
                                         └─────────────────────────┘  │                         │
                                                                      └─────────────────────────┘
```

## Pattern Relationships

### Facade Pattern Structure:
1. **Facade (MediaConverterFacade)**: 
   - Provides simplified interface to complex subsystem
   - Delegates work to appropriate subsystem components
   - Hides complexity from client

2. **Subsystem Classes**:
   - **FileValidator**: Handles file format validation and type detection
   - **FileProcessor**: Manages file processing and optimization
   - **Converters (Image/Video/Document)**: Handle specific format conversions

3. **Client**: 
   - Only interacts with the Facade
   - Doesn't need knowledge of subsystem complexity
   - Uses simple, unified interface

### Key Design Benefits:

#### Before Facade (Complex):
```
Client ←→ FileValidator
       ←→ FileProcessor  
       ←→ ImageConverter
       ←→ VideoConverter
       ←→ DocumentConverter
```

#### With Facade (Simplified):
```
Client ←→ MediaConverterFacade ←→ [All Subsystems]
```

### Interaction Flow:
1. Client calls `ConvertFile()` on Facade
2. Facade delegates to `FileValidator` for validation
3. Facade delegates to `FileProcessor` for optimization
4. Facade determines file type and delegates to appropriate converter
5. Facade returns unified result to client

This UML diagram demonstrates how the Facade pattern encapsulates complexity and provides a clean, simple interface for file conversion operations.