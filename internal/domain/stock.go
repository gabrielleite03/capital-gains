package domain

type Stock struct {
	Name string `json:"name"`
	// StockQuantity represents the number of stocks held in the portfolio.
	StockQuantity        uint64  `json:"stock-quantity"`
	AveragePurchasePrice float64 `json:"average-purchase-price"`

	Loss       float64     `json:"loss"`
	Operations []Operation `json:"operations"`
}
