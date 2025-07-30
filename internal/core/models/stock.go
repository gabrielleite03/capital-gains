// Stock representa uma ação (ou ativo) mantida no portfólio.
// Contém informações sobre a quantidade de ações, preço médio de compra,
// prejuízo acumulado e a lista de operações realizadas com a ação.
package models

type Stock struct {
	// StockQuantity representa a quantidade de ações mantidas no portfólio.
	StockQuantity uint64 `json:"stock-quantity"`
	// AveragePurchasePrice representa o preço médio pelo qual as ações foram compradas.
	AveragePurchasePrice float64 `json:"average-purchase-price"`
	// Loss representa o prejuízo acumulado proveniente das operações com a ação.
	Loss float64 `json:"loss"`
	// Operations representa a lista de operações realizadas com a ação.
	Operations []*Operation `json:"operations"`
}
