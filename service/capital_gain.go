// Pacote service contém a definição e implementação do CapitalGainService.
// Este serviço é responsável por calcular os ganhos de capital (e impostos associados)
// a partir de uma lista de operações de compra e venda de ações.
package service

import (
	"strings"

	"koto.com/internal/core/models"
	"koto.com/internal/core/ports"
)

// CapitalGainService define os métodos relacionados ao cálculo de ganhos de capital.
type CapitalGainService interface {
	// GetCapitalGain processa um conjunto de linhas contendo operações de ações,
	// retornando uma matriz com os ganhos de capital de cada operação.
	GetCapitalGain(name string) (*[][]models.CapitalGains, error)
}

// capitalGainService é a implementação concreta de CapitalGainService.
type capitalGainService struct {
	// stockService é usado para processar as operações e gerar os dados consolidados de ações.
	stockService ports.StockService
}

// NewCapitalGainService retorna uma nova instância do CapitalGainService,
// recebendo como dependência o serviço de ações (StockService).
func NewCapitalGainService(stockService ports.StockService) CapitalGainService {
	return &capitalGainService{
		stockService: stockService,
	}
}

// GetCapitalGain implementa CapitalGainService.
// Ele processa um texto contendo múltiplas linhas de operações (no formato JSON),
// utiliza o StockService para calcular os dados de cada ação e retorna
// os ganhos de capital (impostos) de cada operação.
func (c *capitalGainService) GetCapitalGain(row string) (*[][]models.CapitalGains, error) {
	// Divide a string recebida em linhas, usando quebra de linha como separador
	s := strings.Split(row, "\r\n")

	// cgAll armazenará os ganhos de capital de todas as linhas (operações)
	var cgAll [][]models.CapitalGains

	for _, row := range s {
		// Ignora linhas vazias
		if row != "" {
			var cg []models.CapitalGains

			// Usa o StockService para processar a linha (contendo as operações de ações)
			stock, err := c.stockService.GetStock(row)
			if err != nil {
				return nil, err
			}

			// Para cada operação da ação processada, extrai o imposto calculado
			for _, op := range stock.Operations {
				cg = append(cg, models.CapitalGains{
					Tax: models.MyFloat64(op.Tax), // Guarda o valor do imposto
				})
			}

			// Adiciona o conjunto de ganhos de capital desta linha ao resultado final
			cgAll = append(cgAll, cg)
		}
	}

	// Retorna a matriz com os ganhos de capital de todas as linhas processadas
	return &cgAll, nil
}
