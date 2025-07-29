package service

import (
	"strings"

	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
)

type CapitalGainService interface {
	GetCapitalGain(name string) (*[][]models.CapitalGains, error)
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
func (c *capitalGainService) GetCapitalGain(row string) (*[][]models.CapitalGains, error) {
	s := strings.Split(row, "\r\n")

	var cgAll [][]models.CapitalGains
	for _, row := range s {
		if row != "" {
			var cg []models.CapitalGains
			// Retrieve the stock using the stock service
			stock, err := c.stockService.GetStock(row)
			if err != nil {
				return nil, err

			}

			for _, op := range stock.Operations {
				cg = append(cg, models.CapitalGains{
					Tax: models.MyFloat64(op.Tax),
				})
			}
			cgAll = append(cgAll, cg)
		}

	}
	return &cgAll, nil
}
