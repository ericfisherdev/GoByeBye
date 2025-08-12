# assets/

Static assets and resources for the application.

The `/assets` directory contains non-Go files that are needed by the application at runtime.

## Expected Contents

- `models/` - AI models (ONNX files for U²-Net and other background removal models)
- `icons/` - Application icons for different platforms
- `images/` - Sample images or UI resources
- `shaders/` - GPU shaders if applicable

## Guidelines

- Keep binary files organized by type
- Consider using go:embed for embedding assets in the binary
- Document model sources and licenses
- Optimize images and assets for size

## Embedding Assets

For Fyne applications, assets can be bundled using:
```bash
fyne bundle -o bundled.go assets/
```

Or use go:embed (Go 1.16+):
```go
//go:embed models/*
var modelsFS embed.FS
```