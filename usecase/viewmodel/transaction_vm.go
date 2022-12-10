package viewmodel

// TransactionVM ...
type TransactionVM struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	MerchID   string  `json:"merchant_id"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}
