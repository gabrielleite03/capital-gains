// Stock represents a stock holding in the portfolio.
// It contains information about the quantity held, average purchase price,
// accumulated loss, and a list of operations performed on the stock.
package models

type Stock struct {
	// StockQuantity represents the number of stocks held in the portfolio.
	StockQuantity uint64 `json:"stock-quantity"`
	// AveragePurchasePrice represents the average price at which the stocks were purchased.
	AveragePurchasePrice float64 `json:"average-purchase-price"`
	// Loss represents the accumulated loss from the stock operations.
	Loss float64 `json:"loss"`
	// Operations represents the list of operations performed on the stock.
	Operations []*Operation `json:"operations"`
}
