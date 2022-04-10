package model

import (
	"database/sql"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

var (
	userSelectQuery = `SELECT def."id", def."first_name", def."last_name", def."unique_id", def."password", def."amount",
	def."created_at", def."updated_at", def."deleted_at",
	ur."name"
	from users def
	LEFT JOIN user_roles ur on ur.id = def.role_id `
)

// IUser ...
type IUser interface {
	FindByUniqueID(UniqueID string) (UserEntity, error)
	FindByID(id string) (UserEntity, error)
	Store(data viewmodel.UserVM, now time.Time) (string, error)
	AddFund(userID string, data viewmodel.UserVM, now time.Time) (string, error)
	SubFund(userID string, data viewmodel.UserVM, now time.Time) (string, error)
}

// IUserWithoutTX ...
type IUserWithoutTX interface {
	FindByIDWithoutTX(id string) (UserEntity, error)
}

// UserEntity ....
type UserEntity struct {
	ID        string         `db:"id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	UniqueID  sql.NullString `db:"unique_id"`
	Password  sql.NullString `db:"password"`
	Amount    float64        `db:"amount"`
	CreatedAt sql.NullString `db:"createdAt"`
	UpdatedAt sql.NullString `db:"updatedAt"`
	DeletedAt sql.NullString `db:"deletedAt"`

	UserRoleName sql.NullString `db:"user_role_name"`
}

// userModel ...
type userModel struct {
	DB *sql.Tx
}

// NewUserModel ...
func NewUserModel(db *sql.Tx) IUser {
	return &userModel{DB: db}
}

// userModelWithoutTX ...
type userModelWithoutTX struct {
	DB *sql.DB
}

// NewUserModel ...
func NewUserModelWithoutTX(db *sql.DB) IUserWithoutTX {
	return &userModelWithoutTX{DB: db}
}

func (model userModel) scanRows(rows *sql.Rows) (d UserEntity, err error) {
	err = rows.Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.UniqueID, &d.Password, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.UserRoleName,
	)

	return d, err
}

func (model userModel) scanRow(row *sql.Row) (d UserEntity, err error) {
	err = row.Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.UniqueID, &d.Password, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.UserRoleName,
	)

	return d, err
}

func (model userModelWithoutTX) scanRows(rows *sql.Rows) (d UserEntity, err error) {
	err = rows.Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.UniqueID, &d.Password, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.UserRoleName,
	)

	return d, err
}

func (model userModelWithoutTX) scanRow(row *sql.Row) (d UserEntity, err error) {
	err = row.Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.UniqueID, &d.Password, &d.Amount,
		&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.UserRoleName,
	)

	return d, err
}

// FindByID ...
func (model userModel) FindByID(id string) (data UserEntity, err error) {
	query := userSelectQuery +
		` WHERE def."deleted_at" is null AND def."id"=$1`

	row := model.DB.QueryRow(query, id)
	data, err = model.scanRow(row)

	return data, err
}

// FindByIDWithoutTX ...
func (model userModelWithoutTX) FindByIDWithoutTX(id string) (data UserEntity, err error) {
	query := userSelectQuery +
		` WHERE def."deleted_at" is null AND def."id"=$1`
	row := model.DB.QueryRow(query, id)
	data, err = model.scanRow(row)

	return data, err
}

//FindByUniqueID ...
func (model userModel) FindByUniqueID(uniqueID string) (data UserEntity, err error) {
	query := userSelectQuery +
		` WHERE def."deleted_at" is null AND def."unique_id"=$1`
	row := model.DB.QueryRow(query, uniqueID)
	data, err = model.scanRow(row)

	return data, err
}

// Store ...
func (model userModel) Store(data viewmodel.UserVM, now time.Time) (res string, err error) {
	query := `
		INSERT INTO "users" (
			"first_name", "last_name", "unique_id", "password", "role_id", 
			"created_at", "updated_at"
		)
		VALUES (
			$1, $2, $3, $4, $5,
			$6, $6
		)
		RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		data.FirstName, data.LastName, data.UniqueID, data.Password, data.RoleID,
		now,
	).Scan(&res)

	return res, err
}

// AddFund ...
func (model userModel) AddFund(userID string, data viewmodel.UserVM, now time.Time) (res string, err error) {
	query := `
		UPDATE "users" SET "amount" = amount+$1, "updated_at" = $2
		WHERE "deleted_at" IS NULL AND "id" = $3 RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		data.Amount, now,
		userID,
	).Scan(&res)

	return res, err
}

// SubFund ...
func (model userModel) SubFund(userID string, data viewmodel.UserVM, now time.Time) (res string, err error) {
	query := `
		UPDATE "users" SET "amount" = amount-$1, "updated_at" = $2
		WHERE "deleted_at" IS NULL AND "id" = $3 RETURNING "id"
	`
	err = model.DB.QueryRow(query,
		data.Amount, now,
		userID,
	).Scan(&res)

	return res, err
}
