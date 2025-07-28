// Operation represents a financial transaction involving an asset, such as buying or selling.
// It includes details about the type of operation, unit cost, quantities before and after the transaction,
// and any tax applied.
//
// Fields:
//   - Operation: The type of operation (buy or sell).
//   - UnitCost: The cost per individual unit of the asset.
//   - Quantity: The number of units involved in the operation.
//   - InitialQuantity: The quantity of the asset before the operation.
//   - FinalQuantity: The quantity of the asset after the operation.
//   - Tax: The tax applied to the transaction.
//
// Op defines the type for operation identifiers ("buy" or "sell").
//
// OpBuy and OpSell are constants representing buy and sell operations.
//
// IsBuy returns true if the operation is a buy operation.
//
// IsSell returns true if the operation is a sell operation.
// Operation represents a financial operation involving an asset, such as buying or selling.
// It contains details about the type of operation, unit cost, quantities before and after the operation,
// and any tax applied to the transaction.
package models

// Operation represents a financial transaction involving an asset.
// It contains details about the type of operation, unit cost, quantities before and after the operation,
// and any tax applied to the transaction.
type Operation struct {
	// Operation represents the type of operation (buy or sell).
	Operation Op `json:"operation"`
	// UnitCost represents the cost per individual unit of the asset.
	UnitCost float64 `json:"unit-cost"`
	// Quantity represents the number of units involved in the operation.
	Quantity uint64 `json:"quantity"`
	// InitialQuantity represents the quantity of the asset before the operation.
	InitialQuantity uint64 `json:"initial-quantity"`
	// FinalQuantity represents the quantity of the asset after the operation.
	FinalQuantity uint64 `json:"final-quantity"`
	// Tax represents the tax applied to the operation.
	Tax MyFloat64 `json:"tax"`
}

// Op defines the type for operation identifiers (buy or sell).
type Op string

// OpBuy and OpSell are constants representing buy and sell operations.
var OpBuy Op = "buy"

// OpSell is a constant representing a sell operation.
var OpSell Op = "sell"

// IsBuy returns true if the operation is a buy operation.
func (op *Operation) IsBuy() bool {
	return op.Operation == OpBuy
}

// IsSell returns true if the operation is a sell operation.
func (op *Operation) IsSell() bool {
	return op.Operation == OpSell
}
