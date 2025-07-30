// Pacote models contém definições de dados e lógicas associadas a operações financeiras,
// como compra e venda de ativos.
package models

// Operation representa uma transação financeira envolvendo um ativo,
// como uma operação de compra ou venda.
//
// Essa estrutura inclui:
//   - O tipo da operação (compra ou venda).
//   - O custo por unidade do ativo.
//   - A quantidade de unidades envolvidas.
//   - As quantidades antes e depois da operação.
//   - O imposto aplicado à transação.
type Operation struct {
	// Operation indica o tipo da operação: "buy" (compra) ou "sell" (venda).
	Operation Op `json:"operation"`

	// UnitCost é o custo por unidade do ativo na operação.
	UnitCost float64 `json:"unit-cost"`

	// Quantity é a quantidade de unidades envolvidas na operação.
	Quantity uint64 `json:"quantity"`

	// InitialQuantity representa a quantidade do ativo antes da operação.
	InitialQuantity uint64 `json:"initial-quantity"`

	// FinalQuantity representa a quantidade do ativo após a operação.
	FinalQuantity uint64 `json:"final-quantity"`

	// Tax é o imposto aplicado à transação, com formatação personalizada para JSON.
	Tax MyFloat64 `json:"tax"`
}

// Op define o tipo para identificar operações financeiras.
//
// Os valores esperados são as constantes OpBuy ("buy") e OpSell ("sell").
type Op string

// OpBuy representa uma operação de compra.
var OpBuy Op = "buy"

// OpSell representa uma operação de venda.
var OpSell Op = "sell"

// IsBuy retorna true se a operação for do tipo "buy" (compra).
func (op *Operation) IsBuy() bool {
	return op.Operation == OpBuy
}

// IsSell retorna true se a operação for do tipo "sell" (venda).
func (op *Operation) IsSell() bool {
	return op.Operation == OpSell
}
