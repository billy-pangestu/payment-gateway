package viewmodel

//TransactionWithTotalAmount ...
type TransactionWithTotalAmount struct {
	TotalMoneyIn  float64         `json:"total_money_in"`
	TotalMoneyOut float64         `json:"total_money_Out"`
	Transactions  []TransactionVM `json:"transactions"`
}

// TransactionVM ...
type TransactionVM struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	MoneyOut  float64 `json:"money_out"`
	MoneyIn   float64 `json:"money_in"`
	Notes     string  `json:"notes"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}
