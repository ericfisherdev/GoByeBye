# cmd/

Main applications for this project.

The `/cmd` directory contains the entry points for the GoByeBye application. Each subdirectory should contain a `main.go` file that can be compiled into an executable.

## Structure

- `gobyebye/` - The main GoByeBye desktop application

## Usage

Build the application:
```bash
go build -o gobyebye ./cmd/gobyebye
```

Or use the Makefile:
```bash
make build
```

## Best Practices

- Keep `main.go` files minimal - they should only handle initialization and configuration
- Business logic should be in `/internal` packages
- Each application gets its own subdirectory
- Command-line parsing and setup belongs here