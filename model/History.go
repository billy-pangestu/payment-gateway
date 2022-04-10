package model

import (
	"database/sql"
	"time"
)

var ()

// IHistory ...
type IHistory interface {
	Store(qid, payload, api, status, errorstring string, now time.Time) (string, error)
}

// HistoryEntity ....
type HistoryEntity struct {
	ID        string         `db:"id"`
	Code      sql.NullString `db:"code"`
	Name      sql.NullString `db:"name"`
	Status    sql.NullBool   `db:"status"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}

// historyModel ...
type historyModel struct {
	DB *sql.Tx
}

// NewHistoryModel ...
func NewHistoryModel(db *sql.Tx) IHistory {
	return &historyModel{DB: db}
}

// Store ...
func (model historyModel) Store(qid, payload, api, status, errorstring string, now time.Time) (res string, err error) {
	query := `
		INSERT INTO "history" (
			"qid", "payload", "api", "status", "error_string",
			"created_at", "updated_at"
		)
		VALUES (
			$1, $2, $3, $4, $5,
			$6, $6
		)
		RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		qid, payload, api, status, errorstring,
		now,
	).Scan(&res)

	return res, err
}
