package tests

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "koto.com/internal/core/models"
)

// MockCapitalGainService is a mock of CapitalGainService interface.
type MockCapitalGainService struct {
	ctrl     *gomock.Controller
	recorder *MockCapitalGainServiceMockRecorder
}

// MockCapitalGainServiceMockRecorder is the mock recorder for MockCapitalGainService.
type MockCapitalGainServiceMockRecorder struct {
	mock *MockCapitalGainService
}

// NewMockCapitalGainService creates a new mock instance.
func NewMockCapitalGainService(ctrl *gomock.Controller) *MockCapitalGainService {
	mock := &MockCapitalGainService{ctrl: ctrl}
	mock.recorder = &MockCapitalGainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCapitalGainService) EXPECT() *MockCapitalGainServiceMockRecorder {
	return m.recorder
}

// GetCapitalGain mocks base method.
func (m *MockCapitalGainService) GetCapitalGain(name string) (*[]models.CapitalGains, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapitalGain", name)
	ret0, _ := ret[0].(*[]models.CapitalGains)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCapitalGain indicates an expected call of GetCapitalGain.
func (mr *MockCapitalGainServiceMockRecorder) GetCapitalGain(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapitalGain", reflect.TypeOf((*MockCapitalGainService)(nil).GetCapitalGain), name)
}
