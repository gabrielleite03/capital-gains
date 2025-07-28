package service

import (
	"encoding/json"
	"testing"

	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
)

func TestGetCapitalGain_ErrorFromStockService(t *testing.T) {
	sService := ports.NewStockService()

	service := NewCapitalGainService(sService)

	cg, err := service.GetCapitalGain("test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if cg != nil {
		t.Errorf("Expected nil capital gains, got %v", cg)
	}

}

func TestGetCapitalGain_EmptyOperations(t *testing.T) {
	sService := ports.NewStockService()
	service := NewCapitalGainService(sService)

	cg, err := service.GetCapitalGain("test")
	if err == nil {
		t.Errorf("invalid character 'e' in literal true (expecting 'r') %v", err)
	}
	if cg != nil && len(*cg) != 0 {
		t.Errorf("Expected empty capital gains, got %v", cg)
	}
}

// koto arrumar testes
func TestGetCapitalGain_WithOperations(t *testing.T) {
	ops := []*models.Operation{
		{Tax: 10.5},
		{Tax: 0},
		{Tax: 7.25},
	}
	sService := ports.NewStockService()
	service := NewCapitalGainService(sService)

	data, _ := json.Marshal(ops)
	cg, err := service.GetCapitalGain(string(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cg == nil || len(*cg) != len(ops) {
		t.Fatalf("Expected %d capital gains, got %v", len(ops), cg)
	}
	for i, gain := range *cg {
		if gain.Tax != models.MyFloat64(ops[i].Tax) {
			t.Errorf("Expected Tax %f, got %f", ops[i].Tax, gain.Tax)
		}
	}
}
