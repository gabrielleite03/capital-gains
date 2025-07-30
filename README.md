# Capital Gains
![My Skills](https://skillicons.dev/icons?i=go)



This project calculates capital gains for stock operations using Go.
O objetivo deste exercício é implementar um programa de linha de comando (CLI) que calcula o imposto a
ser pago sobre lucros ou prejuízos de operações no mercado financeiro de ações.


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

executando a aplicação para receber as linhas por comando:
```
Win:
go run .\cmd\api\main.go

Linux/Mac:
go run cmd/api/main.go

```


executando a aplicação utilizando Input Redirection:
```
Win:
go run .\cmd\api\main.go < .\tests\case_1_input.txt

Linux/Mac:
go run ./cmd/api/main.go < ./tests/case_1_input.txt
```

## Requirements
- Go 1.18+

## License
MIT

## Mockgen
mockgen -source=D:\Projects\capital-gains\service\capital_gain.go -destination=D:\Projects\capital-gains\service\capital_gain_mock.go



# Capital Gains
![My Skills](https://skillicons.dev/icons?i=go)

Este projeto calcula **ganhos de capital** para operações de ações utilizando Go.  
O objetivo deste exercício é implementar um programa de linha de comando (CLI) que
