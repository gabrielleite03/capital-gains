# Capital Gains
![My Skills](https://skillicons.dev/icons?i=go)



O objetivo deste exercício é implementar um programa de linha de comando (CLI) que calcula o imposto a
ser pago sobre lucros ou prejuízos de operações no mercado financeiro de ações.

## Decisões Técnicas e Arquiteturais

O projeto foi desenvolvido em Go adotando a Arquitetura Hexagonal (Ports and Adapters). Essa escolha teve como objetivo principal garantir um baixo acoplamento entre as camadas da aplicação, permitindo maior flexibilidade e facilidade de manutenção. Essa abordagem facilita a substituição de componentes externos (como banco de dados ou serviços de mensageria) sem impactar a lógica de negócio.

Optou-se por não utilizar logs durante a execução do sistema, evitando assim a poluição da saída do console, especialmente em ambientes de desenvolvimento e execução de testes. Essa decisão teve como foco manter a saída padrão mais limpa e objetiva.

Para garantir a qualidade do código, foram implementados testes unitários e de integração, cobrindo os principais fluxos da aplicação e validando a interação entre os módulos. No entanto, não foram realizados testes de carga, estresse ou volume, ficando essa etapa como uma possível evolução futura do projeto para avaliar seu comportamento em cenários de alta demanda.

---

⚙️ Requisitos
Go 1.18+

## ▶️ Uso

### Execução interativa (linha por linha)
Permite inserir as operações manualmente no terminal:
```bash
# Windows
go run .\cmd\api\main.go

# Linux/Mac
go run cmd/api/main.go
```

### Execução com Input Redirection (utilizando arquivos)
Permite usar arquivos de entrada já definidos:
```bash
# Windows
go run .\cmd\api\main.go < .\tests\case_1_input.txt

# Linux/Mac
go run ./cmd/api/main.go < ./tests/case_1_input.txt

```
### Execução dos testes
```bash
go test -v ./...

```

---

## 📖 Exemplo de Operações
Arquivo case_1_input.txt:

```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 100},{"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]

```

Saída esperada:

```json
[{"tax": 0.0},{"tax": 0.0},{"tax": 0.0}]
```



## 📂 Estrutura do Projeto
- `cmd/api/main.go`: Ponto de entrada principal. Lê as entradas, processa as operações e gera as saídas.
- `internal/core/models/`: Modelos de dados para ganhos de capital, operações e ações.
- `internal/core/ports/`: Interfaces de serviços.
- `internal/service/`: Lógica de negócios dos ganhos de capital.
- `internal/domain/`: Objetos de transferência de dados (DTOs) e tipos de domínio.
- `tests/`: Arquivos de testes e exemplos de entrada.

---


## Mockgen
mockgen -source=D:\Projects\capital-gains\service\capital_gain.go -destination=D:\Projects\capital-gains\service\capital_gain_mock.go
