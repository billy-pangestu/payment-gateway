package model

import (
	"database/sql"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

// ITransaction ...
type ITransaction interface {
	Store(body viewmodel.TransactionVM, now time.Time) (string, error)
}

// TransactionEntity ...
type TransactionEntity struct {
	ID        string          `db:"id"`
	UserID    string          `db:"user_id"`
	MoneyIn   sql.NullFloat64 `db:"money_in"`
	MoneyOut  sql.NullFloat64 `db:"money_out"`
	Notes     string          `db:"notes"`
	Tags      sql.NullString  `db:"tags"`
	CreatedAt string          `db:"created_at"`
	UpdatedAt sql.NullString  `db:"updated_at"`
	DeletedAt sql.NullString  `db:"deleted_at"`
}

// TransactionModel ...
type TransactionModel struct {
	DB *sql.Tx
}

// NewTransactionModel ...
func NewTransactionModel(db *sql.Tx) ITransaction {
	return &TransactionModel{DB: db}
}

// Store ...
func (model TransactionModel) Store(body viewmodel.TransactionVM, now time.Time) (data string, err error) {
	query :=
		`INSERT INTO "transactions" ("user_id", "merchant_id", "amount", "created_at", "updated_at")
		VALUES($1, $2, $3, $4, $4)
		RETURNING "id"`
	err = model.DB.QueryRow(query,
		body.UserID, body.MerchID, body.Amount, now,
	).Scan(&data)

	return data, err
}
