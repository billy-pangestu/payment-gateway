package model

import (
	"database/sql"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

// ITransaction ...
type ITransaction interface {
	FindByID(id, userID string) (TransactionEntity, error)
	FindByUserID(userID string) ([]TransactionEntity, error)
	FindByTag(userID, tag string) ([]TransactionEntity, error)
	FindTotalAmount(userID string) (TransactionEntity, error)
	FindTotalMoneyByTag(userID, tag string) (TransactionEntity, error)
	Store(body viewmodel.TransactionVM, now time.Time, tx *sql.Tx) (string, error)
	Update(id string, body viewmodel.TransactionVM, changedAt time.Time, tx *sql.Tx) (TransactionEntity, error)
	Destroy(id string, changedAt time.Time) (TransactionEntity, error)
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
	DB *sql.DB
}

// NewTransactionModel ...
func NewTransactionModel(db *sql.DB) ITransaction {
	return &TransactionModel{DB: db}
}

// FindByID ...
func (model TransactionModel) FindByID(id, userID string) (data TransactionEntity, err error) {
	query := `SELECT t."id", t."user_id", t."money_in", t."money_out", t."notes", t."created_at", t."updated_at",
		array_to_string(array_agg(tag.hash_tags),',')
		FROM "transactions" t
		LEFT JOIN "tags" tag ON tag."transaction_id" = t."id"
		WHERE t."id" = $1 AND t."user_id" = $2 AND t."deleted_at" IS NULL
		GROUP BY t."id"`
	err = model.DB.QueryRow(query, id, userID).Scan(
		&data.ID, &data.UserID, &data.MoneyIn, &data.MoneyOut, &data.Notes, &data.CreatedAt, &data.UpdatedAt, &data.Tags,
	)

	return data, err
}

// FindByUserID ...
func (model TransactionModel) FindByUserID(userID string) (data []TransactionEntity, err error) {
	query :=
		`SELECT t."id", t."user_id", t."money_in", t."money_out", t."notes", t."created_at", t."updated_at",
		array_to_string(array_agg(tag.hash_tags),',')
		FROM "transactions" t
		LEFT JOIN "tags" tag ON tag."transaction_id" = t."id"
		WHERE t."user_id" = $1 AND t."deleted_at" IS NULL AND tag."deleted_at" IS NULL
		GROUP BY t."id"
		ORDER BY t."created_at" DESC`
	rows, err := model.DB.Query(query, userID)
	dataTemp := TransactionEntity{}
	for rows.Next() {
		err = rows.Scan(
			&dataTemp.ID, &dataTemp.UserID, &dataTemp.MoneyIn, &dataTemp.MoneyOut, &dataTemp.Notes,
			&dataTemp.CreatedAt, &dataTemp.UpdatedAt, &dataTemp.Tags,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

// FindByTag ...
func (model TransactionModel) FindByTag(userID, tag string) (data []TransactionEntity, err error) {
	query :=
		`SELECT transactions."id", transactions."money_in", transactions."money_out", transactions."notes",
		transactions."created_at", transactions."updated_at", transactions."deleted_at"
		FROM "transactions"
		JOIN "tags"
		ON transactions."id" = tags."transaction_id"
		WHERE transactions."user_id" = $1 AND LOWER(tags."hash_tags") = LOWER($2) AND transactions."deleted_at" IS NULL
		ORDER BY transactions."created_at" DESC`
	rows, err := model.DB.Query(query, userID, tag)
	dataTemp := TransactionEntity{}
	for rows.Next() {
		err = rows.Scan(
			&dataTemp.ID, &dataTemp.MoneyIn, &dataTemp.MoneyOut, &dataTemp.Notes, &dataTemp.CreatedAt, &dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

// FindTotalMoneyByTag ...
func (model TransactionModel) FindTotalMoneyByTag(userID, tag string) (data TransactionEntity, err error) {
	query :=
		`SELECT SUM("money_in") AS "Total Debit", SUM("money_out") AS "Total Credit"
		FROM "transactions"
		JOIN "tags"
		ON transactions."id" = tags."transaction_id"
		WHERE transactions."user_id" = $1 AND LOWER(tags."hash_tags") = LOWER($2) AND transactions."deleted_at" IS NULL`
	err = model.DB.QueryRow(query, userID, tag).Scan(&data.MoneyIn, &data.MoneyOut)

	return data, err
}

// FindTotalAmount ...
func (model TransactionModel) FindTotalAmount(userID string) (data TransactionEntity, err error) {
	query :=
		`SELECT SUM("money_in") AS "Total Debit", SUM("money_out") AS "Total Credit"
		FROM "transactions"
		WHERE "user_id" = $1 AND "deleted_at" IS NULL`
	err = model.DB.QueryRow(query, userID).Scan(&data.MoneyIn, &data.MoneyOut)

	return data, err
}

// Store ...
func (model TransactionModel) Store(body viewmodel.TransactionVM, now time.Time, tx *sql.Tx) (data string, err error) {
	query :=
		`INSERT INTO "transactions" ("user_id", "money_out", "money_in", "notes", "created_at", "updated_at")
		VALUES($1, $2, $3, $4, $5, $5)
		RETURNING "id"`
	if tx != nil {
		err = tx.QueryRow(query,
			body.UserID, body.MoneyOut, body.MoneyIn, body.Notes, now,
		).Scan(&data)
	} else {
		err = model.DB.QueryRow(query,
			body.UserID, body.MoneyOut, body.MoneyIn, body.Notes, now,
		).Scan(&data)
	}

	return data, err
}

// Update ...
func (model TransactionModel) Update(id string, body viewmodel.TransactionVM, changedAt time.Time, tx *sql.Tx) (data TransactionEntity, err error) {
	query :=
		`UPDATE "transactions"
		SET "money_in" = $1, "money_out" = $2, "notes" = $3, "updated_at" = $4
		WHERE "deleted_at" IS NULL AND "id" = $5
		RETURNING "id", "money_in", "money_out", "notes", "created_at"`
	if tx != nil {
		err = tx.QueryRow(query,
			body.MoneyIn, body.MoneyOut, body.Notes, changedAt, id,
		).Scan(&data.ID, &data.MoneyIn, &data.MoneyOut, &data.Notes, &data.CreatedAt)
	} else {
		err = model.DB.QueryRow(query,
			body.MoneyIn, body.MoneyOut, body.Notes, changedAt, id,
		).Scan(&data.ID, &data.MoneyIn, &data.MoneyOut, &data.Notes, &data.CreatedAt)
	}

	return data, err
}

// Destroy ...
func (model TransactionModel) Destroy(id string, changedAt time.Time) (data TransactionEntity, err error) {
	query :=
		`UPDATE "transactions" 
		SET "updated_at" = $1, "deleted_at" = $1
		WHERE "deleted_at" IS NULL AND "id" = $2 
		RETURNING "id", "user_id", "deleted_at"`
	err = model.DB.QueryRow(query, changedAt, id).Scan(&data.ID, &data.UserID, &data.DeletedAt)

	return data, err
}
