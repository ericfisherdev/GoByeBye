# tests/

Test files and testing utilities.

The `/tests` directory contains integration tests, test data, and testing utilities that span multiple packages.

## Structure

- Integration tests that test multiple components together
- End-to-end tests for the complete application
- Test fixtures and sample data
- Performance benchmarks
- Testing utilities and helpers

## Guidelines

- Unit tests should live alongside the code they test (in the same package)
- Put integration and e2e tests here
- Use table-driven tests where appropriate
- Include benchmarks for performance-critical code
- Document how to run specific test suites

## Running Tests

Run all tests:
```bash
go test ./...
```

Run integration tests:
```bash
go test ./tests/...
```

Run with coverage:
```bash
go test -cover ./...
```

## Test Data

Store test images and model files in subdirectories:
- `testdata/` - Standard Go convention for test data (ignored by go tools)
- `fixtures/` - Reusable test fixtures