// Pacote ports contém a definição da interface StockService e sua implementação.
// Ele é responsável por processar operações de compra e venda de ações,
// recalculando a quantidade em carteira, o preço médio de compra e o imposto sobre lucro.
package ports

import (
	"encoding/json"
	"errors"

	"koto.com/internal/core/infrastructure"
	"koto.com/internal/core/models"
)

// StockService define os métodos relacionados ao processamento de ações em um portfólio.
type StockService interface {
	// GetStock processa uma string JSON contendo operações de compra e venda de ações,
	// retornando um objeto Stock com as informações calculadas.
	GetStock(name string) (*models.Stock, error)

	// recalculateStockQuantity recalcula a quantidade de ações, o preço médio e as perdas
	// acumuladas com base em uma operação específica.
	recalculateStockQuantity(s *models.Stock, o *models.Operation) error
}

// stockServiceImpl é a implementação concreta da interface StockService.
type stockServiceImpl struct {
}

// NewStockService retorna uma nova instância de StockService.
func NewStockService() StockService {
	return &stockServiceImpl{}
}

// GetStock implementa StockService.
// Ele recebe uma string JSON contendo operações (compras/vendas),
// reconstrói o objeto Stock e recalcula o estado após cada operação.
func (s *stockServiceImpl) GetStock(in string) (*models.Stock, error) {
	var operations []*models.Operation

	// Valida se a entrada não está vazia
	if in == "" {
		return nil, errors.New("a string de entrada está vazia")
	}

	// Converte o JSON de entrada em uma lista de operações
	err := json.Unmarshal([]byte(in), &operations)
	if err != nil {
		return nil, err
	}

	// Cria um objeto Stock inicial com a lista de operações
	stock := &models.Stock{
		Operations: operations,
	}

	// Processa cada operação, recalculando preço médio, quantidade e perdas
	for _, o := range stock.Operations {
		s.recalculateStockQuantity(stock, o)
	}
	return stock, nil
}

// recalculateStockQuantity implementa StockService.
// Ele recalcula a quantidade de ações, preço médio, perdas acumuladas e impostos,
// de acordo com a operação (compra ou venda).
func (*stockServiceImpl) recalculateStockQuantity(s *models.Stock, o *models.Operation) error {

	// Se a operação for uma compra
	if o.IsBuy() {

		// Se for a primeira compra, define o preço médio como o valor da operação
		if s.AveragePurchasePrice == 0 {
			s.AveragePurchasePrice = o.UnitCost
		} else {
			// Recalcula o preço médio ponderado pelo número de ações
			s.AveragePurchasePrice = (s.AveragePurchasePrice*float64(s.StockQuantity) + o.UnitCost*float64(o.Quantity)) /
				float64(s.StockQuantity+uint64(o.Quantity))
		}

		// Atualiza as quantidades inicial e final da operação
		o.InitialQuantity = s.StockQuantity
		s.StockQuantity += uint64(o.Quantity)
		o.FinalQuantity = s.StockQuantity

		// Se a operação for uma venda
	} else if o.IsSell() {

		// Valor total da venda (preço unitário * quantidade)
		totalSale := float64(o.UnitCost * float64(o.Quantity))

		// Valor de compra das ações vendidas (preço médio * quantidade)
		purchasePrice := s.AveragePurchasePrice * float64(o.Quantity)

		// Apenas vendas com valor acima de 20.000 são consideradas para tributação
		if totalSale > 20000 {
			// Calcula o lucro (ou prejuízo) da operação
			lucro := totalSale - purchasePrice

			if lucro < 0 {
				// Caso o resultado seja prejuízo, acumula a perda
				s.Loss += -lucro
				o.Tax = 0 // Não há imposto sobre prejuízo
			} else {
				// Caso haja lucro, deduz eventuais prejuízos acumulados
				taxValue := lucro - s.Loss

				if taxValue < 0 {
					// Se o valor a ser taxado for negativo, significa que não há lucro tributável
					s.Loss = -taxValue // Mantém o prejuízo remanescente
					taxValue = 0
				} else {
					// Se o valor a ser taxado for positivo, zera o prejuízo acumulado
					s.Loss = 0
				}

				// Aplica imposto de 20% sobre o lucro tributável
				o.Tax = models.MyFloat64(taxValue * infrastructure.TaxOnProfit)
			}
		}

		// Atualiza as quantidades inicial e final da operação
		o.InitialQuantity = s.StockQuantity
		s.StockQuantity -= uint64(o.Quantity)
		o.FinalQuantity = s.StockQuantity
	}

	return nil
}
