// Pacote service_test contém testes unitários e de integração para o CapitalGainService.
// Esses testes validam o comportamento do serviço, incluindo cenários de erro,
// operações vazias, instanciamento correto e uso de mocks com o gomock.
package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
)

// TestGetCapitalGain_ErrorFromStockService valida o comportamento do método GetCapitalGain
// quando ocorre um erro vindo do StockService. Espera-se que um erro seja retornado
// e que o resultado (ganhos de capital) seja nulo.
func TestGetCapitalGain_ErrorFromStockService(t *testing.T) {
	sService := ports.NewStockService()
	service := NewCapitalGainService(sService)

	cg, err := service.GetCapitalGain("test")
	if err == nil {
		t.Errorf("Esperava erro, mas recebeu nil")
	}
	if cg != nil {
		t.Errorf("Esperava capital gains nulo, mas recebeu %v", cg)
	}
}

// TestGetCapitalGain_EmptyOperations valida o comportamento quando a entrada
// não contém operações válidas. Espera-se um erro e uma lista vazia ou nula de ganhos de capital.
func TestGetCapitalGain_EmptyOperations(t *testing.T) {
	sService := ports.NewStockService()
	service := NewCapitalGainService(sService)

	cg, err := service.GetCapitalGain("test")
	if err == nil {
		t.Errorf("Esperava erro de parsing, mas recebeu: %v", err)
	}
	if cg != nil && len(*cg) != 0 {
		t.Errorf("Esperava lista vazia de capital gains, mas recebeu %v", cg)
	}
}

// TestNewCapitalGainService valida a criação de uma nova instância do CapitalGainService.
// Garante que a instância não seja nula e que implemente a interface correta.
func TestNewCapitalGainService(t *testing.T) {
	sService := ports.NewStockService()
	cgService := NewCapitalGainService(sService)

	// Verifica se a instância retornada não é nula
	assert.NotNil(t, cgService, "NewCapitalGainService deveria retornar uma instância não nula")

	// Verifica se a instância implementa a interface CapitalGainService
	_, ok := cgService.(CapitalGainService)
	assert.True(t, ok, "NewCapitalGainService deveria retornar um tipo que implementa CapitalGainService")
}

// TestReportService_GenerateTax_Success demonstra como usar mocks com o gomock
// para validar chamadas ao método GetCapitalGain do CapitalGainService.
// São testados cenários de retorno nulo, retorno válido e erro.
func TestReportService_GenerateTax_Success(t *testing.T) {
	// 1. Cria um controlador gomock para gerenciar o ciclo de vida dos mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 2. Cria uma instância mock do CapitalGainService
	mockCapitalGainService := NewMockCapitalGainService(ctrl)

	// Dados de teste que simulam os retornos possíveis do método GetCapitalGain
	expectedCapitalGains := [][]models.CapitalGains{
		{
			{Tax: models.MyFloat64(5.0)},
			{Tax: models.MyFloat64(10.5)},
		},
		{
			{Tax: models.MyFloat64(20.0)},
		},
	}
	inputString := "[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 100},{\"operation\":\"sell\", \"unit-cost\":15.00, \"quantity\": 50},{\"operation\":\"sell\", \"unit-cost\":15.00, \"quantity\": 50}]"

	// 3. Configura as expectativas no mock:
	// - Primeiro: retorna nil para ganhos de capital e erro nil
	// - Segundo: retorna ganhos de capital válidos
	// - Terceiro: retorna erro
	mockCapitalGainService.EXPECT().GetCapitalGain(inputString).Return(nil, nil).Times(1)
	mockCapitalGainService.EXPECT().GetCapitalGain(inputString).Return(&expectedCapitalGains, nil).Times(1)
	mockCapitalGainService.EXPECT().GetCapitalGain(inputString).Return(nil, errors.New("a string de entrada está vazia")).Times(1)

	// Valida retorno do primeiro cenário
	ops, err := mockCapitalGainService.GetCapitalGain(inputString)
	assert.Nil(t, err)
	assert.Nil(t, ops)

	// Valida retorno do segundo cenário
	ops, err = mockCapitalGainService.GetCapitalGain(inputString)
	assert.Nil(t, err)
	assert.NotNil(t, ops)

	// Valida retorno do terceiro cenário (erro)
	ops, err = mockCapitalGainService.GetCapitalGain(inputString)
	assert.NotNil(t, err)
	assert.Nil(t, ops)
}
