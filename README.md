# Capital Gains

This project calculates capital gains for stock operations using Go.

## Project Structure
- `cmd/api/main.go`: Main entry point. Reads input, processes operations, and outputs results.
- `internal/core/models/`: Data models for capital gains, operations, and stocks.
- `internal/core/ports/`: Service interfaces.
- `internal/service/`: Business logic for capital gains.
- `internal/domain/`: Domain transfer objects and types.
- `tests/`: Test files and input samples.

## Usage
Run the main program and input stock operations line by line. Type `exit` or `quit` to stop. If the input is empty, the program outputs all processed operations as JSON and exits.

## Example
```
go run cmd/api/main.go
```

## Requirements
- Go 1.18+

## License
MIT
