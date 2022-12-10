package request

// TransactionRequest ...
type TransactionRequest struct {
	MerchandID string  `json:"merchant_id" validate:"required"`
	Amount     float64 `json:"amount"`
}
