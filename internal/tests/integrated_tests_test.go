// Pacote tests contém testes automatizados para validar o cálculo de ganho de capital (CapitalGains)
// utilizando os serviços implementados no projeto. Ele garante que a lógica de cálculo
// esteja correta em diferentes cenários de casos de teste.
package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	reflect "reflect"
	"strings"
	"testing"

	"koto.com/internal/core/infrastructure"
	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
	"koto.com/service"
)

// TestCase_1_case_2 valida os casos de teste 1 e 2 usando arquivos de entrada e saída específicos.
// Ele compara o resultado do serviço de ganho de capital com o resultado esperado no arquivo de saída.
func TestCase_1_case_2(t *testing.T) {
	// Lê os arquivos de entrada e saída do caso de teste
	sInput := readFile(infrastructure.Case1Case2InputFile)
	sOutput := readFile(infrastructure.Case1Case2OutputFile)

	// Cria uma instância do serviço de cálculo de ganho de capital
	cg := service.NewCapitalGainService(ports.NewStockService())
	ops, err := cg.GetCapitalGain(sInput)
	if err != nil {
		t.Errorf("Erro ao obter ganhos de capital: %v", err)
	}

	// Converte o conteúdo do arquivo de saída para modelo
	cGModelsOutput := getModelFromString(sOutput)

	// Valida se a saída calculada é igual à saída esperada
	if !reflect.DeepEqual(ops, cGModelsOutput) {
		t.Errorf("Saída esperada: %v, obtida: %v", cGModelsOutput, ops)
	}
}

// TestCase_all_files executa uma bateria de testes para todos os arquivos de caso de uso disponíveis.
// Cada arquivo é lido e comparado com sua saída esperada.
func TestCase_all_files(t *testing.T) {
	filesQuantity := 9 // quantidade de casos de uso disponíveis
	for i := 1; i < filesQuantity; i++ {
		// Gera o nome do arquivo de caso de teste
		file := fmt.Sprintf("%s%d%s", infrastructure.PrefixUseCase, i, infrastructure.SufixUseCase)
		fmt.Println(file)

		// Executa o teste dinamicamente para cada arquivo encontrado
		t.Run(fmt.Sprintf("Executando testes para %s", file), func(t *testing.T) {
			// Lê entrada e saída do arquivo de caso de uso
			sInput := readFile(file)
			sOutput := readFile(file)

			// Cria instância do serviço de cálculo de ganho de capital
			cg := service.NewCapitalGainService(ports.NewStockService())
			ops, err := cg.GetCapitalGain(sInput)
			if err != nil {
				t.Errorf("Erro ao obter ganhos de capital: %v", err)
			}

			// Converte saída esperada em modelo
			cGModelsOutput := getModelFromString(sOutput)

			// Compara saída calculada com saída esperada
			if !reflect.DeepEqual(ops, cGModelsOutput) {
				t.Errorf("Saída esperada: %v, obtida: %v", cGModelsOutput, ops)
			}
		})
	}
}

// readFile lê o conteúdo de um arquivo de teste e retorna como string.
func readFile(f string) string {
	file, err := os.Open(infrastructure.TestCasesPath + f)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	return string(b)
}

// getModelFromString converte uma string contendo JSONs em um ponteiro para
// matriz de slices de CapitalGains. Cada linha não vazia é interpretada como um JSON.
func getModelFromString(row string) *[][]models.CapitalGains {
	s := strings.Split(row, "\r\n")
	cgs := &[][]models.CapitalGains{}
	for _, r := range s {
		if r != "" {
			var cgRow []models.CapitalGains
			e := json.Unmarshal([]byte(r), &cgRow)
			if e != nil {
				log.Fatalf("Erro ao deserializar operação: %v", e)
			}
			*cgs = append(*cgs, cgRow)
		}
	}

	return cgs
}
