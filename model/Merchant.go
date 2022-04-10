package model

import (
	"database/sql"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

var (
	merhcantSelectQuery = `SELECT def."id", def."name", def."amount",
	def."created_at", def."updated_at", def."deleted_at"
	from merchants def `
)

// IMerchant ...
type IMerchant interface {
	FindAll(offset, limit int) ([]MerchantEntity, int, error)
	FindByID(id string) (MerchantEntity, error)
	Store(data viewmodel.MerchantVM, now time.Time) (string, error)
	AddFund(userID string, data viewmodel.MerchantVM, now time.Time) (string, error)
}

// MerchantEntity ....
type MerchantEntity struct {
	ID        string         `db:"id"`
	Name      sql.NullString `db:"name"`
	Amount    float64        `db:"amount"`
	CreatedAt sql.NullString `db:"createdAt"`
	UpdatedAt sql.NullString `db:"updatedAt"`
	DeletedAt sql.NullString `db:"deletedAt"`
}

// merchantModel ...
type merchantModel struct {
	DB *sql.Tx
}

// NewMerchantModel ...
func NewMerchantModel(db *sql.Tx) IMerchant {
	return &merchantModel{DB: db}
}

func (model merchantModel) scanRows(rows *sql.Rows) (d MerchantEntity, err error) {
	err = rows.Scan(
		&d.ID, &d.Name, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
	)

	return d, err
}

func (model merchantModel) scanRow(row *sql.Row) (d MerchantEntity, err error) {
	err = row.Scan(
		&d.ID, &d.Name, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
	)

	return d, err
}

// FindAll ...
func (model merchantModel) FindAll(offset, limit int) (res []MerchantEntity, count int, err error) {
	query := merhcantSelectQuery + ` WHERE def."deleted_at" IS NULL 
	ORDER BY def."name" ASC OFFSET $1 LIMIT $2`
	rows, err := model.DB.Query(query, offset, limit)

	if err != nil {
		return res, count, err
	}

	defer rows.Close()
	for rows.Next() {
		d, err := model.scanRows(rows)
		if err != nil {
			return res, count, err
		}
		res = append(res, d)
	}
	err = rows.Err()

	if err != nil {
		return res, count, err
	}

	// Query row count
	query = `SELECT COUNT(*) FROM "merchants" def WHERE def."deleted_at" IS NULL `
	err = model.DB.QueryRow(query).Scan(&count)

	return res, count, err
}

// FindByID ...
func (model merchantModel) FindByID(id string) (data MerchantEntity, err error) {
	query := merhcantSelectQuery +
		` WHERE def."deleted_at" is null AND def."id"=$1`
	row := model.DB.QueryRow(query, id)
	data, err = model.scanRow(row)

	return data, err
}

// Store ...
func (model merchantModel) Store(data viewmodel.MerchantVM, now time.Time) (res string, err error) {
	query := `
		INSERT INTO "merchants" (
			"name",
			"created_at", "updated_at"
		)
		VALUES (
			$1,
			$2, $2
		)
		RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		data.Name,
		now,
	).Scan(&res)

	return res, err
}

// AddFund ...
func (model merchantModel) AddFund(userID string, data viewmodel.MerchantVM, now time.Time) (res string, err error) {
	query := `
		UPDATE "merchants" SET "amount" = amount+$1, "updated_at" = $2
		WHERE "deleted_at" IS NULL AND "id" = $3 RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		data.Amount, now,
		userID,
	).Scan(&res)

	return res, err
}
