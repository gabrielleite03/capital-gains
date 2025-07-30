package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"koto.com/internal/core/ports"
	"koto.com/service"
)

// main é o ponto de entrada da aplicação.
//
// Este programa lê entradas do console (stdin), processa cada entrada
// utilizando o serviço de cálculo de ganhos de capital (CapitalGainService)
// e exibe no console a saída em formato JSON.
//
// Fluxo principal:
//  1. Lê entradas linha a linha do stdin utilizando bufio.Scanner.
//  2. Para cada entrada, chama o método GetCapitalGain passando o texto informado.
//  3. Converte os resultados retornados em JSON e acumula em uma lista.
//  4. Exibe todos os resultados processados no console.
//
// Observações:
//   - A leitura é encerrada quando uma linha vazia ("") é informada.
//   - Em caso de erro na obtenção dos dados de ganhos de capital, a entrada é ignorada
//     e o programa segue para a próxima linha.
//
// Exemplo de execução:
//
//	$ echo "ACAO1" | go run main.go
//	[{"date":"2023-01-01","gain":1000.50}, {"date":"2023-02-01","gain":500.00}]
//
// Dependências principais:
//   - ports.NewStockService(): cria uma instância do serviço de ações.
//   - service.NewCapitalGainService(): cria o serviço de cálculo de ganhos de capital
//     que consome o serviço de ações.
func main() {
	s := bufio.NewScanner(os.Stdin)
	var outs []string

	for s.Scan() {
		text := s.Text()
		if text == "" {
			break
		}

		capitalGainService := service.NewCapitalGainService(ports.NewStockService())

		cps, err := capitalGainService.GetCapitalGain(text)
		if err == nil {
			for _, row := range *cps {
				out, _ := json.Marshal(row)
				outs = append(outs, string(out))
			}
		}
	}

	for _, row := range outs {
		fmt.Println(row)
	}
}
