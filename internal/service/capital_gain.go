package service

import (
	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
)

type CapitalGainService interface {
	GetCapitalGain(name string) (*[]models.CapitalGains, error)
}

type capitalGainService struct {
	stockService ports.StockService
}

func NewCapitalGainService(stockService ports.StockService) CapitalGainService {
	return &capitalGainService{
		stockService: stockService,
	}
}

// GetCapitalGain implements CapitalGainService.
func (c *capitalGainService) GetCapitalGain(name string) (*[]models.CapitalGains, error) {
	// Retrieve the stock using the stock service
	stock, err := c.stockService.GetStock(name)
	if err != nil {
		return nil, err

	}
	var cg []models.CapitalGains
	for _, op := range stock.Operations {
		cg = append(cg, models.CapitalGains{
			Tax: models.MyFloat64(op.Tax),
		})
	}

	return &cg, nil
}
