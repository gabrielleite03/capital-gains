package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "koto.com/internal/core/models"
)

// MockCapitalGainService é uma implementação mock (gerada pelo GoMock)
// da interface CapitalGainService.
//
// Este mock é utilizado em testes para simular o comportamento real do serviço,
// permitindo controlar respostas e validar chamadas sem depender de implementações reais.
type MockCapitalGainService struct {
	ctrl     *gomock.Controller
	recorder *MockCapitalGainServiceMockRecorder
}

// MockCapitalGainServiceMockRecorder é o gravador de chamadas para o MockCapitalGainService.
// Ele permite definir expectativas sobre como o mock deve ser chamado durante os testes.
type MockCapitalGainServiceMockRecorder struct {
	mock *MockCapitalGainService
}

// NewMockCapitalGainService cria e retorna uma nova instância de MockCapitalGainService.
//
// Parâmetros:
//   - ctrl: ponteiro para *gomock.Controller responsável pelo gerenciamento do mock.
//
// Retorno:
//   - Um ponteiro para a nova instância de MockCapitalGainService.
func NewMockCapitalGainService(ctrl *gomock.Controller) *MockCapitalGainService {
	mock := &MockCapitalGainService{ctrl: ctrl}
	mock.recorder = &MockCapitalGainServiceMockRecorder{mock}
	return mock
}

// EXPECT retorna um objeto que permite ao chamador definir expectativas
// sobre as chamadas feitas ao mock durante o teste.
func (m *MockCapitalGainService) EXPECT() *MockCapitalGainServiceMockRecorder {
	return m.recorder
}

// GetCapitalGain é um método mock da interface CapitalGainService.
//
// Parâmetros:
//   - name: nome da entidade/identificador utilizado para buscar os ganhos de capital.
//
// Retorno:
//   - Um ponteiro para uma matriz de slices de CapitalGains.
//   - Um erro, caso ocorra falha na busca.
//
// Este método é usado para simular a lógica de obtenção de ganhos de capital
// sem acessar uma implementação real.
func (m *MockCapitalGainService) GetCapitalGain(name string) (*[][]models.CapitalGains, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapitalGain", name)
	ret0, _ := ret[0].(*[][]models.CapitalGains)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCapitalGain define uma expectativa de chamada para o método GetCapitalGain
// durante a execução dos testes.
//
// Parâmetros:
//   - name: valor esperado do parâmetro name.
//
// Retorno:
//   - Um ponteiro para gomock.Call, permitindo configurar respostas e validações.
func (mr *MockCapitalGainServiceMockRecorder) GetCapitalGain(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapitalGain", reflect.TypeOf((*MockCapitalGainService)(nil).GetCapitalGain), name)
}
