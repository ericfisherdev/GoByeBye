# GoByeBye

AI-powered background removal desktop application built with Go.

## Overview

GoByeBye is a cross-platform desktop application that uses advanced AI models (U²-Net, ONNX Runtime) to automatically remove backgrounds from images. It features a user-friendly drag-and-drop interface, batch processing capabilities, multiple AI model support, alpha matting edge refinement, and GPU acceleration for professional-quality results.

## Features

- **AI-Powered Background Removal**: Utilizes U²-Net and other state-of-the-art models
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Drag & Drop Interface**: Simple and intuitive user experience
- **Batch Processing**: Process multiple images simultaneously
- **Multiple AI Models**: Support for U²-Net, MODNet, RVM, and more
- **Alpha Matting**: Advanced edge refinement for clean cutouts
- **GPU Acceleration**: Optional CUDA/OpenCL support for faster processing
- **Multiple Output Formats**: PNG, JPEG, WebP with configurable quality

## Installation

### Prerequisites

- Go 1.21 or later
- ONNX Runtime (for AI model inference)
- GPU drivers (optional, for acceleration)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/ericfisherdev/GoByeBye.git
cd GoByeBye

# Download dependencies
go mod download

# Build the application
go build -o gobyebye cmd/gobyebye/main.go

# Run the application
./gobyebye
```

## Usage

### GUI Mode

Simply launch the application and drag images onto the window:

```bash
./gobyebye
```

### CLI Mode (Coming Soon)

```bash
# Process single image
gobyebye -i input.jpg -o output.png

# Batch processing
gobyebye -i "*.jpg" -o output_dir/

# Use GPU acceleration
gobyebye -i input.jpg -o output.png --gpu

# Specify model
gobyebye -i input.jpg -o output.png --model u2net
```

## Project Structure

```
GoByeBye/
├── cmd/
│   └── gobyebye/       # Application entry point
├── internal/
│   ├── core/           # Core processing logic
│   ├── ui/             # User interface components
│   ├── config/         # Configuration management
│   └── models/         # AI model management
├── pkg/
│   ├── logger/         # Logging utilities
│   ├── errors/         # Error handling
│   └── utils/          # Common utilities
├── configs/            # Configuration files
├── assets/             # Static resources
├── docs/               # Documentation
└── test/               # Integration tests
```

## Configuration

Edit `configs/default.yaml` to customize:

- Model selection and paths
- GPU acceleration settings
- Output format and quality
- UI preferences
- Performance tuning

## Development

### Running Tests

```bash
# Unit tests
go test ./...

# Integration tests
go test ./test/...

# With coverage
go test -cover ./...
```

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Architecture

GoByeBye follows clean architecture principles:

- **Interface-driven design** for testability
- **Dependency injection** for loose coupling
- **SOLID principles** throughout the codebase
- **Comprehensive error handling** with context
- **Structured logging** for debugging

## Performance

- Optimized for batch processing
- Memory-efficient image handling
- Optional GPU acceleration via CUDA/OpenCL
- Configurable worker pools for parallel processing

## Roadmap

- [ ] CLI interface implementation
- [ ] Additional AI model support
- [ ] Video background removal
- [ ] Real-time preview
- [ ] Plugin system
- [ ] Cloud processing option
- [ ] Mobile companion app

## License

See [LICENSE](LICENSE) file for details.

## Acknowledgments

- U²-Net model by Qin et al.
- ONNX Runtime by Microsoft
- Go community for excellent libraries

## Support

For issues, questions, or suggestions:
- Open an issue on [GitHub](https://github.com/ericfisherdev/GoByeBye/issues)
- Check the [documentation](docs/)

---

Built with ❤️ using Go