# Project Commands and Guidelines

## Build Commands
- Build: `go build`
- Run: `go run main.go`
- Install dependencies: `go mod tidy`

## Test Commands
- Run all tests: `go test ./...`
- Run a specific test: `go test -run TestName`
- Run tests with coverage: `go test -cover ./...`

## Lint Commands
- Format code: `go fmt ./...`
- Run linter: `golangci-lint run`
- Check imports: `goimports -local heads-down-slacker -w .`

## Code Style Guidelines
- Use Go standard formatting (gofmt)
- Group imports: standard library, then third-party, then local
- Follow Go naming conventions (CamelCase for exported, camelCase for private)
- Use error handling with explicit checks (if err != nil)
- Prefer early returns over nested conditionals
- Use meaningful variable names
- Include comments for exported functions and types
- Keep functions small and focused on single responsibility