package ports

import (
	"encoding/json"
	"errors"

	"koto.com/internal/core/models"
)

type StockService interface {
	GetStock(name string) (*models.Stock, error)
	recalculateWeightedAveragePrice(s *models.Stock, o *models.Operation) error
	recalculateStockQuantity(s *models.Stock, o *models.Operation) error
}

type stockServiceImpl struct {
}

func NewStockService() StockService {
	return &stockServiceImpl{}
}

// GetStock implements StockService.
func (s *stockServiceImpl) GetStock(in string) (*models.Stock, error) {
	var operations []*models.Operation
	if in == "" {
		return nil, errors.New("input string is empty")
	}
	err := json.Unmarshal([]byte(in), &operations)
	if err != nil {
		return nil, err
	}
	stock := &models.Stock{
		Operations: operations,
	}

	for i, o := range stock.Operations {
		// Recalculate the stock quantity and average price for each operation
		s.recalculateStockQuantity(stock, o)
		s.recalculateWeightedAveragePrice(stock, o)

		stock.Operations[i] = o
	}
	return stock, nil
}

// recalculateStockQuantity implements StockService.
func (*stockServiceImpl) recalculateStockQuantity(s *models.Stock, o *models.Operation) error {
	if o.IsBuy() {
		o.InitialQuantity = s.StockQuantity
		s.StockQuantity += uint64(o.Quantity)
		o.FinalQuantity = s.StockQuantity
	} else if o.IsSell() {
		o.InitialQuantity = s.StockQuantity
		s.StockQuantity -= uint64(o.Quantity)
		o.FinalQuantity = s.StockQuantity
	}
	return nil
}

// recalculateWeightedAveragePrice implements StockService.
func (ss *stockServiceImpl) recalculateWeightedAveragePrice(s *models.Stock, o *models.Operation) error {
	if o.IsBuy() {
		if s.AveragePurchasePrice == 0 {
			s.AveragePurchasePrice = o.UnitCost
		} else {
			// Calculate the new average price based on the previous average and the new purchase
			s.AveragePurchasePrice = (s.AveragePurchasePrice*float64(s.StockQuantity) + o.UnitCost*float64(o.Quantity)) / float64(s.StockQuantity+uint64(o.Quantity))
		}

	} else if o.IsSell() {
		//qual o tal da venda??
		totalSale := float64(o.UnitCost * float64(o.Quantity))
		// qual o valor da compra?

		purchasePrice := s.AveragePurchasePrice * float64(o.Quantity)
		// quantas perdas?
		if totalSale > 20000 {
			// lucro deve levar em considerção sse ouve debito em operções passadas
			lucro := totalSale - purchasePrice

			if lucro < 0 {
				// se o lucro for negativo, significa que houve uma perda
				// Deduzir a perda do total de ações
				s.Loss += -lucro // Add the loss to the stock's total loss
				o.Tax = 0        // No tax on losses
			} else {
				// se o lucro for positivo, calcular o imposto
				// 20% de imposto sobre o lucro
				// calcular valor a ser taxado
				taxValue := lucro - s.Loss
				if taxValue < 0 {
					// se o valor a ser taxado for negativo, significa que não há lucro a ser taxado
					s.Loss = -taxValue // Update the loss to the negative tax value
					taxValue = 0
				} else {
					// se o valor a ser taxado for positivo, significa que há lucro a ser taxado
					// Reset the loss since we have a profit now
					s.Loss = 0
				}
				// calcular o imposto
				tax := models.MyFloat64((taxValue) * 0.2) // Assuming a 20% tax on profit
				o.Tax = tax
			}

		}
	}
	return nil
}
