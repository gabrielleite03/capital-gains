package domain

import (
	"fmt"
	"strconv"
)

type OperationResult struct {
	Tax MyFloat64 `json:"tax"`
}

// Define a local type alias for float64
type MyFloat64 float64

func (mf MyFloat64) MarshalJSON() ([]byte, error) {
	// Format the float64 to ensure at least one decimal place.
	// 'f' format ensures fixed-point notation.
	// You can adjust the precision (e.g., "%.1f", "%.2f") as needed.
	s := strconv.FormatFloat(float64(mf), 'f', -1, 64)

	// Example: ensuring at least one digit after the decimal for integral values
	if float64(mf) == float64(int64(mf)) { // Check if it's an integral value
		return []byte(fmt.Sprintf("%.1f", mf)), nil // Format with at least one decimal place
	}

	return []byte(s), nil
}
