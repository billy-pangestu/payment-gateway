package model

import (
	"database/sql"
)

// SQLGdbc ...
type SQLGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	// If need transaction support, add this interface
	Transactioner
}

// SQLDBTx is the concrete implementation of sqlGdbc by using *sql.DB
type SQLDBTx struct {
	DB *sql.DB
}

// SQLConnTx is the concrete implementation of sqlGdbc by using *sql.Tx
type SQLConnTx struct {
	DB *sql.Tx
}

// Transactioner is the transaction interface for database handler
// It should only be applicable to SQL database
type Transactioner interface {
	// Rollback a transaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise rollback
	// txFunc is the operations wrapped in a transaction
	TxEnd(txFunc func() error) error
	// TxBegin gets *sql.DB from receiver and return a SqlGdbc, which has a *sql.Tx
	TxBegin() (SQLGdbc, error)
}

// TxBegin starts a transaction
func (sdt *SQLDBTx) TxBegin() (*SQLConnTx, error) {
	tx, err := sdt.DB.Begin()
	sct := SQLConnTx{tx}
	return &sct, err
}

// TxEnd ...
func (sct *SQLConnTx) TxEnd(txFunc func() error) error {
	var err error
	tx := sct.DB

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}

// Rollback ...
func (sct *SQLConnTx) Rollback() error {
	return sct.DB.Rollback()
}

// Commit ...
func (sct *SQLConnTx) Commit() error {
	return sct.DB.Commit()
}
