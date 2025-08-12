# internal/

Private application and library code.

The `/internal` directory contains code that is private to this application. The Go compiler enforces that packages in the `internal` directory can only be imported by code in the parent directory tree.

## Structure

- `config/` - Application configuration handling
- `core/` - Core business logic and background removal processing
- `models/` - AI model loading and management
- `ui/` - User interface components (Fyne GUI)

## Guidelines

- This code cannot be imported by external projects
- Put code here that you don't want others to depend on
- Organize by feature or domain, not by technical layer
- Keep packages focused and cohesive

## Note

The `internal` directory is a Go convention enforced by the compiler. Any attempt to import these packages from outside the module will result in a compilation error.