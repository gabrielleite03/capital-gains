package domain

type InputTO struct {
	Operations []Operation
}

type Operation struct {
	Operation Op `json:"operation"`
	// UnitCost represents the cost per individual unit of the asset.
	UnitCost        float64 `json:"unit-cost"`
	Quantity        uint64  `json:"quantity"`
	InitialQuantity uint64  `json:"initial-quantity"`
	FinalQuantity   uint64  `json:"final-quantity"`

	// Tax represents the tax applied to the operation.
	Tax MyFloat64 `json:"tax"`
}

type Op string

var OpBuy Op = "buy"
var OpSell Op = "sell"

func (op *Operation) IsBuy() bool {
	return op.Operation == OpBuy
}
func (op *Operation) IsSell() bool {
	return op.Operation == OpSell
}
