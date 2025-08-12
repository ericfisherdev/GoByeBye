# pkg/

Library code that's safe to use by external applications.

The `/pkg` directory contains Go packages that can be imported by other projects. This is code that you're explicitly making available for others to use.

## Structure

- `errors/` - Custom error types and error handling utilities
- `logger/` - Logging utilities and configuration
- `utils/` - General utility functions (image processing, file handling)

## Guidelines

- Only put code here that you're comfortable with others importing
- Keep APIs stable and well-documented
- Write comprehensive tests for public APIs
- Consider semantic versioning for breaking changes

## Note

While the `/pkg` pattern is debated in the Go community (with some preferring packages at the root level), it's useful in larger applications to explicitly separate public from private code. For this desktop application, it helps organize reusable components that could potentially be extracted into separate modules.