package ports

import (
	"encoding/json"
	"testing"

	"koto.com/internal/core/models"
)

func TestGetStock_EmptyInput(t *testing.T) {
	service := NewStockService()
	stock, err := service.GetStock("")
	if err == nil || stock != nil {
		t.Errorf("Expected error for empty input, got stock: %v, err: %v", stock, err)
	}
}

func TestGetStock_BuyOperation(t *testing.T) {
	service := NewStockService()
	ops := []*models.Operation{
		{
			Operation: "buy",
			Quantity:  10,
			UnitCost:  100.0,
		},
	}
	data, _ := json.Marshal(ops)
	stock, err := service.GetStock(string(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if stock.StockQuantity != 10 {
		t.Errorf("Expected StockQuantity 10, got %d", stock.StockQuantity)
	}
	if stock.AveragePurchasePrice != 100.0 {
		t.Errorf("Expected AveragePurchasePrice 100.0, got %f", stock.AveragePurchasePrice)
	}
}

func TestGetStock_SellOperation_NoTax(t *testing.T) {
	service := NewStockService()
	ops := []*models.Operation{
		{
			Operation: "buy",
			Quantity:  10,
			UnitCost:  100.0,
		},
		{
			Operation: "sell",
			Quantity:  5,
			UnitCost:  100.0,
		},
	}
	data, _ := json.Marshal(ops)
	stock, err := service.GetStock(string(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if stock.StockQuantity != 5 {
		t.Errorf("Expected StockQuantity 5, got %d", stock.StockQuantity)
	}
	// Tax should be zero since totalSale <= 20000
	if stock.Operations[1].Tax != 0 {
		t.Errorf("Expected Tax 0, got %f", stock.Operations[1].Tax)
	}
}

func TestGetStock_SellOperation_WithTax(t *testing.T) {
	service := NewStockService()
	ops := []*models.Operation{
		{
			Operation: "buy",
			Quantity:  100,
			UnitCost:  300.0,
		},
		{
			Operation: "sell",
			Quantity:  50,
			UnitCost:  500.0,
		},
	}
	data, _ := json.Marshal(ops)
	stock, err := service.GetStock(string(data))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Tax should be calculated since totalSale > 20000
	expectedTax := models.MyFloat64(((500.0 * 50) - (300.0 * 50)) * 0.2)
	if stock.Operations[1].Tax != expectedTax {
		t.Errorf("Expected Tax %f, got %f", expectedTax, stock.Operations[1].Tax)
	}
}

func TestRecalculateStockQuantity_BuyAndSell(t *testing.T) {
	service := NewStockService()
	stock := &models.Stock{}
	buy := &models.Operation{Operation: "buy", Quantity: 10}
	sell := &models.Operation{Operation: "sell", Quantity: 5}
	service.(*stockServiceImpl).recalculateStockQuantity(stock, buy)
	if stock.StockQuantity != 10 {
		t.Errorf("Expected StockQuantity 10 after buy, got %d", stock.StockQuantity)
	}
	service.(*stockServiceImpl).recalculateStockQuantity(stock, sell)
	if stock.StockQuantity != 5 {
		t.Errorf("Expected StockQuantity 5 after sell, got %d", stock.StockQuantity)
	}
}
