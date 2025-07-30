// Package infrastructure fornece constantes e variáveis relacionadas à estrutura
// de arquivos e parâmetros de configuração utilizados nos testes e na aplicação.
package infrastructure

import (
	"fmt"
	"os"
)

var (
	// TestCasesPath define o caminho base onde os arquivos de teste estão localizados.
	//
	// O caminho é construído dinamicamente utilizando os separadores de diretório
	// apropriados para o sistema operacional em uso, garantindo compatibilidade.
	// Exemplo de saída: "../../tests/" (em sistemas Unix).
	TestCasesPath = fmt.Sprintf("%s%s%s%s%s%s", "..", string(os.PathSeparator), "..", string(os.PathSeparator), "tests", string(os.PathSeparator))
)

const (
	// TaxOnProfit representa a taxa de imposto sobre o lucro utilizada nos cálculos financeiros.
	// Este tipo de dados poderia estar em um arquivo ou serviço de configuração para qua a aletração seja
	// feita de maneira dinâmica
	TaxOnProfit float64 = 0.2
)

// Constantes relacionadas a nomes de arquivos de teste.
const (
	// Case1Case2InputFile é o nome do arquivo de entrada usado para os testes dos casos 1 e 2.
	Case1Case2InputFile = "case_1_case_2_input.txt"

	// Case1Case2OutputFile é o nome do arquivo de saída esperado para os testes dos casos 1 e 2.
	Case1Case2OutputFile = "case_1_case_2_output.txt"

	// PrefixUseCase define o prefixo comum usado para gerar nomes de arquivos de casos de teste.
	PrefixUseCase string = "case_"

	// SufixUseCase define o sufixo comum para os arquivos de saída dos casos de teste.
	SufixUseCase string = "_output.txt"
)
